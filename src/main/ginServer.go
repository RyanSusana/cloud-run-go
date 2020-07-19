package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/billie-jean", func(context *gin.Context) {
		context.String(http.StatusOK, "Is not my lover!")
	})

	r.POST("/pubsub", func(context *gin.Context) {
		all, _ := ioutil.ReadAll(context.Request.Body)
		fmt.Println("Response: " + string(all))
	})

	r.Run()
}
