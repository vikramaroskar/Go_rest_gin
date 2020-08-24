package handler


import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"example.com/user/ginpoc/src/platform/newsfeed"
)

func NewsfeedGet(feed *newsfeed.Newsfeed) gin.HandlerFunc {
	return func(c *gin.Context){
		results := feed.GetAll()
	
		c.JSON(http.StatusOK, results)

	}
}