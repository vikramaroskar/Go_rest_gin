package main

import (
	"fmt"
	
	"github.com/gin-gonic/gin"
	
	"example.com/user/ginpoc/src/handler"
)

func main() {
	fmt.Println("Starting main in Gin POC")
	
	r := gin.Default()
	
	r.GET("/ping", handler.PingGet())
	
	
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}