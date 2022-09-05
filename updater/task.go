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
	go RunProcess(task.OnStart.Env, task.OnStart.Script, task.Directory)
}

func SetupStopTask(task UpdaterTask, wg *sync.WaitGroup) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)
	wg.Add(1)

	go func() {
		<-sigChan
		defer wg.Done()
		log.Println("Stop")
		RunProcess(task.OnStop.Env, task.OnStop.Script, task.Directory)
	}()
}

func StartUpdateRoutine(task UpdaterTask) {
	for {
		time.Sleep(time.Duration(task.Update.Interval) * time.Second)
		log.Println("Before update")
		ret := RunProcess(task.Update.Before.Env, task.Update.Before.Script, task.Directory)
		if ret != 0 {
			log.Println("Before update failed. Skipping update cycle...")
			continue
		}
		log.Println("Update")
		RunProcess(task.Update.On.Env, task.Update.On.Script, task.Directory)
		log.Println("After update")
		go RunProcess(task.Update.After.Env, task.Update.After.Script, task.Directory)
		log.Println("Update cycle is done. Waiting for new one...")
	}
}
