package shared

import (
	"reflect"
	"testing"
)

func TestUpdateConfiguration(t *testing.T) {
	type args struct {
		c Configuration
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateConfiguration(tt.args.c)
		})
	}
}

func TestReadConfig(t *testing.T) {
	type args struct {
		overriddenPath string
	}
	tests := []struct {
		name string
		args args
		want *Configuration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadConfig(tt.args.overriddenPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
