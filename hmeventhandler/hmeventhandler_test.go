package hmeventhandler

import (
	"testing"

	friendlyname "github.com/dhcgn/hm2mqtt/friendlyamehandler"
	"github.com/dhcgn/hm2mqtt/mqtthandler"
)

func TestHandlingIncomingEventsLoop(t *testing.T) {
	type args struct {
		messages            <-chan string
		mqtthandler         mqtthandler.Handle
		friendlyNameHandler friendlyname.Handle
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandlingIncomingEventsLoop(tt.args.messages, tt.args.mqtthandler, tt.args.friendlyNameHandler)
		})
	}
}
