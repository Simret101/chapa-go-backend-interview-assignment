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

type TransferUseCase struct {
	transferRepo  domain.TransferRepository
	webhookSecret string
}

func NewTransferUseCase(transferRepo domain.TransferRepository, webhookSecret string) *TransferUseCase {
	return &TransferUseCase{
		transferRepo:  transferRepo,
		webhookSecret: webhookSecret,
	}
}

func (uc *TransferUseCase) InitiateTransfer(ctx context.Context, transfer *domain.Transfer) error {
	if transfer.TransferRef == "" || transfer.Amount <= 0 {
		return errors.New("invalid transfer data")
	}

	transfer.Status = "pending"
	transfer.CreatedAt = time.Now()

	err := uc.transferRepo.CreateTransfer(ctx, transfer)
	if err != nil {
		return errors.New("failed to create transfer: " + err.Error())
	}

	return nil
}

func (uc *TransferUseCase) CheckTransferStatus(ctx context.Context, ref string) (*domain.Transfer, error) {
	transfer, err := uc.transferRepo.GetTransferByRef(ctx, ref)
	if err != nil {
		return nil, errors.New("failed to get transfer: " + err.Error())
	}

	if transfer == nil {
		return nil, errors.New("transfer not found")
	}

	return transfer, nil
}

func (uc *TransferUseCase) ProcessWebhook(ctx context.Context, payload map[string]interface{}, signature string) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return errors.New("failed to serialize webhook payload: " + err.Error())
	}

	expectedSignature := uc.computeHMACSHA256(payloadBytes)
	if signature != expectedSignature {
		return errors.New("invalid webhook signature")
	}

	event, ok := payload["event"].(string)
	if !ok {
		return errors.New("missing event in webhook payload")
	}
	status, ok := payload["status"].(string)
	if !ok {
		return errors.New("missing status in webhook payload")
	}
	ref, ok := payload["transfer_ref"].(string)
	if !ok {
		return errors.New("missing transferRef in webhook payload")
	}

	if event != "transfer.success" {
		return errors.New("unsupported webhook event: " + event)
	}

	err = uc.transferRepo.UpdateTransferStatus(ctx, ref, status)
	if err != nil {
		return errors.New("failed to update transfer status: " + err.Error())
	}

	return nil
}

func (uc *TransferUseCase) computeHMACSHA256(data []byte) string {
	h := hmac.New(sha256.New, []byte(uc.webhookSecret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
