package tcp

import (
	"encoding/json"
	"exercicio-05-djaar-rblf/shared"
	"net"
	"time"
)

func Client(invocations int, a [][]int, b [][]int) {
	var response shared.Reply

	r, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {
		panic(err)
	}

	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for i := 0; i < invocations; i++ {
		msgToServer := shared.Request{Operation: "Mul", A: a, B: b}

		startTime := time.Now()

		err := jsonEncoder.Encode(msgToServer)
		if err != nil {
			panic(err)
		}

		err = jsonDecoder.Decode(&response)
		if err != nil {
			panic(err)
		}

		elapsedTime := float64(time.Since(startTime).Milliseconds())

		shared.WriteRTTValue("tcp-results.txt", elapsedTime)
	}
}
