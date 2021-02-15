package friendlyamehandler

import "github.com/dhcgn/gohomematicmqttplugin/shared"

type handler struct {
}

type Handle interface {
	ExtendList(e shared.Event)
	AdjustEvent(e shared.Event) shared.Event
}

func New() Handle {
	return &handler{}
}

func (h handler) ExtendList(e shared.Event) {

}

func (h handler) AdjustEvent(e shared.Event) shared.Event {
	return e
}
