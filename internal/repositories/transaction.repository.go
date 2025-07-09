package repositories

import (
	"chapa/internal/domain"
	"context"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, txn *domain.Transaction) error {
	return r.DB.WithContext(ctx).Create(txn).Error
}

func (r *TransactionRepository) GetTransactionByRef(ctx context.Context, ref string) (*domain.Transaction, error) {
	var txn domain.Transaction
	if err := r.DB.WithContext(ctx).Where("txn_ref = ?", ref).First(&txn).Error; err != nil {
		return nil, err
	}
	return &txn, nil
}

func (r *TransactionRepository) UpdateTransactionStatus(ctx context.Context, txnRef, status string) error {
	return r.DB.WithContext(ctx).Model(&domain.Transaction{}).Where("txn_ref = ?", txnRef).Update("status", status).Error
}
