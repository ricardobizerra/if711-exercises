package udp

import (
	"encoding/json"
	"exercicio-04-djaar-rblf/matrix"
	"exercicio-04-djaar-rblf/shared"
	"fmt"
	"net"
	"os"
)

func Server() {
	r, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	ln, err := net.ListenUDP("udp", r)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server listening on", ln.LocalAddr())
	handleUDPConnection(ln)
}

func handleUDPConnection(conn *net.UDPConn) {
	var msgFromClient shared.Request

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	for {
		var msg []byte
		_, err := conn.Read(msg)

		for {
			if err != nil && err.Error() == "EOF" {
				break
			} else if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			if len(msg) > 0 {
				break
			}
		}
		fmt.Println(msg)
		err = json.Unmarshal(msg, &msgFromClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		operation := msgFromClient.Operation
		if operation != "Mul" {
			panic("The only operation accepted is 'Mul'.")
		}

		m1 := msgFromClient.A
		m2 := msgFromClient.B

		r := matrix.Multiply(m1, m2)

		replyToClient := shared.Reply{R: r}
		msgToClient, err := json.Marshal(replyToClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		_, err = conn.Write(msgToClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}
