package search

import (
	"log" 
	"sync"
)

// package-level variable in lowercase is not exported to other packages.
// a map of registered matchers for searching, key=string value=Matcher instance
var matchers = make(map[string]Matcher) // Matcher type is defined in match.go
										// reference type must be created using make()
										
// This function contains the search logic
func Run(searchTerm string) {
	// get the list of feeds
	feeds, err := RetrieveFeeds() // := declare and initialize
	if err != nil {
		log.Fatal(err) // logs to terminal before terminating program
	}

	// create the unbuffered channel to receive match results
	// channel implements a communication between goroutines via a queue of typed values
	resultsChannel := make(chan *Result)

	// set up a wait group to process all the feeds
	// WaitGroup tracks a group of goroutines complete work before program termination
	var waitGroup sync.WaitGroup

	// set the no. of goroutines we need to wait for while they process the individual feeds
	waitGroup.Add(len(feeds))

	// launch a goroutine for each feed to find the results
	for _, feed := range feeds {
		// get the matcher for the search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"] // use default of matcher not found
		}

		// launch goroutine, matcher and feed instances will change with each iteration
		go func(matcher Matcher, feed *Feed) { // passing by value, including th address of pointer Feed 
			Match(matcher, feed, searchTerm, resultsChannel) // Match() from match.go
			waitGroup.Done() // decrements the WaitGroup semaphore count
		}(matcher, feed)
	}

	// launch goroutine to monitor when all work is done
	go func() {
		// wait for all are processed
		waitGroup.Wait()

		// close the channel to signal Display function that we can exit program
		close(resultsChannel)
	}()

	// start displaying results as soon as they are available
	// and return after the final result is displayed.
	Display(resultsChannel) // Display from match.go
}

// Registers a matcher for use by program
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}




