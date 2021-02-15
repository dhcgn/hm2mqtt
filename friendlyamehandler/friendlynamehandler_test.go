package friendlyamehandler

import (
	"reflect"
	"testing"

	"github.com/dhcgn/hm2mqtt/shared"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Handle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler_ExtendList(t *testing.T) {
	type args struct {
		e shared.Event
	}
	tests := []struct {
		name string
		h    handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.ExtendList(tt.args.e)
		})
	}
}

func Test_handler_AdjustEvent(t *testing.T) {
	type args struct {
		e shared.Event
	}
	tests := []struct {
		name string
		h    handler
		args args
		want shared.Event
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.AdjustEvent(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handler.AdjustEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
