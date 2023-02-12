package test

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

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

func TestService(t *testing.T) {
	//	cmdExec(t, "docker-compose", "up", "--build", "-d", "disp")
	cmdExec(t, "docker-compose", "up", "--build", "-d")
	defer func() {
		cmdExec(t, "docker-compose", "down")
	}()

	t.Logf("start grpc client with %s\n", "172.20.0.3:443")
	conn, err := grpc.Dial("172.20.0.3:443", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), 6*time.Minute)
	defer cancel()

	infoClient := NewInfoClient(conn)

	var checkList *pb.CheckList
	t.Logf("start first check\n")
	for { // wait for available server api
		checkList, err = infoClient.client.Checks(ctx, &pb.Empty{})
		if err != nil {
			t.Logf("%v", err)
			time.Sleep(1 * time.Minute)
			continue
		} else {
			t.Logf("len checks %d\n", len(checkList.Checks))
			if len(checkList.Checks) != 0 {
				err = fmt.Errorf("db is not clear")
				t.Fatalf("%v", err)
				return
			}
			break
		}
	}

	time.Sleep(4 * time.Minute)

	t.Logf("start second check\n")
	checkList, err = infoClient.client.Checks(ctx, &pb.Empty{})
	if err != nil {
		t.Fatalf("%v", err)
		return
	}
	if len(checkList.Checks) == 0 {
		err = fmt.Errorf("db must be setted")
		t.Fatalf("%v", err)
		return
	}
	t.Logf("len checks %d\n", len(checkList.Checks))
	//t.Fatalf("%v", checkList.Checks)
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
