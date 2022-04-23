package main

import (
	"log"
	"net/rpc/jsonrpc"
	"os"
	"fmt"
)

// 引数
type Args struct {
	A, B int
}

func main() {
	host := os.Getenv("HOST")
	path := fmt.Sprintf("%s:18888", host)
	client, err := jsonrpc.Dial("tcp", path)
	if err != nil {
		panic(err)
	}
	var result int
	args := &Args{4, 5}
	err = client.Call("Calculator.Multiply", args, &result)
	if err != nil {
		panic(err)
	}
	log.Printf("4 x 5 = %d\n", result)
}
