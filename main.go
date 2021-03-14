package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	var broker = "broker.hivemq.com"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("test_id")
	opts.SetUsername("")
	opts.SetPassword("")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	client.Publish("testtopic13121", 0, false, "message")

}
