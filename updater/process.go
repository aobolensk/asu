package asu

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func RunProcess(env []string, script string) int {
	cmd := exec.Command("bash", "-")
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
	out, _ := cmd.CombinedOutput()
	retCode := cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	fmt.Printf("%s\n", out)
	return retCode
}
