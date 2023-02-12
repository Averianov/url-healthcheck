package main

import (
	"context"
	"log"
	"time"
	"url-healthcheck/config"
	pb "url-healthcheck/pkg/grpc"

	"google.golang.org/grpc"
)

type infoClient struct {
	client pb.InfoClient
}

func NewInfoClient(conn grpc.ClientConnInterface) infoClient {
	return infoClient{
		client: pb.NewInfoClient(conn),
	}
}

func main() {
	var err error

	config.LoadConfig()

	conn, err := grpc.Dial("localhost:"+config.GRPCPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), 6*time.Minute)
	defer cancel()

	infoClient := NewInfoClient(conn)

	log.Println("Start check")
	var checkList *pb.CheckList
	for {
		checkList, err = infoClient.client.Checks(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("%v", err)
			return
		}
		log.Printf("len checks %d\n", len(checkList.Checks))

		time.Sleep(1 * time.Minute)
	}
}
