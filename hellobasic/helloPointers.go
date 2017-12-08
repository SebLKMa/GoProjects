package main

import "fmt"

// this function receives a pointer to an integer and changes its value to 0
func zero(pNum *int) {
    *pNum = 0; // derefeence pointer then set its value to 0
}

// this function receives a pointer to an integer and changes its value to 1
func one(pNum *int) {
    *pNum = 1;
}

// almost the same as C++
func main() {
    x := 8
    fmt.Println(x);
    zero(&x); // address of x
    fmt.Println(x);
    
    // allocates memory of integer y pointer
    y := new(int)
    one(y);
    fmt.Println(*y);
    
    // y will be garbage-collected when nothing refers to it any more.
}