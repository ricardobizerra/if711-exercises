package main

import (
	"exercicio-04-djaar-rblf/tcp"
	"exercicio-04-djaar-rblf/udp"
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
			tcpResult := tcp.Client(1000, 40)

			fmt.Println("Average RTT time:", tcpResult.Average, "ms")
			fmt.Println("Median RTT time:", tcpResult.Median, "ms")
			fmt.Println("Variance RTT time:", tcpResult.Variance, "ms")
		default:
			fmt.Println("Usage: go run main.go [tcp|udp] [server|client]")
			os.Exit(1)
		}

	case "udp":
		switch os.Args[2] {
		case "server":
			udp.Server()
		case "client":
			udpResult := udp.Client(1000, 40)
			fmt.Println("Average RTT time:", udpResult.Average, "ms")
			fmt.Println("Median RTT time:", udpResult.Median, "ms")
			fmt.Println("Variance RTT time:", udpResult.Variance, "ms")
		default:
			fmt.Println("Usage: go run main.go [tcp|udp] [server|client]")
			os.Exit(1)
		}
	default:
		fmt.Println("Usage: go run main.go [tcp|udp] [server|client]")
		os.Exit(1)
	}
}
