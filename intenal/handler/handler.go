package handler

import (
	"expanses_rest_api/intenal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign_up", h.sign_up)
		auth.POST("/sign_in", h.sign_in)
	}
	expense := router.Group("/expenses")
	{
		expense.POST("/add", h.add)
		expense.PUT("/update", h.update)
		expense.GET("/getall", h.getall)
	}

	return router
}
