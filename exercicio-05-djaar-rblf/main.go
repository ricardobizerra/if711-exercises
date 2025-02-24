package main

import (
	goRpc "exercicio-05-djaar-rblf/go-rpc"
	grpcClient "exercicio-05-djaar-rblf/grpc/client"
	grpcServer "exercicio-05-djaar-rblf/grpc/server"
	"exercicio-05-djaar-rblf/shared"
	"fmt"
	"os"
)

func main() {
	dim := 20
	max_value := 100
	invocations := 10000

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [go-rpc|grpc] [server|client|results]")
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
		case "results":
			rttValues, err := shared.ReadRTTValues("shared-volume/grpc-results.txt")
			if err != nil {
				fmt.Println("Error reading RTT values")
				panic(err)
			}
			shared.CalculateStats(rttValues)
		}
	case "go-rpc":
		switch os.Args[2] {
		case "server":
			goRpc.Server()
		case "client":
			a, b := shared.GenerateRandomMatrixes(dim, max_value)
			invocations := 10000
			goRpc.Client(invocations, a, b)
		case "results":
			rttValues, err := shared.ReadRTTValues("shared-volume/go-rpc-results.txt")
			if err != nil {
				fmt.Println("Error reading RTT values")
				panic(err)
			}
			shared.CalculateStats(rttValues)
		default:
			fmt.Println("Usage: go run main.go [tcp|go-rpc|grpc] [server|client|results]")
			os.Exit(1)
		}
	default:
		fmt.Println("Usage: go run main.go [tcp|go-rpc|grpc] [server|client|results]")
		os.Exit(1)
	}
}
