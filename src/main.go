package main

import (
	"flag"
	"fmt"
	"github.com/dhcgn/gohomematicmqttplugin/src/hmclient"
	"github.com/dhcgn/gohomematicmqttplugin/src/hmeventhandler"
	"github.com/dhcgn/gohomematicmqttplugin/src/hmlistener"
	"github.com/dhcgn/gohomematicmqttplugin/src/userConfigHttpServer"
	"github.com/dhcgn/gohomematicmqttplugin/src/mqttHandler"
	"github.com/dhcgn/gohomematicmqttplugin/src/shared"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	version = "undef"
	flagTokenPtr   = flag.String("config", ``, `Overrides the path to the config file`)
)

const (
	applicationName = "GoHomeMaticMqttPlugin"
)

func main() {
	fmt.Println(applicationName)
	fmt.Println("Version:", version)
	fmt.Println("Project URL: https://github.com/dhcgn/GoHomeMaticMqttPlugin ")

	flag.Parse()

	config := shared.ReadConfig(*flagTokenPtr)

	events := make(chan string, 1000)
	ticker := time.NewTicker(1 * time.Minute)
	tickerStatus := time.NewTicker(1 * time.Second)

	mqttHandler.Init(config)

	go func() { hmeventhandler.UploadLoop(events) }()
	go func() { hmlistener.StartServer(events, config.ListenerPort) }()
	go func() { syncLoop(ticker.C, config) }()
	go func() { statsLoop(tickerStatus.C, events) }()
	go func() { userConfigHttpServer.Start() }()

	c := make(chan os.Signal)
	cleanupDone := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()
	<-cleanupDone
}

func statsLoop(tick <-chan time.Time, events chan string) {
	for range tick {
		eventCount := len(events)
		if eventCount != 0 {
			log.Println("Events: ", eventCount)
		}
	}
}

func syncLoop(tick <-chan time.Time, config *shared.Configuration) {
	if runtime.GOOS == "windows" {
		log.Println("Skipped on windows")
		return
	}

	// devices <- client.GetDevices()
	hmclient.Init(config.ListenerPort, config.InterfaceId, config.HomematicUrl)

	for range tick {
		hmclient.Init(config.ListenerPort, config.InterfaceId, config.HomematicUrl)
		// devices <- client.GetDevices()
	}
}

func cleanup() {
	log.Println("Starting Cleanup")

	mqttHandler.Disconnect()
}
