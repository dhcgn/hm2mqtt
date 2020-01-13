package main

import (
	"fmt"
	"github.com/dhcgn/gohomematicmqttplugin/src/hmclient"
	"github.com/dhcgn/gohomematicmqttplugin/src/hmeventhandler"
	"github.com/dhcgn/gohomematicmqttplugin/src/hmlistener"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	version = "undef"
)

func main() {
	fmt.Println("Version:", version)

	log.Println("Starting")

	config := ReadConfig()
	log.Println("ReadConfig: ", config)

	events := make(chan string, 1000)
	ticker := time.NewTicker(1 * time.Minute)
	tickerStatus := time.NewTicker(1 * time.Second)

	go func() { hmeventhandler.UploadLoop(events) }()
	go func() { hmlistener.StartServer(events, config.ListenerPort) }()
	go func() { SyncLoop(ticker.C, config) }()
	go func() { StatsLoop(tickerStatus.C, events) }()

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

func StatsLoop(tick <-chan time.Time, events chan string) {
	for range tick {
		log.Println("Events: ", len(events))
	}
}

func SyncLoop(tick <-chan time.Time, config *config) {
	// Init
	log.Print("SyncLoop Init")

	if runtime.GOOS == "windows" {
		log.Println("Skipped on windows")
		return
	}

	// devices <- client.GetDevices()
	hmclient.Init(config.ListenerPort, config.InterfaceId, config.HomematicUrl)

	for range tick {
		log.Print("SyncLoop Started")
		hmclient.Init(config.ListenerPort, config.InterfaceId, config.HomematicUrl)
		// devices <- client.GetDevices()
	}
}

func cleanup() {
	log.Println("Starting Cleanup")

	// TODO unsubscript!
}
