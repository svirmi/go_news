package router

import (
	"fmt"
	"go_news/news"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//The Greeting Message
const greetingMessage = `Welcome to the Grapper app!
* This app allow you to search news for some topic.
* Hit /search/technology and then his /result/technology as an example
`

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, greetingMessage)
}

func searchHandler(c *gin.Context) {

	categoryName := c.Param("category")
	log.Println("we're in searchHandler, category: ", categoryName)

	news.SearchFor(categoryName)

	c.String(http.StatusOK, "Search for %s category is intialized", categoryName)
}

func resultHandler(c *gin.Context) {

	categoryName := c.Param("category")
	topics := news.ResultFor(categoryName)

	fmt.Println("category: ", categoryName)
	fmt.Println("topics: ", topics)

	c.JSON(http.StatusOK, topics)
}

func resetHandler(c *gin.Context) {

	news.ResetArchive()
	c.String(http.StatusOK, "The archive is emptied successfully!")
}
