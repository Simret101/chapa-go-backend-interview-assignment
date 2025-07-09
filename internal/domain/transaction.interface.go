package domain

import (
	"context"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, txn *Transaction) error
	GetTransactionByRef(ctx context.Context, txnRef string) (*Transaction, error)
	UpdateTransactionStatus(ctx context.Context, txnRef string, status string) error
}

type TransactionUseCase interface {
	InitiatePayment(ctx context.Context, txn *Transaction) error
	VerifyTransaction(ctx context.Context, txnRef string) (*Transaction, error)
	ProcessWebhook(ctx context.Context, payload map[string]interface{}, signature string) error
}
