package router

import (
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
			Method:  echo.GET,
			Path:    "/users",
			Handler: userHandler.FindAllUser,
		},
	}
}

func PrivateRoutes() []*Route {
	return nil
}
