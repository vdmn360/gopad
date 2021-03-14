package mqtt_handlers

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Println(msg)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("client connected!!")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Println(err)
}

func SetMQTTHandlers(opts *mqtt.ClientOptions) {
	opts.OnConnect = connectHandler
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnectionLost = connectionLostHandler
}
