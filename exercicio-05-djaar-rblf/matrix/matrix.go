package matrix

func Multiply(a, b [][]int) [][]int {
	if len(a[0]) != len(b) {
		return nil
	}

	c := make([][]int, len(a))
	for i := range c {
		c[i] = make([]int, len(b[0]))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c
}

func Multiply32(a, b [][]int32) [][]int32 {
	if len(a[0]) != len(b) {
		return nil
	}

	c := make([][]int32, len(a))
	for i := range c {
		c[i] = make([]int32, len(b[0]))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c
}
