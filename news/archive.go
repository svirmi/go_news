package news

import (
	"fmt"
	"log"
)

//Topic is basic Archive item
type Topic struct {
	Title  string `json:"title"`
	Desc   string `json:"description"`
	Author string `json:"author"`
	URL    string `json:"url"`
}

type newsAPIResponse struct {
	Articles []Topic `json:"articles"`
	Source   string  `json:"source"`
}

//Source is a part of news
type Source struct {
	ID string `json:"id"`
}

type sourceAPIResponse struct {
	Status  string
	Sources []Source `json:"sources"`
}

//Archive is the basic type of in-memory storage
type Archive map[string][]Topic

//NewArchive is a factory for Archive type
func NewArchive() Archive {
	fmt.Println("Create new archive.")
	return make(Archive)
}

func (a Archive) data(category string) []Topic {
	return a[category]
}

func (a Archive) resetData() {
	for c := range a {
		delete(a, c)
	}
	log.Println("Archive is clean.")
}

func (a Archive) collectNews(category string) {
	sources := getSources(category)
	topics := getLatest(sources)
	a[category] = topics
}
