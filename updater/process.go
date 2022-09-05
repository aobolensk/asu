package asu

import (
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func RunProcess(env []string, script string, workingDir string) int {
	cmd := exec.Command("bash", "-")
	cmd.Dir = workingDir
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, env...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(stdin, script)
	if err != nil {
		log.Fatal(err)
	}
	stdin.Close()
	cmd.Start()
	cmd.Wait()
	retCode := cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	return retCode
}
