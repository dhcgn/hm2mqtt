package cmdhandler

import (
	"log"
	"strings"

	"github.com/dhcgn/gohomematicmqttplugin/hmclient"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type CmdHandler interface {
	AddCmd(msg mqtt.Message)
}

type cmdHandler struct {
	homematicUrl string
}

func NewCmdHandler(homematicUrl string) CmdHandler {
	return &cmdHandler{
		homematicUrl: homematicUrl,
	}
}

func (c *cmdHandler) AddCmd(msg mqtt.Message) {
	log.Println("cmd receiver got msg", msg.Topic(), string(msg.Payload()))

	segments := strings.Split(msg.Topic(), "/")
	if len(segments) != 5 {
		log.Println("cmd receiver got invalid msg, should be e.g. hm/set/JEQ0000000:1/Level/ with value: 1")
	}

	valueKey := segments[len(segments)-2]
	address := segments[len(segments)-3]
	value := string(msg.Payload())

	hmclient.SetState(address, valueKey, value, c.homematicUrl)
}
