package main

import (
	"log"
	"os"
	"fmt"
	"flag"
	_ "../searchfeed/matchers"
	"../searchfeed/search"
)

// thia is called before main
func init() {
	// set the device for logging (default log writes to stderr)
	log.SetOutput(os.Stdout)
}

func main() {
	// command line flag -term="some words"
	searchTerm := flag.String("term", "president", "the search term")

    // parse
    flag.Parse()
    
    // after Parse, any non-flag args can also be retrieved with flag.Args() returning []string
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println(args)
	}

	// perform search for this search term
	search.Run(*searchTerm)
}