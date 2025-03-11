package client

import (
	"context"
	pb "exercicio-06-djaar-rblf/grpc/grpc"
	"exercicio-06-djaar-rblf/grpc/sharedGrpc"
	"exercicio-06-djaar-rblf/shared"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client(invocations int, a [][]int32, b [][]int32) {
	// Set up a connection to the server.
	conn, err := grpc.NewClient("rpc-server:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMatrixMulClient(conn)

	for i := 0; i < invocations; i++ {
		startTime := time.Now()
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		matrix1 := sharedGrpc.ParseMatrix(a)
		matrix2 := sharedGrpc.ParseMatrix(b)
		r, err := c.Mul(ctx, &pb.Request{Op: "Mul", M1: matrix1, M2: matrix2})
		if err != nil {
			log.Fatalf("could not multiply: %v", err)
		}
		_ = sharedGrpc.UnparseMatrix(r.GetM())
		//fmt.Printf("%v\n", matrix_res)

		// Tempo em milisegundos mais preciso
		elapsedTime := float64(time.Since(startTime).Nanoseconds()) / 1000000

		shared.WriteRTTValue("/data/grpc-results.txt", elapsedTime)
	}
}
