package router

import (
	"github.com/gin-gonic/gin"
)

//New engine pointer
func New() *gin.Engine {
	r := gin.Default()

	r.GET("/", indexHandler)
	r.GET("/search/:category", searchHandler)
	r.GET("/result/:category", resultHandler)
	r.GET("/reset", resetHandler)

	return r
}
