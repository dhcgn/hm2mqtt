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

	"github.com/dhcgn/hm2mqtt/cmdhandler"
	"github.com/dhcgn/hm2mqtt/friendlyamehandler"

	"github.com/dhcgn/hm2mqtt/hmclient"
	"github.com/dhcgn/hm2mqtt/hmeventhandler"
	"github.com/dhcgn/hm2mqtt/hmlistener"
	"github.com/dhcgn/hm2mqtt/mqttHandler"
	"github.com/dhcgn/hm2mqtt/shared"
	"github.com/dhcgn/hm2mqtt/userConfigHttpServer"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	version      = "undef"
	flagTokenPtr = flag.String("config", ``, `Overrides the path to the config file`)
)

const (
	applicationName = "hm2mqtt"
)

func main() {
	fmt.Println(applicationName)
	fmt.Println("Version:", version)
	fmt.Println("Project URL: https://github.com/dhcgn/hm2mqtt ")

	flag.Parse()

	config := shared.ReadConfig(*flagTokenPtr)

	cmd := cmdhandler.NewCmdHandler(config.HomematicUrl)
	friendlyName := friendlyamehandler.New()

	events := make(chan string, 1000)
	tickerRefreshSubscription := time.NewTicker(1 * time.Minute)
	tickerStatus := time.NewTicker(1 * time.Second)

	mqttHandler := mqttHandler.New(config, func(client mqtt.Client, msg mqtt.Message) { cmd.SendNewStateToHomematic(msg) })

	go func() { hmeventhandler.HandlingIncomingEventsLoop(events, mqttHandler, friendlyName) }()
	go func() { hmlistener.StartServer(events, config.ListenerPort) }()
	go func() { refreshSubscriptionLoop(tickerRefreshSubscription.C, config) }()
	go func() { statsLoop(tickerStatus.C, events) }()
	go func() { userConfigHttpServer.StartWebService() }()

	c := make(chan os.Signal)
	cleanupDone := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup(mqttHandler)
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

func cleanup(mqttHandler mqttHandler.Handle) {
	log.Println("Starting Cleanup")

	mqttHandler.Disconnect()
}
