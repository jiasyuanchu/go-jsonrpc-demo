package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Arith struct{}

type Args struct {
	A, B int
}

func (t *Arith) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	arith := new(Arith)
	err := rpc.Register(arith)
	if err != nil {
		log.Fatal("Failed to register:", err)
	}

	// 監聽 TCP 端口
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	defer listener.Close()

	log.Println("JSON-RPC Server is running on port 1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept:", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
