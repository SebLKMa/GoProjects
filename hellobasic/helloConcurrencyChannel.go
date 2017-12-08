package main

import ("fmt"; "time")

// bidirectional channel
func pinger(c chan string) {
    for i := 0; ; i++ {
        c <- "ping" // send string to channel, blocks until channel receives
    }
}

// eg. of send-only channel
func ponger(c chan<- string) {
    for i := 0; ; i++ {
        c <- "pong" // send string to channel, blocks until channel receives
    }
}

// you can set device as receive-only by <-chan
func device(c chan string) {
    for {
        msg := <- c // receive string from channel
        fmt.Println("Received:", msg)
        time.Sleep(time.Second * 1)
    }
}

func pause() {
    var input string
    fmt.Scanln(&input) // pause
}

func main() {
    var c chan string = make(chan string)
    
    go pinger(c)
    go ponger(c)
    go device(c)
    
    pause()
}