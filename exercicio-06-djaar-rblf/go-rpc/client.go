package goRpc

import (
	"exercicio-06-djaar-rblf/shared"
	"net/rpc"
	"time"
)

func Client(invocations int, a [][]int, b [][]int) {
	var response shared.Reply

	client, err := rpc.Dial("tcp", "rpc-server:8080")

	if err != nil {
		panic(err)
	}

	defer client.Close()

	for i := 0; i < invocations; i++ {
		msgToServer := shared.Request{Operation: "Mul", A: a, B: b}

		startTime := time.Now()

		err := client.Call("Matrix.Multiply", msgToServer, &response)

		if err != nil {
			panic(err)
		}

		elapsedTime := float64(time.Since(startTime).Nanoseconds()) / 1000000

		shared.WriteRTTValue("/data/go-rpc-results.txt", elapsedTime)
	}
}
