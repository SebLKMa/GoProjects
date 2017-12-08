package main

import "fmt"

func average(s []float64) float64 {
    total := 0.0 // float type
    
    for _, v := range s {
        total += v
    }
    return total/float64(len(s))
}

func f42() (int, int) {
    return 4, 2
}

/*
  variadic parameter
*/
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func main() {
    mySlice := []float64{98,93,77,82,83}
    fmt.Println(average(mySlice))

    x, y := f42();
    fmt.Println(x, y);
    
    fmt.Println(add(1,3,5))
}