package handler


import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"example.com/user/ginpoc/src/platform/newsfeed"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post string `json:"post"`

}

func NewsfeedPost(feed *newsfeed.Newsfeed) gin.HandlerFunc {
	return func(c *gin.Context){
	
		requestBody := newsfeedPostRequest{}
		
		//auto map incoming request body to this struct type 
		// based on field matching
		c.Bind(&requestBody) 
		
		
		item := newsfeed.Newsitem {
			Title: requestBody.Title,
			Post: requestBody.Post,
		}
		
		feed.Add(item)
		
		
		//no body in response
		c.Status(http.StatusNoContent)

	}
}