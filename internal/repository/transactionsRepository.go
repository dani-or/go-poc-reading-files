package transactionsrepository

import (
	"nequi.com/poc-reading-files/internal/domain"
)

type TransactionsRepository interface {
	GetTransactions() (transaction.Transaction, error)
}