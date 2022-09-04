package asu

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func ProcessTask(task UpdaterTask) {
	var wg sync.WaitGroup
	log.Println("task: " + task.Name + " (api: " + task.APIVersion + ")")
	go StartTask(task)
	go StartUpdateRoutine(task)
	SetupStopTask(task, &wg)
	wg.Wait()
}

func StartTask(task UpdaterTask) {
	log.Println("Start")
	RunProcess(task.OnStart.Env, task.OnStart.Script)
}

func SetupStopTask(task UpdaterTask, wg *sync.WaitGroup) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)
	wg.Add(1)

	go func() {
		<-sigChan
		defer wg.Done()
		log.Println("Stop")
		RunProcess(task.OnStop.Env, task.OnStop.Script)
	}()
}

func StartUpdateRoutine(task UpdaterTask) {
	for {
		time.Sleep(time.Duration(task.Update.Interval) * time.Second)
		log.Println("Before update")
		ret := RunProcess(task.Update.Before.Env, task.Update.Before.Script)
		if ret != 0 {
			log.Println("Before update failed. Skipping update cycle...")
			continue
		}
		log.Println("Update")
		RunProcess(task.Update.On.Env, task.Update.On.Script)
		log.Println("After update")
		RunProcess(task.Update.After.Env, task.Update.After.Script)
		log.Println("Update cycle is done. Waiting for new one...")
	}
}
