package main

import (
	"log"
	"os"
	_ "matchers" "search"
)

// thia is called before main
func init() {
	// set the device for logging (default log writes to stderr)
	log.SetOutput(os.Stdout)
}

func main() {
	// perform search for this search term
	search.Run("president")
}