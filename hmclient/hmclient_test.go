package hmclient

import "testing"

func TestGetDevices(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDevices(); got != tt.want {
				t.Errorf("GetDevices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetState(t *testing.T) {
	type args struct {
		address      string
		valueKey     string
		value        string
		homematicURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetState(tt.args.address, tt.args.valueKey, tt.args.value, tt.args.homematicURL); got != tt.want {
				t.Errorf("SetState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInit(t *testing.T) {
	type args struct {
		port         int
		interfaceID  int
		homematicURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Init(tt.args.port, tt.args.interfaceID, tt.args.homematicURL); got != tt.want {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}
