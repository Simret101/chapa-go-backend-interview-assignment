package repositories

import (
	"chapa/internal/domain"
	"context"

	"gorm.io/gorm"
)

type TransferRepository struct {
	DB *gorm.DB
}

func NewTransferRepository(db *gorm.DB) *TransferRepository {
	return &TransferRepository{
		DB: db,
	}
}

func (r *TransferRepository) CreateTransfer(ctx context.Context, transfer *domain.Transfer) error {
	return r.DB.WithContext(ctx).Create(transfer).Error
}

func (r *TransferRepository) GetTransferByRef(ctx context.Context, ref string) (*domain.Transfer, error) {
	var transfer domain.Transfer
	if err := r.DB.WithContext(ctx).Where("transfer_ref = ?", ref).First(&transfer).Error; err != nil {
		return nil, err
	}
	return &transfer, nil
}

func (r *TransferRepository) UpdateTransferStatus(ctx context.Context, ref string, status string) error {
	return r.DB.WithContext(ctx).Model(&domain.Transfer{}).Where("transfer_ref = ?", ref).Update("status", status).Error
}
