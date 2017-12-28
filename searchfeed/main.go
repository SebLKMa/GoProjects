package main

import (
	"log"
	"os"
	_ "../searchfeed/matchers" 
	"../searchfeed/search"
)

// thia is called before main
func init() {
	// set the device for logging (default log writes to stderr)
	log.SetOutput(os.Stdout)
}

func main() {
	// perform search for this search term
	search.Run("Dreaming")
}