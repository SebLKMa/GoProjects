package main

import (
	"fmt"
)

func removeEntry(theMap map[string]string, theKey string) {
	delete(theMap, theKey)
}

func displayMap(theMap map[string]string) {
	fmt.Println("=== MAP LISTING ===")
	for key, value := range theMap {
		fmt.Printf("%s : %s\n", key, value)
	}
	fmt.Println("===================")
}

func main() {
	quotesMap := map[string]string{
		"To err":    "Is Human",
		"To be":     "Or not to be",
		"Beauty is": "Skin deep",
		"I think":   "Therefore I am",
	}

	// iterating the map
	displayMap(quotesMap)

	// passing a map to function does not make a copy
	// it references the same map
	removeEntry(quotesMap, "I think")

	// check by boolean returned
	myKey := "I think"
	value, exists := quotesMap[myKey]
	if exists {
		fmt.Printf("Key: %s, Found Value: %s\n", myKey, value)
	}

	// map always returned a nil value, even if value not found
	myKey = "Beauty is"
	value = quotesMap[myKey]
	if value != "" { // string type nil value is ""
		fmt.Printf("Key: %s, Found Value: %s\n", myKey, value)
	}

	// iterating the map
	displayMap(quotesMap)
}
