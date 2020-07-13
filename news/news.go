package news

var (
	searchClause chan string
	returnClause chan string
	result       chan []Topic
	reset        chan bool
)

func init() {

	searchClause = make(chan (string))
	returnClause = make(chan (string))

	reset = make(chan bool)
	result = make(chan []Topic)

}

//SearchFor adds category to searchClose channel (see CollectNews)
func SearchFor(category string) {
	searchClause <- category
}

//ResultFor yields the result of search
func ResultFor(category string) []Topic {
	returnClause <- category
	return <-result
}

//ResetArchive resets in-memory storage
func ResetArchive() {
	reset <- true
}


//CollectNews loop of the main functionality of archive
func (a Archive) CollectNews() {
	for {
		select {
		case category := <-searchClause:
			a.collectNews(category)
		case category := <-returnClause:
			result <- a.data(category)
		case <-reset:
			a.resetData()
		}
	}

}
