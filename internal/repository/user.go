package repository

import (
	"context"

	"github.com/zhikariz/depublic/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	users := make([]entity.User, 0)

	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
