package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func mains(args []string) (int, error) {
	for len(args) > 0 {
		equalPos := strings.IndexByte(args[0], '=')
		if equalPos < 0 {
			break
		}
		os.Setenv(args[0][:equalPos], args[0][equalPos+1:])
		args = args[1:]
	}
	if len(args) <= 0 {
		return 0, nil
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	rc := -1
	if cmd.ProcessState != nil {
		rc = cmd.ProcessState.ExitCode()
	}
	return rc, err
}

func main() {
	rc, err := mains(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	os.Exit(rc)
}
