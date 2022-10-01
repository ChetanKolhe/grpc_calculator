package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/ChetanKolhe/grpc_calculator/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("This is the client")

	conn, err := grpc.Dial("0.0.0.0:50032", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error Occured %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	doUniary(c)
	doMultipResponse(c)

	fmt.Printf("Connection is created %v", c)
}

func doMultipResponse(c greetpb.GreetServiceClient) {

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
