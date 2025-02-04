package main

import (
	"exercicio-04-djaar-rblf/tcp"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [server|client]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "server":
		tcp.Server()
	case "client":
		tcpResult := tcp.Client(10000)

		fmt.Println("Average RTT time:", tcpResult.Average, "ms")
		fmt.Println("Median RTT time:", tcpResult.Median, "ms")
		fmt.Println("Variance RTT time:", tcpResult.Variance, "ms")
	default:
		fmt.Println("Usage: go run main.go [server|client]")
		os.Exit(1)
	}
}
