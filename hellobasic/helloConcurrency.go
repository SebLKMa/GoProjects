package main

import ("fmt"; "time"; "math/rand")

func f(max int) {
    for i := 0; i<max; i++ {
        fmt.Println(max, ":", i)
    }
}

// goroutines are lightweight
func go10() {
    for i := 0; i<10; i++ {
        go f(i)
    }
}

func fsleep(max int) {
    for i := 0; i<max; i++ {
        fmt.Println(max, ":", i)
        seconds := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * seconds)
    }
}

// goroutines are lightweight
func go10sleep() {
    for i := 0; i<10; i++ {
        go fsleep(i)
    }
}

func main() {
    //go f(10)
    
    //go10()
    
    go10sleep()
    
    var input string
    fmt.Scanln(&input) // pause main goroutine
}