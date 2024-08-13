package main

import "fmt"

func main() {
	var stringSlice []string
	stringSlice = append(stringSlice, "string_one", "string_two")
	fmt.Println(stringSlice)
	address := &stringSlice
	fmt.Println(&stringSlice)
	fmt.Println(address)

	makeSliceOne := make([]string, 3, 4)
	makeSliceTwo := make([]int, 3)
	makeSliceThree := make([]float64, 10)
	fmt.Println(makeSliceOne, makeSliceTwo, makeSliceThree)
	fmt.Printf("Address of array before append: %p\n", &makeSliceOne)

	makeSliceOne = append(makeSliceOne, "text")
	fmt.Printf("Address of array after append: %p\n", &makeSliceOne)
	makeSliceOne = append(makeSliceOne, "text", "text", "text", "text")
	fmt.Printf("Address of array after out of length: %p\n", &makeSliceOne)

	fmt.Println("Length of the array:", len(makeSliceOne))
	fmt.Println("Capacity of the array:", cap(makeSliceOne)) // doubled initial value
}
