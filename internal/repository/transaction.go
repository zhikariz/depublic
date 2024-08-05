package repository

import (
	"context"

	"github.com/zhikariz/depublic/internal/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactionByUserID(ctx context.Context, userID int64) ([]entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) FindTransactionByUserID(ctx context.Context, userID int64) ([]entity.Transaction, error) {
	result := make([]entity.Transaction, 0)

	if err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Payment").
		Preload("Details.Product").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
