package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zhikariz/depublic/internal/dto"
	"github.com/zhikariz/depublic/internal/service"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService service.UserService
}

func NewHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GeneratePassword(c echo.Context) error {

	var request struct {
		Password string `json:"password"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"password": string(encodedPassword)})
}

func (h *UserHandler) Login(c echo.Context) error {
	var request dto.LoginRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := h.userService.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

func (h *UserHandler) FindAllUser(c echo.Context) error {
	users, err := h.userService.FindAll(c.Request().Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var request dto.CreateUserRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.userService.Create(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "success create user"})
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var request dto.UpdateUserRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.userService.Update(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success update user"})
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	var request dto.DeleteUserRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.userService.Delete(c.Request().Context(), request.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success delete user"})
}
