package tcp

import (
	"encoding/json"
	"exercicio-04-djaar-rblf/shared"
	"math/rand"
	"net"
	"sort"
	"time"
)

type Result struct {
	Average  float64
	Variance float64
	Median   float64
}

func Client(invocations int) Result {
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
		a, b := generateRandomMatrixes(5)

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

		elapsedTime := float64(time.Since(startTime).Milliseconds())
		RTTList = append(RTTList, elapsedTime)
	}

	average := calculateAverage(RTTList)

	return Result{
		Average:  average,
		Median:   calculateMedian(RTTList),
		Variance: calculateVariance(RTTList, average),
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
			a[i][j] = rand.Intn(10)
			b[i][j] = rand.Intn(10)
		}
	}

	return a, b
}

func calculateAverage(numeros []float64) float64 {
	if len(numeros) == 0 {
		return 0
	}

	soma := 0.0
	for i := range len(numeros) {
		soma += numeros[i]
	}

	return soma / float64(len(numeros))
}

func calculateVariance(arr []float64, media float64) float64 {
	var somaQuadrados float64
	for _, v := range arr {
		somaQuadrados += (v - media) * (v - media)
	}
	return somaQuadrados / float64(len(arr))
}

func calculateMedian(arr []float64) float64 {
	sort.Float64s(arr)

	n := len(arr)
	if n%2 == 1 {
		return arr[n/2]
	} else {
		return (arr[n/2-1] + arr[n/2]) / 2.0
	}
}
