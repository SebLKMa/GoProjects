package main

import "fmt"
import m "hello/mathpackage"

func main() {
    numbers := []float64{1, 2, 3, 4}
    avg := m.Average(numbers)
    fmt.Println("Numbers:", numbers)
    fmt.Println("Average:", avg)
}