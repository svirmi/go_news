package router

import (
	"go_news/news"
	"net/http"

	"github.com/gin-gonic/gin"
)

//The Greeting Message
const greetingMessage = ``

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title" : "Go news!",
	})
}

func searchHandler(c *gin.Context) {
	categoryName := c.Param("category")
	news.SearchFor(categoryName)

	topics := news.ResultFor(categoryName)

	c.HTML(http.StatusOK, "results.tmpl", gin.H{
		"topics": topics,
	})
}

func resetHandler(c *gin.Context) {
	news.ResetArchive()
	c.String(http.StatusOK, "The archive has been emptied successfully.")
}
