package service

import (
	"context"
	"fmt"
	"url-healthcheck/pkg/db"

	pb "url-healthcheck/pkg/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type infoServer struct {
	db db.DB
	pb.UnimplementedInfoServer
}

// NewInfoServer create grpc service InfoServer
func NewInfoServer(conn db.DB) pb.InfoServer {
	return &infoServer{
		db: conn,
	}
}

// Checks get and returned checks list from db
func (s *infoServer) Checks(ctx context.Context, _ *pb.Empty) (list *pb.CheckList, err error) {
	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
			fmt.Println("Incoming grpc request Checks")
			var checks []db.Check
			checks, err = s.db.GetCheckList()
			if err != nil {
				err = status.Error(codes.Internal, err.Error())
				//fmt.Println(err)
				return
			}

			var grpcChecks []*pb.Check = []*pb.Check{}
			for _, chk := range checks {
				//fmt.Printf("%v", chk)
				grpcCheck := dbCheckToGrpcCheck(chk)
				grpcChecks = append(grpcChecks, grpcCheck)
			}

			list = new(pb.CheckList)
			list.Checks = grpcChecks
			//fmt.Printf("extraction checks was complite: %v", list.Checks)
			return
		}
	}
	return
}
