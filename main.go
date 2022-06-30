package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/products", GetProducts)

	return router
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
