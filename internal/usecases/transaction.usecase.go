// chapa/internal/usecases/transaction_usecase.go
package usecases

import (
	"chapa/internal/domain"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"
)

type TransactionUseCase struct {
	transactionRepo domain.TransactionRepository
	webhookSecret   string
}

func NewTransactionUseCase(transactionRepo domain.TransactionRepository, webhookSecret string) *TransactionUseCase {
	return &TransactionUseCase{
		transactionRepo: transactionRepo,
		webhookSecret:   webhookSecret,
	}
}

func (uc *TransactionUseCase) InitiatePayment(ctx context.Context, txn *domain.Transaction) error {

	if txn.TxnRef == "" || txn.Amount <= 0 {
		return errors.New("invalid transaction data")
	}

	txn.Status = "pending"
	txn.CreatedAt = time.Now()

	err := uc.transactionRepo.CreateTransaction(ctx, txn)
	if err != nil {
		return err
	}

	return nil
}

func (uc *TransactionUseCase) VerifyTransaction(ctx context.Context, txnRef string) (*domain.Transaction, error) {
	txn, err := uc.transactionRepo.GetTransactionByRef(ctx, txnRef)
	if err != nil {
		return nil, err
	}

	if txn == nil {
		return nil, errors.New("transaction not found")
	}

	return txn, nil
}

func (uc *TransactionUseCase) ProcessWebhook(ctx context.Context, payload map[string]interface{}, signature string) error {

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return errors.New("failed to serialize webhook payload: " + err.Error())
	}

	expectedSignature := uc.computeHMACSHA256(payloadBytes)
	if signature != expectedSignature {
		return errors.New("invalid webhook signature")
	}

	txnRef, ok := payload["txn_ref"].(string)
	if !ok {
		return errors.New("missing txnRef in webhook payload")
	}
	status, ok := payload["status"].(string)
	if !ok {
		return errors.New("missing status in webhook payload")
	}

	err = uc.transactionRepo.UpdateTransactionStatus(ctx, txnRef, status)
	if err != nil {
		return errors.New("failed to update transaction status: " + err.Error())
	}

	return nil
}

func (uc *TransactionUseCase) computeHMACSHA256(data []byte) string {
	h := hmac.New(sha256.New, []byte(uc.webhookSecret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
