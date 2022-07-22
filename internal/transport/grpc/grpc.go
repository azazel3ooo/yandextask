package grpc

import (
	"context"
	"github.com/azazel3ooo/yandextask/internal/logic"
	"github.com/azazel3ooo/yandextask/internal/models"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type ShortenerServer struct {
	UnimplementedUrlsServer

	storage       models.Storable
	cfg           models.Config
	chanForDelete chan []string
}

func (s *ShortenerServer) Get(ctx context.Context, in *GetterRequest) (*GetterResponse, error) {
	var res GetterResponse

	url, err := logic.GetUrl(in.Id, s.storage)
	if err != nil {
		res.Error = err.Error()
	} else if url == "deleted" {
		res.Url = "Gone"
	} else {
		res.Url = url
	}

	return &res, nil
}

func (s *ShortenerServer) Set(ctx context.Context, in *SetterRequest) (*SetterResponse, error) {
	var (
		res SetterResponse
	)

	id, err := logic.SetUrl(in.Url, getUser(ctx), s.storage, s.cfg.FileStoragePath, s.cfg.URLBase)
	if err != nil {
		res.Error = err.Error()
	} else {
		res.Result = id
	}

	return &res, nil
}

func (s *ShortenerServer) UserUrls(ctx context.Context, in *UserUrlsRequest) (*UserUrlsResponse, error) {
	var res UserUrlsResponse

	pairSlice, err := logic.UserUrlsGet(in.User, s.storage, s.cfg.URLBase)
	if err != nil {
		res.Error = err.Error()
	} else {
		res.FromUserResponseSlice(pairSlice)
	}

	return &res, nil
}

func (s *ShortenerServer) SetMany(ctx context.Context, in *SetManyRequest) (*SetManyResponse, error) {
	var res SetManyResponse

	set, err := logic.SetMany(in.ToSet(), getUser(ctx), s.storage, s.cfg.URLBase)
	if err != nil {
		res.Error = err.Error()
	} else {
		res.FromSet(set)
	}

	return &res, nil
}

func (s *ShortenerServer) AsyncDelete(ctx context.Context, in *AsyncDeleteReq) (*AsyncDeleteResp, error) {
	var res AsyncDeleteResp

	idsForDelete, err := logic.CheckUrlsForDelete(in.Array, getUser(ctx), s.storage)
	if err != nil {
		res.Error = err.Error()
	} else if len(idsForDelete) == 0 {
		res.Status = "No content"
	} else {
		res.Status = "Accepted"
		s.chanForDelete <- idsForDelete
	}

	return &res, nil
}

func (s *ShortenerServer) GetStat(ctx context.Context, in *GetStatReq) (*GetStatResp, error) {
	var res GetStatResp

	p, _ := peer.FromContext(ctx)
	if logic.InWhiteList(p.Addr.String(), s.cfg.Subnet) {
		users, urls, err := logic.GetStat(s.storage)
		if err != nil {
			res.Error = err.Error()
		} else {
			res.Users = int32(users)
			res.Urls = int32(urls)
		}

	} else {
		res.Error = "Forbidden"
	}

	return &res, nil
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var user string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("user")
		if len(values) > 0 {
			user = values[0]
		}
	}

	// если нет метадаты, то создаем нового пользователя и добавляем метаданные
	// по факту принцип, как с cookie у http
	if len(user) == 0 {
		md := metadata.New(map[string]string{"user": uuid.New().String()}) // очень плохо, и нужно бы вынести в отдельную функцию генерацию нового ID пользователя
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	return handler(ctx, req)
}

func GetServer(cfg models.Config, ch4delete chan []string, store models.Storable) (*grpc.Server, error) {
	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))

	var us ShortenerServer
	us.Setup(store, ch4delete, cfg)

	RegisterUrlsServer(s, &us)

	return s, nil
}
