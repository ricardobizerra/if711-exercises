package goRpc

import (
	"fmt"
	"net"
	"net/rpc"
)

func Server() {
	go rpcServer()
	fmt.Scanln()
}

func rpcServer() {
	matrix_service := new(MatrixService)

	server := rpc.NewServer()
	server.RegisterName("Matrix", matrix_service)

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server listening on port 8080")

	server.Accept(ln)
}
