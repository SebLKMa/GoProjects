package main

import ("fmt"; "time")

func pause() {
    var input string
    fmt.Scanln(&input) // pause
}

func SimpleSelect() {
    c1 := make(chan string)
    c2 := make(chan string)
    
    // start first goroutine
    go func() { // anonymous function
        for {
            c1 <- "sent from channel 1"
            time.Sleep(time.Second * 2)
        }
    }()
    
    // start second goroutine
    go func() { // anonymous function
        for {
            c2 <- "sent from channel 2"
            time.Sleep(time.Second * 3)
        }
    }()

    // start third goroutine that randomly picks a channel
    // will block if nothing received
    go func() {
        for {
            select {
            case msg1 := <- c1: // if received from channel 1
                fmt.Println(msg1)
            case msg2 := <- c2: // if received from channel 2
                fmt.Println(msg2)
            }
        }
    }()
    
}

func TimeoutOnSelect() {
    c1 := make(chan string)
    c2 := make(chan string)
    
    // start first goroutine
    go func() { // anonymous function
        for {
            c1 <- "sent from channel 1"
            time.Sleep(time.Second * 2)
        }
    }()
    
    // start second goroutine
    go func() { // anonymous function
        for {
            c2 <- "sent from channel 2"
            time.Sleep(time.Second * 3)
        }
    }()

    // start third goroutine that randomly picks a channel
    // will block if nothing received
    go func() {
        for {
            select {
            case msg1 := <- c1: // if received from channel 1
                fmt.Println(msg1)
            case msg2 := <- c2: // if received from channel 2
                fmt.Println(msg2)
            case msg3 := <- time.After(time.Second): // creates a channel and sends current time
                fmt.Println("timeout", msg3)
            //default:
            //    fmt.Println("no ready channel")
            //    time.Sleep(time.Second * 3)
            }
        }
    }()
    
}

func main() {
    //SimpleSelect()
    TimeoutOnSelect()
    pause()
}