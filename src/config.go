package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	ListenerPort int
	InterfaceId  int
	HomematicUrl string
}

func ReadConfig() *config {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	configPath := filepath.Join(dir, "config.json")
	configSamplePath := filepath.Join(dir, "config.sample.json")

	if _, err := os.Stat(configSamplePath); os.IsNotExist(err) {
		newConfig := config{
			ListenerPort: 8777,
			InterfaceId:  2,
			HomematicUrl: "http://192.168.1.100:2001/",
		}
		f, _ := os.Create(configSamplePath)
		defer f.Close()
		j, _ := json.Marshal(newConfig)
		f.Write(j)
		f.Sync()
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("No config at %v", configPath)
	}

	dat, _ := ioutil.ReadFile(configPath)
	c := config{}
	json.Unmarshal(dat, &c)

	return &c
}
