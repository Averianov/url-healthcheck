package test

import (
	"os"
	"os/exec"
	"testing"
)

func TestService(t *testing.T) {
	//	cmdExec(t, "docker-compose", "up", "--build", "-d", "disp")
	cmdExec(t, "docker-compose", "up", "--build", "-d")

	// some test

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
