package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zhikariz/depublic/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) FindAllUser(c echo.Context) error {
	users, err := h.userService.FindAll(c.Request().Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}
