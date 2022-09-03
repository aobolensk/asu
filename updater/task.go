package asu

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func ProcessTask(task UpdaterTask) {
	var wg sync.WaitGroup
	log.Println("task: " + task.Name + " (api: " + task.APIVersion + ")")
	go StartTask(task)
	SetupStopTask(task, &wg)
	wg.Wait()
}

func StartTask(task UpdaterTask) {
	RunProcess(task.OnStart.Env, task.OnStart.Script)
}

func SetupStopTask(task UpdaterTask, wg *sync.WaitGroup) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)
	wg.Add(1)

	go func() {
		<-sigChan
		defer wg.Done()
		RunProcess(task.OnStop.Env, task.OnStop.Script)
	}()
}
