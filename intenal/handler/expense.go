package handler

import (
	"expanses_rest_api/intenal/expense"
	"fmt"
	"net/http"
	"strconv"

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

func (h *Handler) get(c *gin.Context) {

	str_id := c.Param("id")

	id, err := strconv.ParseInt(str_id, 10, 64)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err)
		return
	}

	response := h.services.GetExpense(id)

	logrus.Info("Get expense with id " + str_id)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) getall(c *gin.Context) {

	response := h.services.GetAllExpenses()

	logrus.Info("Get all expenses")
	c.JSON(http.StatusOK, response)

}

func (h *Handler) getlast(c *gin.Context) {
	var input struct {
		Limit  int64 `json:"limit"`
		Offset int64 `json:"offset"`
	}
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err)
		return
	}
	response := h.services.GetLastExpenses(input.Limit, input.Offset)

	logrus.Info(fmt.Sprintf("Get last expense with limit %d and offset %d", input.Limit, input.Offset))
	c.JSON(http.StatusOK, response)

}
