package main

import (
	"log"
	"seanHome/pkg/data"
	"seanHome/pkg/server"
)

func main() {
	data, err := data.Read()
	if err != nil {
		log.Println(err)
	}
	server.Run(data)
}
