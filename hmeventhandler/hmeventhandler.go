package hmeventhandler

import (
	friendlyname "github.com/dhcgn/hm2mqtt/friendlyamehandler"
	mqtthandler "github.com/dhcgn/hm2mqtt/mqtthandler"
)

//HandlingIncomingEventsLoop parse incoming messages from chan to Events and send them via mqtt to the broker
func HandlingIncomingEventsLoop(messages <-chan string, mqttHandler mqtthandler.Handle, friendlyNameHandler friendlyname.Handle) {
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
