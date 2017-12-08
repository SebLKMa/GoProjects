package main

import "fmt"

func f1() {
    fmt.Println("f1")
}

func f2() {
    fmt.Println("f2")
}

// deferred functions are always called, often used to free resources deterministically 
func fopen() {
    fmt.Println("fopen")
}
func fclose() {
    fmt.Println("fclose")
}
func fcompute() {
    fmt.Println("fcompute")
}

func f5() {
    fopen()
    defer fclose() // deferred functions are always called
    fcompute()
}


func main() {
    defer f2() // deferred functions are always called
    f1()
    
    f5()
}