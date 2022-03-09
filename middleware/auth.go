package middleware

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "fmt"
)

type ApiAuth struct {
  gorm.Model
  ApiKey  string
}

func InitDb(db *gorm.DB) {
  db.AutoMigrate(&ApiAuth{})
}

func respondWithError(c *gin.Context, code int, message interface{}) {
  c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
  var apiAuth ApiAuth

  return func(c *gin.Context) {
    token := c.GetHeader("X-API-Key")
    fmt.Println(token)

    if token == "" {
      respondWithError(c, 401, "Must Supply Api Token Header")
      return
    }

    result := db.Where("api_key = ?", token).First(&apiAuth)

    if result.Error == gorm.ErrRecordNotFound {
      respondWithError(c, 401, "API Token Invalid")
      return
    }

    c.Next()
  }
}