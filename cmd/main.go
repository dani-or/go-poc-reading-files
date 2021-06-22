package main

import (
	"fmt"
	//"log"
	"nequi.com/poc-reading-files/internal/platform/storage"
	"nequi.com/poc-reading-files/internal/services/transactions"
)

func main() {
	fmt.Println("Hello, World! Daniela readinf files")
	transactionsrepository := s3.NewS3Repository()
	getTransactionsService := services.NewGetTransactionsService(transactionsrepository)
	fmt.Println(getTransactionsService.GetTransactions())
}