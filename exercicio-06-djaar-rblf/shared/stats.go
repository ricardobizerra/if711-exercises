package shared

import (
	"fmt"
	"math"
	"sort"
)

func CalculateAverage(numeros []float64) float64 {
	if len(numeros) == 0 {
		return 0
	}

	soma := 0.0
	for i := range len(numeros) {
		soma += numeros[i]
	}

	return soma / float64(len(numeros))
}

func CalculateVariance(arr []float64, media float64) float64 {
	var somaQuadrados float64
	for _, v := range arr {
		somaQuadrados += (v - media) * (v - media)
	}
	return somaQuadrados / float64(len(arr))
}

func CalculateStandardDeviation(arr []float64, media float64) float64 {
	return math.Sqrt(CalculateVariance(arr, media))
}

func CalculateMedian(arr []float64) float64 {
	sort.Float64s(arr)

	n := len(arr)
	if n%2 == 1 {
		return arr[n/2]
	} else {
		return (arr[n/2-1] + arr[n/2]) / 2.0
	}
}

func CalculateStats(arr []float64) {
	average := CalculateAverage(arr)
	stdDev := CalculateStandardDeviation(arr, average)
	median := CalculateMedian(arr)

	fmt.Println("Média:        ", average, "ms")
	fmt.Println("Desvio padrão:", stdDev, "ms")
	fmt.Println("Mediana:      ", median, "ms")
}
