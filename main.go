package main

import (
	"context"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
	if err := rpc.Register(arith); err != nil {
		log.Fatal("Failed to register:", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	defer listener.Close()

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Shutting down server...")
		cancel()
		listener.Close()
	}()

	log.Println("JSON-RPC Server is running on port 1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				wg.Wait()
				log.Println("Server stopped.")
				return
			default:
				log.Println("Failed to accept:", err)
			}
			continue
		}

		wg.Add(1)
		go func(conn net.Conn) {
			defer wg.Done()
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
