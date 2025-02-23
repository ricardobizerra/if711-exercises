package goRpc

import (
	"exercicio-05-djaar-rblf/shared"
	"net/rpc"
	"sync"
	"time"
)

func Client(invocations int, a [][]int, b [][]int, number_clients int) {
	var wg sync.WaitGroup

	for i := 0; i < number_clients; i++ {
		wg.Add(1)
		go rpcClient(&wg, invocations, a, b)
	}

	wg.Wait()
}

func rpcClient(wg *sync.WaitGroup, invocations int, a [][]int, b [][]int) {
	defer wg.Done()

	var response shared.Reply

	client, err := rpc.Dial("tcp", ":8080")

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

		elapsedTime := float64(time.Since(startTime).Milliseconds())

		shared.WriteRTTValue("go-rpc-results.txt", elapsedTime)
	}
}
