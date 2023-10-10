package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) add(c *gin.Context) {}
func (h *Handler) getall(c *gin.Context) {
	fmt.Println("[info] test")
}
