package mqttHandler

import (
	"github.com/dhcgn/gohomematicmqttplugin/src/shared"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

const ClientID = "HomeMaticMqttPlugin"
const AutoReconnect = true

var (
	// TODO Set Last Will
	opts *mqtt.ClientOptions
	c    mqtt.Client
)

func Init(config *shared.Configuration){
	log.Println("Connect to Broker", config.BrokerUrl, "as",ClientID )
	opts= mqtt.NewClientOptions().AddBroker(config.BrokerUrl).SetClientID(ClientID).SetAutoReconnect(AutoReconnect)
	c = mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func Disconnect(){
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

