package hmlistener

import "testing"

func TestStartServer(t *testing.T) {
	type args struct {
		messages chan<- string
		port     int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartServer(tt.args.messages, tt.args.port)
		})
	}
}
