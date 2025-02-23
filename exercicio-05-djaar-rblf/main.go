package main

import (
	goRpc "exercicio-05-djaar-rblf/go-rpc"
	"exercicio-05-djaar-rblf/shared"
	"exercicio-05-djaar-rblf/tcp"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [tcp|go-rpc] [server|client]")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "tcp":
		switch os.Args[2] {
		case "server":
			tcp.Server()
		case "client":
			a, b := shared.GenerateRandomMatrixes(60, 100)

			tcp.Client(10000, a, b)

			rttValues, err := shared.ReadRTTValues("tcp-results.txt")

			if err != nil {
				panic(err)
			}

			shared.CalculateStats(rttValues)
		default:
			fmt.Println("Usage: go run main.go [tcp|go-rpc] [server|client]")
			os.Exit(1)
		}

	case "go-rpc":
		switch os.Args[2] {
		case "server":
			goRpc.Server()
		case "client":
			a, b := shared.GenerateRandomMatrixes(60, 100)

			invocations := 10000
			number_clients := 20
			goRpc.Client(invocations, a, b, number_clients)

			rttValues, err := shared.ReadRTTValues("go-rpc-results.txt")

			if err != nil {
				panic(err)
			}

			shared.CalculateStats(rttValues)
		default:
			fmt.Println("Usage: go run main.go [tcp|go-rpc] [server|client]")
			os.Exit(1)
		}
	default:
		fmt.Println("Usage: go run main.go [tcp|go-rpc] [server|client]")
		os.Exit(1)
	}
}
