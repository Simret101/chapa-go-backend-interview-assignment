package domain

import (
	"context"
)

type TransferUseCase interface {
	InitiateTransfer(ctx context.Context, transfer *Transfer) error
	CheckTransferStatus(ctx context.Context, ref string) (*Transfer, error)
	ProcessWebhook(ctx context.Context, payload map[string]interface{}, signature string) error
}

type TransferRepository interface {
	CreateTransfer(ctx context.Context, transfer *Transfer) error
	GetTransferByRef(ctx context.Context, ref string) (*Transfer, error)
	UpdateTransferStatus(ctx context.Context, ref string, status string) error
}
