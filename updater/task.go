package asu

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func ProcessTask(task UpdaterTask) {
	log.Println("task: " + task.Name + " (api: " + task.APIVersion + ")")
	go StartTask(task)
	for {
		time.Sleep(10 * time.Second)
	}
}

func StartTask(task UpdaterTask) {
	cmd := exec.Command("bash", "-")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, task.OnStart.Env...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(stdin, task.OnStart.Script)
	if err != nil {
		log.Fatal(err)
	}
	stdin.Close()
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}
