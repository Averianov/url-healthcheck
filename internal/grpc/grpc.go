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

func StartGRPCServer(
	db db.DB,
) (err error) {
	var srv *grpc.Server

	srv = grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor()),
	)

	infoSrv := service.NewInfoServer(db)
	pb.RegisterInfoServer(srv, infoSrv)
	reflection.Register(srv)

	var lis net.Listener
	lis, err = net.Listen("tcp", ":443")
	if err != nil {
		err = fmt.Errorf("failed to listen: %v", err)
		return
	}
	return srv.Serve(lis)

}
