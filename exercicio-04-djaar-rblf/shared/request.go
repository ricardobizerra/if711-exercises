package shared

type Request struct {
	Operation string  `json:"operation"`
	A         [][]int `json:"a"`
	B         [][]int `json:"b"`
}
