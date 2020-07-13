package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getSourcesFor(category string) []Source {

	fmt.Println("we are in getSources with category:", category)
	body := getRequest(sourceURL(category))
	fmt.Println("body: ", string(body))

	var sourceAPI sourceAPIResponse
	json.Unmarshal(body, &sourceAPI)

	x := sourceAPI.Sources
	fmt.Println("sourceApi.Sources: ", x)

	return x
}

func getLatest(sources []Source) []Topic {

	log.Println("we are in getTopics with sources:", sources)
	var topics []Topic

	for _, source := range sources {
		body := getRequest(articleURL(source))

		var articleAPI newsAPIResponse
		json.Unmarshal(body, &articleAPI)

		topics = append(topics, articleAPI.Articles...)
	}

	return topics
}

func getRequest(url string) []byte {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error! Cannot get sources!")
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)

	}

	defer res.Body.Close()

	return body

}

func sourceURL(category string) string {
	return fmt.Sprintf("https://newsapi.org/v1/sources?language=en&category=%s", category)
}

func articleURL(source Source) string {
	return fmt.Sprintf("https://newsapi.org/v1/articles?source=%s&sortBy=latest&apiKey=e4304f0df55b4d66b10f69f1d7d3521e", source.ID)
}
