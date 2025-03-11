package sharedGrpc

import (
	pb "exercicio-06-djaar-rblf/grpc/grpc"
)

func ParseMatrix(m [][]int32) *pb.Matrix {
	rows := make([]*pb.Row, 0)
	for i := range m {
		rows = append(rows, &pb.Row{Row: m[i]})
	}
	return &pb.Matrix{M: rows}
}

func UnparseMatrix(m *pb.Matrix) [][]int32 {
	result := make([][]int32, 0)
	matrix := m.GetM()
	for i := range matrix {
		row := matrix[i]
		result = append(result, row.Row)
	}
	return result
}
