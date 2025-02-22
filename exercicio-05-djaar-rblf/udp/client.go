package udp

import (
	"encoding/json"
	"exercicio-05-djaar-rblf/shared"
	"net"
	"time"
)

func Client(invocations int, a [][]int, b [][]int) {
	var response shared.Reply

	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	for i := 0; i < invocations; i++ {
		msgToServer := shared.Request{Operation: "Mul", A: a, B: b}

		startTime := time.Now()

		msg, err := json.Marshal(msgToServer)
		if err != nil {
			panic(err)
		}
		SendUDPMessage(conn, msg)
		msg, _ = ReceiveUDPMessage(conn)
		err = json.Unmarshal(msg, &response)
		if err != nil {
			panic(err)
		}

		elapsedTime := float64(time.Since(startTime).Milliseconds())

		shared.WriteRTTValue("udp-results.txt", elapsedTime)
	}
}
