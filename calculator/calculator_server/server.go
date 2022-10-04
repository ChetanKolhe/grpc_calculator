package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/ChetanKolhe/grpc_calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	calculatorpb.UnimplementedCalculateServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Serving the Sum Request %v", req)

	result := req.GetFirstNumber() + req.SecondNumber

	response := &calculatorpb.SumResponse{
		SumResult: result,
	}

	return response, nil
}

func (*server) SquareRoot(context context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	input_number := req.GetNumber()
	fmt.Println("Recived Square root rpc ")
	if input_number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalidi Argument is provided %v", input_number),
		)
	}

	return &calculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(input_number)),
	}, nil
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
