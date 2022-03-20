package main

import (
	"context"
	"log"
	"test-grpc-go/calculator/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//definisikan grpc dial
	cc, err := grpc.Dial("localhost:5005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("gagal saat dial grpc: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCountServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CountServiceClient) {
	//definisikan request
	req := &calculatorpb.CountRequest{
		Counting: &calculatorpb.Counting{
			FirstNumber:  10,
			SecondNumber: 3,
		},
	}
	// handle response dan error
	res, err := c.Count(context.Background(), req)
	if err != nil {
		log.Fatalf("gagal saat request : %v", err)
	}
	log.Printf("%v", res.Result)
}
