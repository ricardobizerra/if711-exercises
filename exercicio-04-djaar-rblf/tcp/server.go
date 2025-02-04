package tcp

import (
	"encoding/json"
	"exercicio-04-djaar-rblf/matrix"
	"exercicio-04-djaar-rblf/shared"
	"fmt"
	"net"
	"os"
)

func Server() {
	r, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	ln, err := net.ListenTCP("tcp", r)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server listening on port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	var msgFromClient shared.Request

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for {
		err := jsonDecoder.Decode(&msgFromClient)
		if err != nil && err.Error() == "EOF" {
			break
		}

		operation := msgFromClient.Operation
		if operation != "Mul" {
			panic("The only operation accepted is 'Mul'.")
		}

		m1 := msgFromClient.A
		m2 := msgFromClient.B

		r := matrix.Multiply(m1, m2)

		msgToClient := shared.Reply{R: r}

		err = jsonEncoder.Encode(msgToClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}
