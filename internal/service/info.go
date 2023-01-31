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

func NewInfoServer(conn db.DB) pb.InfoServer {
	return &infoServer{
		db: conn,
	}
}
func (s *infoServer) Checks(ctx context.Context, _ *pb.Empty) (list *pb.CheckList, err error) {
	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:

			var checks []*db.Check
			checks, err = s.db.GetCheckList()
			if err != nil {
				err = status.Error(codes.Internal, err.Error())
				fmt.Printf("%v", err)
				return
			}

			var grpcChecks []*pb.Check = []*pb.Check{}
			for _, chk := range checks {
				fmt.Printf("%v", chk)
				grpcCheck := new(pb.Check)
				grpcCheck.Id = chk.Id
				grpcCheck.Url = chk.Url
				grpcCheck.Type = pb.Check_CheckType(pb.Check_CheckType_value[chk.Type])
				grpcCheck.Comment = chk.Comment
				grpcChecks = append(grpcChecks, grpcCheck)
			}

			list = new(pb.CheckList)
			list.Checks = grpcChecks
			fmt.Printf("extraction checks was complite: %v", list.Checks)
			return
		}
	}
	return
}
