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
	logrus.Info("Expense successfully added")
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "ok",
		"message": "Expense successfully added",
	})
}
func (h *Handler) update(c *gin.Context) {
	var input expense.Expense
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err)
		return
	}
	err := h.services.UpdateExpense(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Error(err)
		return
	}
	logrus.Info("Expense successfully updated")
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "ok",
		"message": "Expense successfully updated",
	})
}

func (h *Handler) getall(c *gin.Context) {
	fmt.Println("[info] test")
}
