package service

import (
	"url-healthcheck/pkg/db"
	pb "url-healthcheck/pkg/grpc"
)

func dbCheckToGrpcCheck(chk db.Check) (grpcCheck *pb.Check, err error) {
	grpcCheck = new(pb.Check)
	grpcCheck.Id = int64(chk.ID)
	grpcCheck.Url = chk.Url
	grpcCheck.Type = pb.Check_CheckType(pb.Check_CheckType_value[chk.Type])
	grpcCheck.Status = pb.Check_CheckStatus(pb.Check_CheckStatus_value[chk.Status])
	grpcCheck.Comment = chk.Comment
	return
}
