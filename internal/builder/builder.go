package builder

import (
	"github.com/zhikariz/depublic/internal/http/handler"
	"github.com/zhikariz/depublic/internal/http/router"
	"github.com/zhikariz/depublic/internal/repository"
	"github.com/zhikariz/depublic/internal/service"
	"gorm.io/gorm"
)

func BuildPrivateRoutes() []*router.Route {
	return router.PrivateRoutes()
}

func BuildPublicRoutes(db *gorm.DB) []*router.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewHandler(userService)
	return router.PublicRoutes(userHandler)
}
