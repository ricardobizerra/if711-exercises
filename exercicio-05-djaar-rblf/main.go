package main

import (
	goRpc "exercicio-05-djaar-rblf/go-rpc"
	grpcClient "exercicio-05-djaar-rblf/grpc/client"
	grpcServer "exercicio-05-djaar-rblf/grpc/server"
	"exercicio-05-djaar-rblf/shared"
	"exercicio-05-djaar-rblf/tcp"
	"fmt"
	"os"
)

func main() {
	dim := 20
	max_value := 100
	invocations := 10000

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [tcp|go-rpc|grpc] [server|client]")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "grpc":
		switch os.Args[2] {
		case "server":
			grpcServer.Server()
		case "client":
			a, b := shared.GenerateRandomMatrixes32(dim, max_value)

			grpcClient.Client(invocations, a, b)

			rttValues, err := shared.ReadRTTValues("grpc-results.txt")
			if err != nil {
				panic(err)
			}
			shared.CalculateStats(rttValues)
		}
	case "tcp":
		switch os.Args[2] {
		case "server":
			tcp.Server()
		case "client":
			a, b := shared.GenerateRandomMatrixes(dim, max_value)

			tcp.Client(invocations, a, b)

			rttValues, err := shared.ReadRTTValues("tcp-results.txt")

			if err != nil {
				panic(err)
			}

			shared.CalculateStats(rttValues)
		default:
			fmt.Println("Usage: go run main.go [tcp|go-rpc|grpc] [server|client]")
			os.Exit(1)
		}

	case "go-rpc":
		fmt.Println("Usage: go run main.go [tcp] [server|client]")
		switch os.Args[2] {
		case "server":
			goRpc.Server()
		case "client":
			a, b := shared.GenerateRandomMatrixes(dim, max_value)

			number_clients := 20
			goRpc.Client(invocations, a, b, number_clients)

			rttValues, err := shared.ReadRTTValues("go-rpc-results.txt")

			if err != nil {
				panic(err)
			}

			shared.CalculateStats(rttValues)
		default:
			fmt.Println("Usage: go run main.go [tcp|go-rpc|grpc] [server|client]")
			os.Exit(1)
		}
	default:
		fmt.Println("Usage: go run main.go [tcp|go-rpc|grpc] [server|client]")
		os.Exit(1)
	}
}
