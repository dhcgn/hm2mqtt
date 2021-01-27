package cmdhandler

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type CmdHandler interface {
	AddCmd(msg mqtt.Message)
}

type cmdHandler struct {
}

func NewCmdHandler() CmdHandler {
	return &cmdHandler{}
}

func (c *cmdHandler) AddCmd(msg mqtt.Message) {
	log.Println("cmd receiver got msg", msg.Topic(), string(msg.Payload()))
}
