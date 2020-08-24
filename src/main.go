package main

import (
	"fmt"
	
	"github.com/gin-gonic/gin"
	
	"example.com/user/ginpoc/src/handler"
	"example.com/user/ginpoc/src/platform/newsfeed"
)

func main() {
	fmt.Println("Starting main in Gin POC")
	
	
	feed := newsfeed.New()
	/*
	fmt.Println(feed)
	
	feed.Add(newsfeed.Newsitem{"Firstpost", "This is body of first post"})
	feed.Add(newsfeed.Newsitem{"Secondpost", "This is body of second post"})
	
	fmt.Println(feed)
	*/
	
	r := gin.Default()
	
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	
	
	
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}