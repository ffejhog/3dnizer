package main

import "github.com/gin-gonic/gin"

type dependencies struct {
	gin *gin.Engine
}

func loadDependancies() dependencies {
	r := gin.Default()
	return dependencies{
		gin: r,
	}
}
