package runner

import (
	"bytes"
	"os/exec"
)

func RunCmdOnHost(command string, host string) (*bytes.Buffer, error) {
	cmd := exec.Command("ssh", host, command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return &out, nil
}
