package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/dhcgn/gohomematicmqttplugin/cmdhandler"

	"github.com/dhcgn/gohomematicmqttplugin/hmclient"
	"github.com/dhcgn/gohomematicmqttplugin/hmeventhandler"
	"github.com/dhcgn/gohomematicmqttplugin/hmlistener"
	"github.com/dhcgn/gohomematicmqttplugin/mqttHandler"
	"github.com/dhcgn/gohomematicmqttplugin/shared"
	"github.com/dhcgn/gohomematicmqttplugin/userConfigHttpServer"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	version      = "undef"
	flagTokenPtr = flag.String("config", ``, `Overrides the path to the config file`)
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

	cmd := cmdhandler.NewCmdHandler()

	events := make(chan string, 1000)
	tickerRefreshSubscription := time.NewTicker(1 * time.Minute)
	tickerStatus := time.NewTicker(1 * time.Second)

	mqttHandler.Init(config, func(client mqtt.Client, msg mqtt.Message) { cmd.AddCmd(msg) })

	go func() { hmeventhandler.UploadLoop(events) }()
	go func() { hmlistener.StartServer(events, config.ListenerPort) }()
	go func() { refreshSubscriptionLoop(tickerRefreshSubscription.C, config) }()
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
			log.Println("Events queued: ", eventCount)
		}
	}
}

func refreshSubscriptionLoop(tick <-chan time.Time, config *shared.Configuration) {
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
