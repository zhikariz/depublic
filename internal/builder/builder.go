package builder

import (
	"github.com/zhikariz/depublic/config"
	"github.com/zhikariz/depublic/internal/http/handler"
	"github.com/zhikariz/depublic/internal/http/router"
	"github.com/zhikariz/depublic/internal/repository"
	"github.com/zhikariz/depublic/internal/service"
	"gorm.io/gorm"
)

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(cfg, userRepository)
	userHandler := handler.NewHandler(userService)
	return router.PrivateRoutes(userHandler)
}

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	transactionRepository := repository.NewTransactionRepository(db)
	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(cfg, userRepository)
	transactionService := service.NewTransactionService(transactionRepository)

	userHandler := handler.NewHandler(userService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	return router.PublicRoutes(userHandler, transactionHandler)
}
