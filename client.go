package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	// connect to the server
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer client.Close()

	// prepare arguments
	args := &Args{3, 5}
	var reply int

	// call the remote procedure
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		log.Fatal("Failed to call Arith.Add:", err)
	}

	log.Printf("Result: %d + %d = %d\n", args.A, args.B, reply)
}
