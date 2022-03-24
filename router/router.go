package router

import (
	"github.com/gin-gonic/gin"
)

//New engine pointer
func New() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("frontend/templates/*.tmpl")
	router.Static("/static", "./frontend/static")

	router.GET("/", indexHandler)
	router.GET("/search/:category", searchHandler)
	router.GET("/reset", resetHandler)

	return router
}
