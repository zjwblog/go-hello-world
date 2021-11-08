package util

import (
	"bytes"
	"os/exec"
)

func ExecLinuxShell(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	var result bytes.Buffer
	cmd.Stdout = &result
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return result.String(), err
}
