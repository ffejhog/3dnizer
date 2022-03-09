package main

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type dependencies struct {
	gin *gin.Engine
  db *gorm.DB
}

func loadDependancies() dependencies {
	r := gin.Default()
  db, err := gorm.Open(sqlite.Open("3dnizer.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  
	return dependencies{
		gin: r,
    db: db,
	}
}
