package search

import (
	"fmt"
	"log"
)

// definition of Result
type Result struct {
	Field string
	Content string
}

// interface type for Matcher
type Matcher interface {
	// Feed type from feed.go
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match would be launched as a concurrent goroutine
func Match(matcher Matcher, feed *Feed, searchTerm string, resultsChannel chan<- *Result) {
	// perform search based on search term
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// sends the results to channel
	for _, result := range searchResults {
		resultsChannel <- result
	}
}

// Writes results to terminal as soon as they are received by individual goroutines
func Display(resultsChannel chan *Result) {
	// blocks until results is sent to channel
	// once the channel is closed, this for loop terminates
	for result := range resultsChannel {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}