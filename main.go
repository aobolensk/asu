package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	asu "github.com/aobolensk/asu/updater"
	"gopkg.in/yaml.v3"
)

func main() {
	fmt.Println("asu")
	var dir string
	flag.StringVar(&dir, "dir", ".", "directory where asu.yaml file should be found")
	flag.Parse()
	cfg := processDirectory(dir)
	asu.ProcessTask(cfg.Task)
}

func processDirectory(dir string) asu.Config {
	yamlBytes, err := os.ReadFile(filepath.Join(dir, "asu.yaml"))
	if err != nil {
		log.Fatal(err)
	}
	config := asu.Config{}
	err = yaml.Unmarshal(yamlBytes, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
