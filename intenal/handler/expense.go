package handler

import (
	"expanses_rest_api/intenal/expense"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) add(c *gin.Context) {
	var input expense.Expense
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err)
		return
	}
	err := h.services.AddExpense(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Error(err)
		return
	}
	logrus.Info(input)
	c.JSON(http.StatusOK, input)
}
func (h *Handler) getall(c *gin.Context) {
	fmt.Println("[info] test")
}
