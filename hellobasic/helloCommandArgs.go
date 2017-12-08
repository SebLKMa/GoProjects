package main

import ("fmt"; "flag"; "math/rand")

// go run helloCommandArgs.go -max=42 hello seb
func main() {
    // define our flags
    max := flag.Int("max", 6, "the max value")
    
    // parse
    flag.Parse()
    
    // after Parse, any non-flag args can also be retrieved with flag.Args() returning []string
    args := flag.Args()
    fmt.Println(args)
   
    // Generate a number between 0 and max
    fmt.Println(rand.Intn(*max))
}