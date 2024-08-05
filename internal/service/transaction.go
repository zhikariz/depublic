package service

import (
	"context"

	"github.com/zhikariz/depublic/internal/entity"
	"github.com/zhikariz/depublic/internal/repository"
)

type TransactionService interface {
	FindTransactionByUserID(ctx context.Context, userID int64) ([]entity.Transaction, error)
}

type transactionService struct {
	repository repository.TransactionRepository
}

func NewTransactionService(repository repository.TransactionRepository) TransactionService {
	return &transactionService{repository}
}

func (s *transactionService) FindTransactionByUserID(ctx context.Context, userID int64) ([]entity.Transaction, error) {
	return s.repository.FindTransactionByUserID(ctx, userID)
}
