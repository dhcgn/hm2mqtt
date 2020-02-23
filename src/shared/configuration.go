package shared

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var Config Configuration;
var ConfigFilePath string;

type Configuration struct {
	ListenerPort int
	InterfaceId  int
	HomematicUrl string
	BrokerUrl    string
}

func UpdateConfiguration(c Configuration) {
	Config = c

	f, _ := os.Create(ConfigFilePath)
	defer f.Close()
	j, _ := json.MarshalIndent(c, "", "   ")
	f.Write(j)
	f.Sync()
}

func ReadConfig(overriddenPath string) *Configuration {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	ConfigFilePath = filepath.Join(dir, "config.json")
	configSamplePath := filepath.Join(dir, "config.sample.json")

	if overriddenPath != "" {
		ConfigFilePath = overriddenPath
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

	if _, err := os.Stat(ConfigFilePath); os.IsNotExist(err) {
		log.Fatalf("No config at %v", ConfigFilePath)
	}

	dat, _ := ioutil.ReadFile(ConfigFilePath)
	c := Configuration{}
	json.Unmarshal(dat, &c)

	Config = c
	return &c
}
