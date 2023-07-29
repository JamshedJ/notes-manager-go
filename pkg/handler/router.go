package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	task := router.Group("/task")
	{
		task.POST("/", h.createTask)
		task.GET("/", h.getTask)
		task.GET("/:id", h.getByIdTask)
		task.PUT("/:id", h.updateTask)
		task.DELETE("/:id", h.deleteTask)
	}

	return router
}
