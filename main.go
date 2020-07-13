package main

import (
	"go_news/news"
	"go_news/router"
)

func main() {

	r := router.New()
	a := news.NewArchive()

	go a.CollectNews()

	r.Run()
}
