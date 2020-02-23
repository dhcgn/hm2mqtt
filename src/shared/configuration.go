package shared

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Configuration struct {
	ListenerPort int
	InterfaceId  int
	HomematicUrl string
	BrokerUrl    string
}

func ReadConfig(overriddenPath string) *Configuration {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	configPath := filepath.Join(dir, "config.json")
	configSamplePath := filepath.Join(dir, "config.sample.json")

	if overriddenPath != "" {
		configPath = overriddenPath
		configSamplePath = overriddenPath
	}

	if _, err := os.Stat(configSamplePath); os.IsNotExist(err) {
		newConfig := Configuration{
			ListenerPort: 8777,
			InterfaceId:  2,
			HomematicUrl: "http://127.0.0.1:2001/",
			BrokerUrl:    "tcp://192.168.10.31:1883",
		}
		f, _ := os.Create(configSamplePath)
		defer f.Close()
		j, _ := json.MarshalIndent(newConfig, "", "   ")
		f.Write(j)
		f.Sync()
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("No config at %v", configPath)
	}

	dat, _ := ioutil.ReadFile(configPath)
	c := Configuration{}
	json.Unmarshal(dat, &c)

	return &c
}
