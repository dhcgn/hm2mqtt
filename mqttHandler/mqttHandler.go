package mqtthandler

import (
	"log"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/dhcgn/hm2mqtt/shared"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// TOOO Change Impl to https://github.com/eclipse/paho.mqtt.golang/blob/43c9c445a89e7dca549a9bd445e3553dade9a2bc/cmd/docker/subscriber/main.go#L80

// Handle for send data to a broker
type Handle interface {
	SendToBroker(e shared.Event)
	Disconnect()
}

type handle struct {
	options *mqtt.ClientOptions
	client  mqtt.Client
}

const (
	autoReconnect    = true
	subscribeChannel = "hm/set/#"
)

var (
	// TODO Set Last Will
	id, err  = machineid.ProtectedID("HomeMaticMqttPlugin")
	clientID = "HomeMaticMqttPlugin_" + id[0:16]
)

// New creates a new mqttHandler to creates a client for subscription and publishing
func New(config *shared.Configuration, handler mqtt.MessageHandler) Handle {
	log.Println("Connect to Broker", config.BrokerURL, "as", clientID)
	opts := mqtt.NewClientOptions().AddBroker(config.BrokerURL).SetClientID(clientID).SetAutoReconnect(autoReconnect)
	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	t := c.Subscribe(subscribeChannel, 1, handler)
	go func() {
		<-t.Done()
		if t.Error() != nil {
			log.Println("ERROR SUBSCRIBING:", t.Error())
		} else {
			log.Println("subscribed to", subscribeChannel)
		}
	}()

	return &handle{
		options: opts,
		client:  c,
	}
}

func (h handle) Disconnect() {
	log.Println("Disconnect from Broker")

	h.client.Disconnect(100)
}

func (h handle) SendToBroker(e shared.Event) {
	start := time.Now()

	topic := "hm/" + e.SerialNumber + "/" + e.Type
	token := h.client.Publish(topic, 1, false, e.DataValue)
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
