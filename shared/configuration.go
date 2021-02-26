package shared

// TODO with Go 1.16 should use FS for testing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Config is the instance of the actual Configuration
var Config Configuration

// ConfigFilePath is the path to the config file
var ConfigFilePath string

// Configuration is the config which is necessarily to run hm2mqtt
type Configuration struct {
	// ListenerPort is the own port which hm2mqtt is listening to
	ListenerPort int
	// InterfaceID is an ID used for RPC calls
	InterfaceID int
	// HomematicURL the url points to homematic rpc
	HomematicURL string
	// BrokerURL the url points to a mqtt broker
	BrokerURL string
	// Retain mqtt messages
	Retain bool
}

// UpdateConfiguration save new config to disk
func UpdateConfiguration(c Configuration) {
	Config = c

	f, _ := os.Create(ConfigFilePath)
	defer f.Close()
	j, _ := json.MarshalIndent(c, "", "   ")
	f.Write(j)
	f.Sync()
}

// ReadConfig read config from disk
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
			InterfaceID:  2,
			HomematicURL: "http://127.0.0.1:2001/",
			BrokerURL:    "tcp://192.168.10.31:1883",
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
