package udp

import (
	"encoding/json"
	"exercicio-05-djaar-rblf/matrix"
	"exercicio-05-djaar-rblf/shared"
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
	var msgRequest shared.Request

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	for {
		msg_byte, addr := ReceiveUDPMessage(conn)
		msg := string(msg_byte[:])
		err := json.Unmarshal([]byte(msg), &msgRequest)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		operation := msgRequest.Operation
		if operation != "Mul" {
			panic("The only operation accepted is 'Mul'.")
		}

		m1 := msgRequest.A
		m2 := msgRequest.B

		r := matrix.Multiply(m1, m2)

		replyToClient := shared.Reply{R: r}
		msgToClient, err := json.Marshal(replyToClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		SendUDPMessageAddr(conn, addr, msgToClient)
	}
}
