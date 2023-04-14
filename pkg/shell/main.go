package shell

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func getCurrentShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	return shell
}

func getShell() string {
	os := runtime.GOOS
	if os == "windows" {
		return "powershell"
	}

	return getCurrentShell()
}

var shellToUse = getShell()

func Out(command string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(shellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Start()
	err := cmd.Wait()

	return stdout.String(), stderr.String(), err
}

func OutLive(command string) (string, string, error) {
	cmd := exec.Command(shellToUse, "-c", command)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	outStr, errStr := stdoutBuf.String(), stderrBuf.String()

	if err != nil {
		return outStr, errStr, err
	}

	return outStr, errStr, nil
}

/*Execute a command and return the output*/
func ExecuteCommand(cmd string, print ...bool) string {
	var ignore bool
	out, errout, err := Out(cmd)

	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)

		log.Fatal("")
	}

	if len(print) > 0 && !print[0] {
		ignore = true
	}

	if !ignore {
		fmt.Println(out)
		fmt.Println(errout)
	}

	return out
}