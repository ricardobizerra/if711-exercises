package tcp

import (
	"encoding/json"
	"exercicio-04-djaar-rblf/shared"
	"net"
	"time"
)

type Result struct {
	Average  float64
	Variance float64
	Median   float64
}

func Client(invocations int, matrix_size int, max_value_matrix int) Result {
	var response shared.Reply
	RTTList := [](float64){}

	r, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {
		panic(err)
	}

	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for i := 0; i < invocations; i++ {
		a, b := shared.GenerateRandomMatrixes(matrix_size, max_value_matrix)

		msgToServer := shared.Request{Operation: "Mul", A: a, B: b}

		startTime := time.Now()

		err := jsonEncoder.Encode(msgToServer)
		if err != nil {
			panic(err)
		}

		err = jsonDecoder.Decode(&response)
		if err != nil {
			panic(err)
		}

		elapsedTime := float64(time.Since(startTime).Microseconds()) / 1000
		RTTList = append(RTTList, elapsedTime)
	}

	average := shared.CalculateAverage(RTTList)

	return Result{
		Average:  average,
		Median:   shared.CalculateMedian(RTTList),
		Variance: shared.CalculateVariance(RTTList, average),
	}
}
