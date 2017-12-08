package main

import "fmt"

func f1() {
    // add is local variable of the type func(int, int) int
    add := func(x, y int) int {
        return x + y
    }
    
    fmt.Println(add(4,2))
}

func f2() {
    x := 0 // closure - local increment func has access to in-scope variable x
    increment := func() int {
        x++
        return x
    }
    
    fmt.Println(increment()) // 1
    fmt.Println(increment()) // 2
}

// this function returns a (function that returns an uint)
func makeEvenNumberGenerator() func() uint {
    i := uint(0) // first value is 0, will be updated by the returned function object invocations
    return func() (val uint) {
        val = i  // return value references i
        i += 2   // i is always incremented by 2
        return
    }
}

func main() {
    f1()
    f2()

    nextEvenNum := makeEvenNumberGenerator() // variable nextEvenNum is declared-and-assigned as a function object type.
    fmt.Println(nextEvenNum()) // 0
    fmt.Println(nextEvenNum()) // 2
    fmt.Println(nextEvenNum()) // 4
}