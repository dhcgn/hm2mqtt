package hmeventhandler

import (
	friendlyname "github.com/dhcgn/gohomematicmqttplugin/friendlyamehandler"
	"github.com/dhcgn/gohomematicmqttplugin/mqttHandler"
)

//HandlingIncomingEventsLoop parse incoming messages from chan to Events and send them via mqtt to the broker
func HandlingIncomingEventsLoop(messages <-chan string, mqttHandler mqttHandler.Handle, friendlyNameHandler friendlyname.Handle) {
	for {
		stringBody := <-messages
		var events, err = parseEventMultiCall(stringBody)
		if err != nil {
			continue
		}

		for _, e := range events {
			friendlyNameHandler.ExtendList(e)
			e = friendlyNameHandler.AdjustEvent(e)

			mqttHandler.SendToBroker(e)
		}
	}
}
