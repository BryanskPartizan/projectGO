package main

import "fmt"

func main() {
	a := [2]int{1, 2}
	b := [4]int{1, 3, 4}
	c := [...]int{2, 3, 4, 5}

	fmt.Println(a, b)
	fmt.Println(c)

	initArray := [...]int{1, 2, 3}
	copyArray := initArray

	fmt.Printf("Address of initArray: %p\n", &initArray)
	fmt.Printf("Address of copyArray: %p\n", &copyArray)

}
