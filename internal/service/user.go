package service

import (
	"context"

	"github.com/zhikariz/depublic/internal/dto"
	"github.com/zhikariz/depublic/internal/repository"
)

type UserService interface {
	FindAll(ctx context.Context) ([]dto.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository}
}

func (u *userService) FindAll(ctx context.Context) ([]dto.User, error) {
	users, err := u.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	usersDTO := make([]dto.User, 0)

	for _, v := range users {
		usersDTO = append(usersDTO, dto.User{ID: v.ID, Name: v.Name})
	}

	return usersDTO, nil
}
