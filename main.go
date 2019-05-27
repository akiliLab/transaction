package main

import (
	"context"
	"log"
	"time"

	transaction "github.com/akililab/transaction/proto"
	"github.com/micro/go-micro"
)

// Transaction : This defines Transaction struct
type Transaction struct{}

// GetTransactions : Return All transactions per given account_id
func (t *Transaction) GetTransactions(ctx context.Context, req *transaction.TransactionRequest, rsp *transaction.TransactionReply) error {
	log.Print("Received GetTransaction request")
	return nil
}

func main() {
	// Setup dial with balance service
	service := micro.NewService(
		micro.Name("akililab.transaction"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	transaction.RegisterTransactionHandler(service.Server(), new(Transaction))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
