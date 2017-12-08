package main

import ("fmt"; "strings")

func main() {
    // uses the strings package
    fmt.Println(strings.ToLower("HELLO"))
    fmt.Println(strings.ToUpper("world"))
    fmt.Println(strings.Contains("Ich bin", "bin"))
    fmt.Println(strings.Count("hallo", "l"))
    fmt.Println(strings.HasPrefix("sebastian", "seb"))
    fmt.Println(strings.HasSuffix("hallo", "lo"))
    fmt.Println(strings.Index("hallo", "a")) // returns 1, based on zero index
    fmt.Println(strings.Replace("boah boah boah", "o", "l", 2))
    
    strList1 := []string{"To", "err", "is", "human"}
    fmt.Println(strList1)
    strList2 := strings.Join(strList1, "-") // join list of strings to a string
    fmt.Println(strList2)    
    strList3 := strings.Split(strList2, "-") // split a string into a list of strings
    fmt.Println(strList3)
    
    bytes := []byte("text to binary")   // converts string to bytes
    fmt.Println(bytes)
    str := string(bytes)                // converts bytes to string
    fmt.Println(str)
}