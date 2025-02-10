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
