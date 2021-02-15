package cmdhandler

import (
	"reflect"
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestNewCmdHandler(t *testing.T) {
	type args struct {
		homematicUrl string
	}
	tests := []struct {
		name string
		args args
		want CmdHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCmdHandler(tt.args.homematicUrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCmdHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cmdHandler_SendNewStateToHomematic(t *testing.T) {
	type args struct {
		msg mqtt.Message
	}
	tests := []struct {
		name string
		c    *cmdHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.SendNewStateToHomematic(tt.args.msg)
		})
	}
}
