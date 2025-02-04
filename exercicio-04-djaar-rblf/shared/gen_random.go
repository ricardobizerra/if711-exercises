package shared

import "math/rand"

func GenerateRandomMatrixes(size int) ([][]int, [][]int) {
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
