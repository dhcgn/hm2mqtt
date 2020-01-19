package hmeventhandler

import (
	"github.com/dhcgn/gohomematicmqttplugin/src/mqttHandler"
)

//UploadLoop parse incoming messages from chan to Events and end them via mqtt to the broker
func UploadLoop(messages <-chan string) {
	for {
		stringBody := <-messages
		var events, err = parseEventMultiCall(stringBody)
		if err != nil {
			continue
		}

		for _, e := range events {
			mqttHandler.SendToBroker(e)
		}
	}
}
