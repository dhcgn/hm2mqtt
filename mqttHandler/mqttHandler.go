package mqttHandler

import (
	"log"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/dhcgn/gohomematicmqttplugin/shared"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// TOOO Change Impl to https://github.com/eclipse/paho.mqtt.golang/blob/43c9c445a89e7dca549a9bd445e3553dade9a2bc/cmd/docker/subscriber/main.go#L80

const (
	AutoReconnect   = true
	subsribeChannel = "hm/cmd/#"
)

var (
	// TODO Set Last Will
	opts *mqtt.ClientOptions
	c    mqtt.Client

	id, err  = machineid.ProtectedID("HomeMaticMqttPlugin")
	ClientID = "HomeMaticMqttPlugin_" + id[0:16]
)

func Init(config *shared.Configuration, handler mqtt.MessageHandler) {
	log.Println("Connect to Broker", config.BrokerUrl, "as", ClientID)
	opts = mqtt.NewClientOptions().AddBroker(config.BrokerUrl).SetClientID(ClientID).SetAutoReconnect(AutoReconnect)
	c = mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	t := c.Subscribe(subsribeChannel, 1, handler)
	go func() {
		<-t.Done()
		if t.Error() != nil {
			log.Println("ERROR SUBSCRIBING:", t.Error())
		} else {
			log.Println("subscribed to", subsribeChannel)
		}
	}()

}

func Disconnect() {
	log.Println("Disconnect from Broker")

	c.Disconnect(100)
}

func SendToBroker(e shared.Event) {
	start := time.Now()

	topic := "hm/" + e.SerialNumber + "/" + e.Type
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
