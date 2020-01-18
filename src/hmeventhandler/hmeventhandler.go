package hmeventhandler

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

var (
	// TODO Set Last Will
	opts = mqtt.NewClientOptions().AddBroker("tcp://192.168.10.31:1883").SetClientID("HomeMaticMqtt").SetAutoReconnect(true)
	c    = mqtt.NewClient(opts)
)

func UploadLoop(messages <-chan string) {
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		stringBody := <-messages
		var events, err = parseEventMultiCall(stringBody)
		if err != nil {
			continue
		}

		for _, e := range events {
			sendToBroker(e)
		}
	}
}

func sendToBroker(e Event) {
	topic := "hm/" + e.SerialNumber + "/" + e.Type

	start := time.Now()

	token := c.Publish(topic, 1, false, e.DataValue)
	wait := token.WaitTimeout(2 * time.Second)
	err := token.Error()

	elapsed := time.Since(start)

	if wait && err == nil {
		log.Println("OK:    topic:", topic, "with value: ", e.DataValue, elapsed)
	} else {
		log.Println("ERROR: topic:", topic, "with value: ", e.DataValue, "Wait", wait, "Error:", err, elapsed)
		time.Sleep(500 * time.Millisecond)
	}
}
