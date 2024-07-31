package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zhikariz/depublic/internal/http/handler"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func PublicRoutes(userHandler *handler.UserHandler) []*Route {
	return []*Route{
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: userHandler.Login,
		},
		{
			Method:  http.MethodPost,
			Path:    "/generate-password",
			Handler: userHandler.GeneratePassword,
		},
	}
}

func PrivateRoutes(userHandler *handler.UserHandler) []*Route {
	return []*Route{
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: userHandler.FindAllUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: userHandler.CreateUser,
		},
		{
			Method:  http.MethodPut,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
		},
	}
}
