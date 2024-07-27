package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	arrayA := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceA := arrayA[0:len(arrayA)]
	fmt.Println(sliceA)

	arrayOfNames := [4]string{
		"NameOne",
		"NameTwo",
		"NameThree",
		"NameFour",
	}

	fmt.Println(arrayOfNames[0:2])
	arrayOfNames[0] = "NewName"
	fmt.Println(arrayOfNames)
	fmt.Println(arrayOfNames[len(arrayOfNames)-1])

	// booleanArray := [3]bool{true, false, true}
	// indexArray := [3]int{1, 2, 3}

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)

	sliceOne := []int{2, 3, 5, 7, 11, 13}
	printSlice(sliceOne)

	// Slice the slice to give it zero length.
	sliceOne = sliceOne[:0]
	printSlice(sliceOne)

	// Extend its length.
	sliceOne = sliceOne[:4]
	printSlice(sliceOne)

	// Drop its first two values.
	sliceOne = sliceOne[2:]
	printSlice(sliceOne)

	var nilSlice []int
	printSlice(nilSlice)

	dynamicArrayOne := make([]int, 3)
	printSlice(dynamicArrayOne)

	dynamicArrayOne[:4] = 4
	printSlice(dynamicArrayOne)

}
