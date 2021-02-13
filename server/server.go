package main

import (
	"encoding/json"
	"fmt"
	"inmemorydb/model"
	"net"
)

var db = map[string]string{}

func main() {
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		fmt.Println("connection closing")
		conn.Close()
	}()
	fmt.Println("connection opened")
	requestBytes := make([]byte, 4096)
	n, err := conn.Read(requestBytes)
	if err != nil {
		return
	}
	requestTrim := requestBytes[:n]
	data := &model.Message{}
	err = json.Unmarshal(requestTrim, data)
	if err != nil {
		response(conn, false, "")
		return
	}

	if data.Type == model.SET {
		db[data.Key] = data.Value
		response(conn, true, "")
		return
	}

	if data.Type == model.GET {
		getData, ok := db[data.Key]
		if !ok {
			response(conn, false, "")
			return
		}
		response(conn, true, getData)
	}
}

func response(conn net.Conn, status bool, data string) bool {
	r := model.Response{
		Status: status,
		Data:   data,
	}
	responseJson, _ := json.Marshal(r)
	_, err := conn.Write(responseJson)
	if err != nil {
		return false
	}
	return true
}
