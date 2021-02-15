package mqtthandler

import (
	"reflect"
	"testing"

	"github.com/dhcgn/hm2mqtt/shared"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestNew(t *testing.T) {
	type args struct {
		config  *shared.Configuration
		handler mqtt.MessageHandler
	}
	tests := []struct {
		name string
		args args
		want Handle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.config, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handle_Disconnect(t *testing.T) {
	tests := []struct {
		name string
		h    handle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Disconnect()
		})
	}
}

func Test_handle_SendToBroker(t *testing.T) {
	type args struct {
		e shared.Event
	}
	tests := []struct {
		name string
		h    handle
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.SendToBroker(tt.args.e)
		})
	}
}
