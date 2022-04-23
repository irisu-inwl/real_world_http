package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Calculator int

// Methods called externally by RPC
// cf: https://go-tour-jp.appspot.com/methods/1
func (c *Calculator) Multiply(args Args, result *int) error {
	log.Printf("Multiply called: %d, %d\n", args.A, args.B)
	*result = args.A * args.B
	return nil
}

// Argument called externally
type Args struct {
	A, B int
}

func main() {
	calculator := new(Calculator)
	// Create rpc server
	server := rpc.NewServer()
	// Register procedure
	server.Register(calculator)
	http.Handle(rpc.DefaultRPCPath, server)
	log.Println("start http listening :18888")
	listener, err := net.Listen("tcp", ":18888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
