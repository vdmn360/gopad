package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type MQTTString struct {
	DeviceID string
	Host     string
	Port     int16
	Topic    string
	SASToken string
}

func read_config_file() ([]byte, error) {
	config_path := os.Getenv("GOPAD_CONFIG")
	data, err := ioutil.ReadFile(config_path)
	return data, err
}

func ReadMQTTConfig() (MQTTString, error) {
	data, err := read_config_file()
	var mqttString MQTTString
	json.Unmarshal(data, &mqttString)
	return mqttString, err

}
