package main

import "fmt"

func reverse_string(some_string string) string {
	for i := 0; i < len(some_string); i++ {
		defer fmt.Print(string(some_string[i]))
	}
	return some_string
}

func main() {
	/*
		fmt.Println("counting")
		for i := 0; i < 5; i++ {
			defer fmt.Println(i)
		}
		fmt.Println("done")
	*/

	reverse_string("abcdefg")
}
