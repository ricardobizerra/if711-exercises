package main

import (
	grcpClient "exercicio-05-djaar-rblf/grpc/client"
	grcpServer "exercicio-05-djaar-rblf/grpc/server"
	"exercicio-05-djaar-rblf/shared"
	"exercicio-05-djaar-rblf/tcp"
	"exercicio-05-djaar-rblf/udp"
	"fmt"
	"os"
)

func main() {
	dim := 20
	max_value := 100
	invocations := 10000

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [tcp|udp|grpc] [server|client]")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "grpc":
		switch os.Args[2] {
		case "server":
			grcpServer.Server()
		case "client":
			a, b := shared.GenerateRandomMatrixes32(dim, max_value)
			grcpClient.Client(invocations, a, b)
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
			fmt.Println("Usage: go run main.go [tcp|udp|grpc] [server|client]")
			os.Exit(1)
		}

	case "udp":
		switch os.Args[2] {
		case "server":
			udp.Server()
		case "client":
			a, b := shared.GenerateRandomMatrixes(dim, max_value)

			udp.Client(invocations, a, b)

			rttValues, err := shared.ReadRTTValues("udp-results.txt")

			if err != nil {
				panic(err)
			}

			shared.CalculateStats(rttValues)
		default:
			fmt.Println("Usage: go run main.go [tcp|udp|grpc] [server|client]")
			os.Exit(1)
		}
	default:
		fmt.Println("Usage: go run main.go [tcp|udp|grpc] [server|client]")
		os.Exit(1)
	}
}
