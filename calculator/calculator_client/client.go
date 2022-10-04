package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ChetanKolhe/grpc_calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
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

	calculateWithNegativeValue(c)
}

func calculateWithNegativeValue(c calculatorpb.CalculateServiceClient) {

	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{
		Number: -20,
	})

	respErr, ok := status.FromError(err)

	if ok {
		fmt.Println(respErr.Proto())
		fmt.Println(respErr.Code())
		fmt.Println(respErr.Err())
		fmt.Println(respErr.String())
		fmt.Println(respErr.Proto().Message)
	} else {

		log.Fatalf("Error Occur for this number %v", err)
	}

	fmt.Printf("Response :%v", res)

}
