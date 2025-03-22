package main

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type Args struct {
	A, B int
}

func main() {
	var client *rpc.Client
	var err error

	for i := 0; i < 3; i++ {
		client, err = jsonrpc.Dial("tcp", "localhost:1234")
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer client.Close()

	args := &Args{3, 5}
	var reply int

	for i := 0; i < 3; i++ {
		err = client.Call("Arith.Add", args, &reply)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to call Arith.Add:", err)
	}

	log.Printf("Result: %d + %d = %d\n", args.A, args.B, reply)
}
