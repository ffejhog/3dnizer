package main

import (
  "github.com/gin-gonic/gin"
  "github.com/ffejhog/3dnizer/controller"
  "github.com/ffejhog/3dnizer/middleware"
  "gorm.io/gorm"
  
)

func LoadRouter(router *gin.Engine) {
  router.GET("/ping", controller.GetPing)
}

func LoadMiddleware(router *gin.Engine, db *gorm.DB) {
  middleware.InitDb(db)
  router.Use(middleware.TokenAuthMiddleware(db))
}