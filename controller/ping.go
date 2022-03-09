package controller

import (
  "github.com/gin-gonic/gin"

  "net/http"
)

func GetPing(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
} 