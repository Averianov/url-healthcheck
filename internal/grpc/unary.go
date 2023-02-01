package grpc

import (
	"context"

	"google.golang.org/grpc"
)

// UnaryInterceptor do preparing incoming requests
func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (i interface{}, err error) {
		defer func(err error) {
			if err != nil {
				//fmt.Println(err)
			}
		}(err)

		// there will be checked grpc authentication

		return handler(ctx, req)
	}
}
