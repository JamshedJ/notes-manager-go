package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTask(c *gin.Context) {
	id, _ := c.Get("user_id")
	c.JSON(http.StatusOK, gin.H{"message": "OK", "id": id})
}

func (h *Handler) getTask(c *gin.Context) {}

func (h *Handler) getByIdTask(c *gin.Context) {}

func (h *Handler) updateTask(c *gin.Context) {}

func (h *Handler) deleteTask(c *gin.Context) {}
