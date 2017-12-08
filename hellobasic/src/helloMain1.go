package main

import "fmt"
import "hello/mathpackage"

func main() {
    numbers := []float64{1, 2, 3, 4}
    avg := math.Average(numbers)
    fmt.Println("Numbers:", numbers)
    fmt.Println("Average:", avg)
}