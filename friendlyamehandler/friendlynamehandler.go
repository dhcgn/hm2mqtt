package friendlyamehandler

import "github.com/dhcgn/hm2mqtt/shared"

type handler struct {
}

// Handle to iteract with the friendly names
type Handle interface {
	ExtendList(e shared.Event)
	AdjustEvent(e shared.Event) shared.Event
}

// New creates a Handle to manage the FriendlyNames for homematic devices
func New() Handle {
	return &handler{}
}

func (h handler) ExtendList(e shared.Event) {

}

func (h handler) AdjustEvent(e shared.Event) shared.Event {
	return e
}
