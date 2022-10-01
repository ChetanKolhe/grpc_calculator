package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ChetanKolhe/grpc_calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("This is the client")

	conn, err := grpc.Dial("0.0.0.0:50032", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error Occured %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculateServiceClient(conn)

	request := calculatorpb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 30,
	}

	sum_result, _ := c.Sum(context.Background(), &request)

	fmt.Println(sum_result)
	fmt.Printf("Connection is created %v", c)
}
