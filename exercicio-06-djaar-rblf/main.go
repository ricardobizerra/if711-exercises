package main

import (
	"exercicio-06-djaar-rblf/rabbitmq"
	"exercicio-06-djaar-rblf/shared"
	"fmt"
	"os"
	"strings"
)

func printAndExit() {
	protocols := []string{"rabbitmq"}
	operations := []string{"server", "client", "results"}

	fmt.Printf("Usage: go run main.go [%s] [%s]\n",
		strings.Join(protocols, "|"),
		strings.Join(operations, "|"))

	os.Exit(1)
}

func main() {
	dim := 20
	max_value := 100
	invocations := 10000

	if len(os.Args) < 3 {
		printAndExit()
	}

	switch os.Args[1] {

	case "rabbitmq":

		switch os.Args[2] {

		case "server":
			rabbitmq.Server()

		case "client":
			a, b := shared.GenerateRandomMatrixes(dim, max_value)

			rabbitmq.Client(invocations, a, b)

		case "results":
			rttValues, err := shared.ReadRTTValues("/app/data/rabbitmq-results.txt")
			if err != nil {
				fmt.Println("Error reading RTT values")
				panic(err)
			}
			shared.CalculateStats(rttValues)

		default:
			printAndExit()
		}

	default:
		printAndExit()
	}
}
