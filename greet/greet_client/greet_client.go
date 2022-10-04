package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/ChetanKolhe/grpc_calculator/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("This is the client")

	conn, err := grpc.Dial("0.0.0.0:50032", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error Occured %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	// doUniary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBidiStreaming(c)
	doDeadlineRequest(c, 5*time.Second)
	doDeadlineRequest(c, 1*time.Second)

	fmt.Printf("Connection is created %v", c)
}

func doBidiStreaming(c greetpb.GreetServiceClient) {

	stream, _ := c.GreetEveryOne(context.Background())

	waitc := make(chan struct{})
	// send data
	go func() {
		for i := 0; i < 10; i++ {

			req := &greetpb.GreetEveryoneRequest{
				Greet: &greetpb.Greeting{
					FirstName: "chetan " + strconv.Itoa(i),
					LastName:  "kolhe",
				},
			}

			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error Occured to send value :%v", err)
			}
			time.Sleep(1000 * time.Microsecond)

		}
		stream.CloseSend()
	}()

	// recive data
	go func() {
		for {
			response, err := stream.Recv()

			if err == io.EOF {
				fmt.Println("Channel is closed ")
				break
			}

			if err != nil {
				log.Fatalf("Not able to recive data : %v", err)
				break
			}

			result := response.GetResult()

			fmt.Printf("Result : %v \n", result)

		}
		close(waitc)
	}()

	<-waitc

}

func doClientStreaming(c greetpb.GreetServiceClient) {

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error Occured %v", stream)
	}
	for i := 0; i < 10; i++ {

		request := &greetpb.LongGreetRequest{
			Greet: &greetpb.Greeting{
				FirstName: "Chetan" + strconv.Itoa(i),
				LastName:  "Kolhe",
			},
		}

		stream.Send(request)

	}

	result, _ := stream.CloseAndRecv()

	fmt.Printf("Recived Result : %v \n", result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {

	greet := greetpb.Greeting{
		FirstName: "Chetan",
		LastName:  "Kolhe",
	}

	request := &greetpb.GreetManyTimeRequest{
		Greeting: &greet,
	}

	response, err := c.GreetManyTime(context.Background(), request)

	if err != nil {
		log.Fatalf("Error Occur %v", err)
	}

	for {

		msg, err := response.Recv()

		if err == io.EOF {
			// break the loop , read all the strem
			break
		}
		if err != nil {
			log.Fatalf("Error Occured while reading the stream %v", err)
		}

		fmt.Printf("Response Stream : %v \n", msg.GetResult())
	}

	fmt.Println(response)

}

func doDeadlineRequest(c greetpb.GreetServiceClient, time time.Duration) {
	greet := greetpb.Greeting{
		FirstName: "Chetan",
		LastName:  "Kolhe",
	}

	request := &greetpb.GreetDeadlineRequest{
		Greeting: &greet,
	}

	ct := context.Background()
	ct, cancel := context.WithTimeout(ct, time)
	defer cancel()

	response, err := c.GreetWithDeadline(ct, request)

	if err != nil {

		statusError, ok := status.FromError(err)

		if ok {
			fmt.Println(statusError.Proto().Message)

			if statusError.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout occur , Dead line exceed ")
				fmt.Printf("Error from clint %v", statusError.Message())
			}

		} else {
			log.Fatalf("Unexpected Error Occur %v", err)
		}

	}

	fmt.Println(response)

}

func doUniary(c greetpb.GreetServiceClient) {
	greet := greetpb.Greeting{
		FirstName: "Chetan",
		LastName:  "Kolhe",
	}

	request := &greetpb.GreetRequest{
		Greeting: &greet,
	}

	response, err := c.Greet(context.Background(), request)

	if err != nil {
		log.Fatalf("Error Occur %v", err)
	}

	fmt.Println(response)

}
