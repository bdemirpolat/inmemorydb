package main

import (
	"encoding/json"
	"fmt"
	"inmemorydb/model"
	"net"
)

var publisherAddr = ":9091"

func main() {
	fmt.Println(process(model.SET, "NAME", "AHMET"))
	fmt.Println(process(model.GET, "NAME"))

	fmt.Println(process(model.SET, "NAME", "MEHMET"))
	fmt.Println(process(model.GET, "NAME"))
}

func process(messageType int, key string, values ...string) model.Response {
	conn, err := net.Dial("tcp", publisherAddr)
	if err != nil {
		panic(err)
	}

	value := ""
	if len(values) > 0 {
		value = values[0]
	}

	registerMessage := model.Message{
		Type:  messageType,
		Key:   key,
		Value: value,
	}
	jsonMessage, err := json.Marshal(registerMessage)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write(jsonMessage)
	if err != nil {
		panic(err)
	}
	responseBytes := make([]byte, 4096)
	n, err := conn.Read(responseBytes)
	if err != nil {
		panic(err)
	}
	responseTrim := responseBytes[:n]
	response := &model.Response{}
	err = json.Unmarshal(responseTrim, response)
	if err != nil {
		panic(err)
	}
	return *response
}
