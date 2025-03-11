package server

import (
	"context"
	pb "exercicio-06-djaar-rblf/grpc/grpc"
	"exercicio-06-djaar-rblf/grpc/sharedGrpc"
	"exercicio-06-djaar-rblf/matrix"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMatrixMulServer
}

type OpError struct {
	message string
}

func (e *OpError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

// Mul implements grpc.MatrixMulServer
func (s *server) Mul(_ context.Context, in *pb.Request) (*pb.Reply, error) {
	op := in.GetOp()
	// log.Printf("Received: %v", op)
	if op != "Mul" {
		return nil, &OpError{"Wrong Operation"}
	}
	matrix1 := sharedGrpc.UnparseMatrix(in.GetM1())
	matrix2 := sharedGrpc.UnparseMatrix(in.GetM2())
	matrix_res := matrix.Multiply32(matrix1, matrix2)
	matrix_parse := sharedGrpc.ParseMatrix(matrix_res)
	return &pb.Reply{M: matrix_parse}, nil
}

func Server() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMatrixMulServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
