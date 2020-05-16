package main

import (
	"log"
	"net/rpc"
)

type Args struct{}

func main() {
	var reply int64
	args := Args{}

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Call("Reciever.Time", args, &reply); err != nil {
		log.Fatal(err)
	}

	log.Printf("[REPLY]: %+v\n", reply)
}
