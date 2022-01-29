package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var flagChdir = flag.String("C", "", "Set current working directory")

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
		fmt.Fprintln(os.Stderr, "Usage: zenv {options..} {ENVNAME=VALUE...} COMMAND {ARGS...}")
		flag.PrintDefaults()
		return 0, nil
	}

	if *flagDllDir != "" {
		if err := addDllDirectories(strings.Split(*flagDllDir, string(os.PathListSeparator))...); err != nil {
			return 1, nil
		}
	}

	if *flagChdir != "" {
		if err := os.Chdir(*flagChdir); err != nil {
			return 1, err
		}
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
	flag.Parse()
	rc, err := mains(flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	os.Exit(rc)
}
