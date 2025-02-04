package udp

import (
	"encoding/json"
	"exercicio-04-djaar-rblf/shared"
	"fmt"
	"math/rand"
	"net"
	"time"
)

type Result struct {
	Average  float64
	Variance float64
	Median   float64
}

func Client(invocations int, matrix_size int) Result {
	var response shared.Reply
	RTTList := [](float64){}

	r, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, r)
	if err != nil {
		panic(err)
	}

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for i := 0; i < invocations; i++ {
		a, b := generateRandomMatrixes(matrix_size)

		msgToServer := shared.Request{Operation: "Mul", A: a, B: b}

		startTime := time.Now()

		err := jsonEncoder.Encode(msgToServer)
		if err != nil {
			panic(err)
		}
		fmt.Println("hi")
		err = jsonDecoder.Decode(&response)
		if err != nil {
			panic(err)
		}

		elapsedTime := float64(time.Since(startTime).Milliseconds())
		RTTList = append(RTTList, elapsedTime)
	}

	average := shared.CalculateAverage(RTTList)

	return Result{
		Average:  average,
		Median:   shared.CalculateMedian(RTTList),
		Variance: shared.CalculateVariance(RTTList, average),
	}
}

func generateRandomMatrixes(size int) ([][]int, [][]int) {
	a := make([][]int, size)
	b := make([][]int, size)

	for i := range a {
		a[i] = make([]int, size)
		b[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			a[i][j] = rand.Intn(100)
			b[i][j] = rand.Intn(100)
		}
	}

	return a, b
}
