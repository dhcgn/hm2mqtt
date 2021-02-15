package userConfigHttpServer

import "testing"

func Test_createTemplate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTemplate(); got != tt.want {
				t.Errorf("createTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartWebService(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartWebService()
		})
	}
}
