package main

import "fmt"

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
func fcomputeException() {
    panic("COMPUTE EXCEPTION!")
}

func fdoCompute() {
    defer func() {
        str := recover()
        fmt.Println("fdoCompute Exception caught: " + str.(string))
    }()

    fopen()
    defer fclose() // deferred functions are always called
    fcomputeException() // this function throws
}

func frecoverDemo() {
    defer func() {
        str := recover()
        fmt.Println("frecoverDemo Exception caught: " + str.(string))
    }()
    
    panic("EXCEPTION!")
}

func fpanic() {
    panic("HELP!")
}

func main() {
    // there will be compile error is there is no other panic needs to be recovered
    // i.e. recover() returns nil
    defer func() {
        str := recover()
        if str != nil {
            fmt.Println("main Exception caught: " + str.(string))
        }
    }()
    
    frecoverDemo()
    fdoCompute()
    fpanic();
}