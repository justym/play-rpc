package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {}

type Reciever int64

func (r *Reciever) Time(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	log.Printf("Got Request at %v",*reply)
	return nil
}

func main() {
	timeRcrv := new(Reciever)
	if err := rpc.Register(timeRcrv); err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()

	listenner, err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(listenner, nil))
}

