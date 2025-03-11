package shared

import "math/rand"

func GenerateRandomMatrixes(dim int, max_value int) ([][]int, [][]int) {
	a := make([][]int, dim)
	b := make([][]int, dim)

	for i := range a {
		a[i] = make([]int, dim)
		b[i] = make([]int, dim)
	}

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			a[i][j] = rand.Intn(max_value)
			b[i][j] = rand.Intn(max_value)
		}
	}

	return a, b
}

func GenerateRandomMatrixes32(dim int, max_value int) ([][]int32, [][]int32) {
	a := make([][]int32, dim)
	b := make([][]int32, dim)

	for i := range a {
		a[i] = make([]int32, dim)
		b[i] = make([]int32, dim)
	}

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			a[i][j] = int32(rand.Intn(max_value))
			b[i][j] = int32(rand.Intn(max_value))
		}
	}

	return a, b
}
