package test

import (
	"context"
	"io"
	"os"
	"os/exec"
	"testing"
	"time"
	"url-healthcheck/config"

	pb "url-healthcheck/pkg/grpc"

	"cloud.google.com/go/logging"
	"google.golang.org/grpc"
)

func TestService(t *testing.T) {
	//	cmdExec(t, "docker-compose", "up", "--build", "-d", "disp")
	cmdExec(t, "docker-compose", "up", "--build", "-d")

	conn, err := grpc.Dial("localhost:"+config.GRPCPort, grpc.WithInsecure())
	if err != nil {
		logging.Default.Fatalln(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), 6*time.Hour)
	defer cancel()

	if err = launchClient(ctx, conn); err != nil {
		logging.Default.Fatalln(err)
	}

	defer func() {
		cmdExec(t, "docker-compose", "down")
	}()
}

func cmdExec(t *testing.T, args ...string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		t.Fatalf("cannot %s due to %s", args, err.Error())
	}
	t.Logf("success %s for the integration test\n", args)
}

func launchClient(ctx context.Context, conn *grpc.ClientConn) (err error) {
	infoClient := NewInfoClient(conn)
	infoClient.Open(ctx) // нужно добавить канал для останова стрима

	return
}

type infoClient struct {
	client pb.InfoClient
}

func NewInfoClient(conn grpc.ClientConnInterface) infoClient {
	return infoClient{
		client: pb.NewInfoClient(conn),
	}
}

func (hbc heartBeatClient) Open(ctx context.Context) {
	//func (hbc heartBeatClient) Open(ctx context.Context) (err error) {
	var err error
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var requests []*pb.HeartBeatRequest
	var incomingJSON []string
	// for i := 0; i < 5; i++ { // make buffer from 5 json items
	// 	incomingJSON = append(incomingJSON, `{"partner_id":"037b7e17-68a2-11ed-94a5-525400112eb0",
	// 	"device_id":"123456786","sim_card_imsi":"df789qwe321",
	// 	"battery":75,"comment":"testing incident 3","roaming":false}`)
	// }

	requests, err = makeRequests(incomingJSON) // make pb.HeartBeatRequest from json items
	if err != nil {
		logging.Default.Errorf("make requests: %v", err)
		return
	}

	var stream pb.HeartBeat_BeatClient
	stream, err = hbc.client.Beat(ctx)
	if err != nil {
		logging.Default.Errorf("create stream: %v", err)
		return
	}

	var response *pb.Response

	// ticker := time.NewTicker(5 * time.Minute)
	// for _ := range ticker.C {
	// 	time.Sleep(60 * time.Second)
	// 	logging.Default.Infof("%v", "send reaquest")
	// 	if err = stream.Send(requests[0]); err != nil { // send pb.HeartBeatRequest to server
	// 		logging.Default.Errorf("send stream: %v", err)
	// 		return
	// 	}
	// }
	//time.Sleep(300 * time.Second)

	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {

		logging.Default.Infof("send request: \n%v\n\n_", requests[0])
		err = stream.Send(requests[0])
		if err != nil { // send pb.HeartBeatRequest to server
			logging.Default.Debugf("%v", err)
			if err == io.EOF {
				err = stream.CloseSend()
				if err != nil {
					logging.Default.Errorf("close and receive: %v", err)
				}
				return
			}
		}

		response, err = stream.Recv()
		if err != nil { // send pb.HeartBeatRequest to server
			logging.Default.Debugf("%v", err)
			if err == io.EOF {
				err = stream.CloseSend()
				if err != nil {
					logging.Default.Errorf("close and receive: %v", err)
				}
				return
			}
		}
		logging.Default.Infof("%v", response)

		// err = stream.Send(requests[0])
		// if err != nil { // send pb.HeartBeatRequest to server
		// 	logging.Default.Debugf("%v", err)
		// 	if err == io.EOF {
		// 		response, err = stream.CloseAndRecv()
		// 		if err != nil {
		// 			logging.Default.Errorf("close and receive: %v", err)
		// 			return
		// 		}
		// 	}
		// 	logging.Default.Errorf("error stream: %v", err)
		// 	return
		// }
	}

	// response, err = stream.CloseAndRecv()
	// if err != nil {
	// 	logging.Default.Errorf("close and receive: %v", err)
	// 	return
	// }
	return
}
