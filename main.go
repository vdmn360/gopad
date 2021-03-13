package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	var broker = ""
	var port = 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("")
	opts.SetUsername("")
	opts.SetPassword("")
	client := mqtt.NewClient(opts)
	client.Publish("test topic", 0, false, "message")

}
