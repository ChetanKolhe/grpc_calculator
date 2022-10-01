package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

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

func (*server) GreetManyTime(req *greetpb.GreetManyTimeRequest, stream greetpb.GreetService_GreetManyTimeServer) error {
	fmt.Println("GreetManyRequest Stream get executed ")
	firstName := req.Greeting.GetFirstName()

	for i := 0; i < 10; i++ {

		result := "Hello " + firstName + strconv.Itoa(i)
		stream.Send(&greetpb.GreetManyTimeResponse{
			Result: result,
		})

		time.Sleep(1000 * time.Microsecond)
	}

	return nil
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
