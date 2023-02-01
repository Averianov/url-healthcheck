package grpc

import (
	"fmt"
	"net"
	"url-healthcheck/internal/service"
	"url-healthcheck/pkg/db"
	pb "url-healthcheck/pkg/grpc"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"
)

// StartGRPCServer launch grpc api
func StartGRPCServer(
	db db.DB,
	port string,
) (err error) {
	var srv *grpc.Server

	srv = grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor()),
	)

	infoSrv := service.NewInfoServer(db)
	pb.RegisterInfoServer(srv, infoSrv)
	reflection.Register(srv)

	var lis net.Listener
	lis, err = net.Listen("tcp", ":"+port)
	if err != nil {
		err = fmt.Errorf("failed to listen: %v", err)
		return
	}
	return srv.Serve(lis)

}
