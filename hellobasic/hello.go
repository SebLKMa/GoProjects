package main

import "fmt"

func appendSlices() {
    slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)
}

func copySlices() {
    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
}

func averageInSlices() {
    var total float64 = 0
    slice1 := []float64{1,2,3}
    
    for i:=0; i<len(slice1); i++ {
        total += slice1[i]
    }
    fmt.Println(slice1, total/float64(len(slice1)))
}

func makeMap() {
    m := make(map[string]int)
    m["four"] = 4
    m["two"] = 2
    fmt.Println(m)
    fmt.Println(m["two"])
    delete(m, "four")
    fmt.Println(m)
    
    if number, ok := m["two"]; ok {
        fmt.Println(number, ok)
    }
}

func main() {
    s := "hello go!"
    fmt.Println(s)
    
    appendSlices()
    copySlices()
    averageInSlices()
    makeMap()
}
