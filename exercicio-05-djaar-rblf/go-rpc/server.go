package goRpc

import (
	"fmt"
	"net"
	"net/rpc"
)

func Server() {
	matrix_service := new(MatrixService)

	server := rpc.NewServer()
	server.RegisterName("Matrix", matrix_service)

	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server listening on port 8080")

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		go server.ServeConn(conn)
	}
}
