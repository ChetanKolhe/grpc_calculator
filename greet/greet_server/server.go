package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/ChetanKolhe/grpc_calculator/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetDeadlineRequest) (*greetpb.GreetDeadlineResponse, error) {
	fmt.Printf("Greet With Dead Line function invoked %v", req)

	for i := 0; i < 3; i++ {

		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Dead Line Exceded")

			return nil, status.Errorf(
				codes.DeadlineExceeded, "Dead Line Exceded",
			)
		}
		time.Sleep(1 * time.Second)
	}

	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	response := &greetpb.GreetDeadlineResponse{
		Result: result,
	}

	return response, nil
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

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {

	fmt.Println("Reciving Client Stream Request ")
	result := ""
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error occur while serving client request %v", err)
			return nil
		}

		firstName := res.GetGreet().GetFirstName()

		result = "Hello " + firstName + "!"
		fmt.Printf("Recived : %v \n", result)
	}

}

func (*server) GreetEveryOne(stream greetpb.GreetService_GreetEveryOneServer) error {

	fmt.Println("Reciving BiDi Stream Request ")
	result := ""
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error occur while serving client request %v", err)
			return nil
		}

		firstName := res.GetGreet().GetFirstName()

		result = "Hello " + firstName + "!"
		fmt.Printf("Recived : %v \n", result)
		stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
	}
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
