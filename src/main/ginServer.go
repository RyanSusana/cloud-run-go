package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.String(204, "No content.")
	})
	r.GET("/billie-jean", func(context *gin.Context) {
		context.String(http.StatusOK, "Is not my lover!")
	})

	r.Run()
}
