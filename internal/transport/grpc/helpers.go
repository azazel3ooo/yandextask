package grpc

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func getUser(ctx context.Context) string {
	var user string

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md.Get("user")
		if len(values) > 0 {
			user = values[0]
		}
	}
	return user
}
