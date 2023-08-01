package handler

import (
	"github.com/JamshedJ/todo/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	task := router.Group("/task", h.userIdentity)
	{
		task.POST("/", h.createTask)
		task.GET("/", h.getTask)
		task.GET("/:id", h.getByIdTask)
		task.PUT("/:id", h.updateTask)
		task.DELETE("/:id", h.deleteTask)
	}

	return router
}
