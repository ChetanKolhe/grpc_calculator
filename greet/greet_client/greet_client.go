package main

import (
	"context"
	"fmt"
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
	fmt.Printf("Connection is created %v", c)
}
