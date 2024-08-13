package main

import (
	"fmt"
	"strings"
)

func main() {
	//kipling_string := "if you can keep your head when all about you"

	/*
		for i := 0; i < len(kipling_string); i++ {
			fmt.Printf("%c", kipling_string[i])
		}
	*/

	first_row := "first row"
	second_row := "second row"

	// string compare
	fmt.Println(strings.Compare(first_row, second_row))
	fmt.Println(strings.Compare(first_row, first_row))
	fmt.Println(strings.Compare(second_row, first_row))

	testString := "String with some words"
	test := strings.Fields(testString)
	fmt.Println(test)

}
