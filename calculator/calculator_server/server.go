package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ChetanKolhe/grpc_calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculateServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Serving the Sum Request %v", req)

	result := req.GetFirstNumber() + req.SecondNumber

	response := &calculatorpb.SumResponse{
		SumResult: result,
	}

	return response, nil
}

func main() {
	fmt.Println("This is Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50032")

	if err != nil {
		log.Fatalf("Not able to listen , %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
