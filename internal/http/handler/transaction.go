package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zhikariz/depublic/internal/service"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (h *TransactionHandler) FindTransactionByUserID(c echo.Context) error {
	var request struct {
		UserID int64 `param:"id"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	transactions, err := h.transactionService.FindTransactionByUserID(c.Request().Context(), request.UserID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": transactions})
}
