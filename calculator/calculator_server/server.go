package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"test-grpc-go/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Count(ctx context.Context, req *calculatorpb.CountRequest) (*calculatorpb.CountResponse, error) {
	firstNumber := req.GetCounting().GetFirstNumber()
	lastNumber := req.GetCounting().GetSecondNumber()
	penjumlahan := firstNumber + lastNumber
	result := "Hasil Dari penjumlahan " + strconv.Itoa(int(firstNumber)) + " dan " + strconv.Itoa(int(lastNumber)) + " adalah " + strconv.Itoa(int(penjumlahan))
	res := &calculatorpb.CountResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("hai ini adalah server")

	lis, err := net.Listen("tcp", "0.0.0.0:5005")
	if err != nil {
		log.Fatalf("gagal saat listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCountServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("gagal saat serve: %v", err)
	}

}
