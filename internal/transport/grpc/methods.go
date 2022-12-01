package grpc

import "github.com/azazel3ooo/yandextask/internal/models"

// FromUserResponseSlice Заполняет UserUrlsResponse.Result с помощью []models.UserResponse
func (u *UserUrlsResponse) FromUserResponseSlice(s []models.UserResponse) {
	for _, el := range s {
		cur := UserUrlsResponse_UrlsPair{Original: el.Original, Short: el.Short}
		u.Result = append(u.Result, &cur)
	}
}

// ToSet конвертирует SetManyRequest.Set в []models.CustomIDSet
func (sReq *SetManyRequest) ToSet() []models.CustomIDSet {
	res := make([]models.CustomIDSet, len(sReq.Set))

	for idx, el := range sReq.Set {
		cur := models.CustomIDSet{OriginalURL: el.Original, CorrelationID: el.Id}
		res[idx] = cur
	}
	return res
}

// FromSet заполняет SetManyResponse.Set с помощью []models.CustomIDSet
func (sResp *SetManyResponse) FromSet(s []models.CustomIDSet) {
	for _, el := range s {
		cur := SetManyResponse_CustomIDSetResp{Original: el.OriginalURL, Short: el.ShortURL, Id: el.CorrelationID}
		sResp.Set = append(sResp.Set, &cur)
	}
}

func (s *ShortenerServer) Setup(store models.Storable, ch chan []string, cfg models.Config) {
	s.storage = store
	s.chanForDelete = ch
	s.cfg = cfg
}
