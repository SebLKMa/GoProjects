package main

import (
	"fmt"
)

func main() {
	defer func() {
		str := recover()
		fmt.Println("Exception caught: " + str.(string))
	}()

	slice := []int{10, 20, 30, 40, 50}

	// length 2, capacity 4
	newSlice := slice[1:3]

	newSlice[1] = 35

	fmt.Println(slice[2])
	fmt.Println(newSlice[1])

	//newSlice[2] = 35 // index out of range

	newSlice = append(newSlice, 60)
	fmt.Println(newSlice)
	fmt.Println(newSlice[2])
	fmt.Println(slice[3])
}
