package main

import (
	"fmt"
	"log"
	"net"

	"github.com/MAKOSCAFEE/malengo-pay/db"
	balance "github.com/ubunifupay/balance/pb"

	pb "github.com/ubunifupay/transaction/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var balanceClient balance.BalanceServiceClient

func main() {
	// Setup dial with balance service
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	balanceClient = balance.NewBalanceServiceClient(conn)

	// Setting up grpc server
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransactionServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) StoreTransaction(ctx context.Context, request *pb.TransactionRequest) (*pb.TransactionReply, error) {
	rep := &pb.TransactionReply{Completed: true}
	// Storing the transaction into the database
	db.StoreTransaction(request)
	fmt.Println("stored in db")
	// Credit or debit the balance by sending a new request to BalanceService
	balanceRequest := &balance.BalanceRequest{
		AccountID: request.AccountID,
		Value:     request.Amount,
		Currency:  request.Currency,
	}
	res, err := balanceClient.ManageBalance(ctx, balanceRequest)
	if err != nil || !res.Completed {
		rep.Completed = false
		return rep, err
	}
	return rep, nil
}
