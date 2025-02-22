package main

import (
	"exercicio-05-djaar-rblf/shared"
	"exercicio-05-djaar-rblf/tcp"
	"exercicio-05-djaar-rblf/udp"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [tcp|udp] [server|client]")
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
			fmt.Println("Usage: go run main.go [tcp|udp] [server|client]")
			os.Exit(1)
		}

	case "udp":
		switch os.Args[2] {
		case "server":
			udp.Server()
		case "client":
			a, b := shared.GenerateRandomMatrixes(60, 100)

			udp.Client(10000, a, b)

			rttValues, err := shared.ReadRTTValues("udp-results.txt")

			if err != nil {
				panic(err)
			}

			shared.CalculateStats(rttValues)
		default:
			fmt.Println("Usage: go run main.go [tcp|udp] [server|client]")
			os.Exit(1)
		}
	default:
		fmt.Println("Usage: go run main.go [tcp|udp] [server|client]")
		os.Exit(1)
	}
}
