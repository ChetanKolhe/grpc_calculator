package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ChetanKolhe/grpc_calculator/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Printf("Greet function invoked %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	response := &greetpb.GreetResponse{
		Result: result,
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

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
