package main

import "fmt"

func main() {
	i, j := 10, 20

	p := &i //pointer to i
	/*
		& - define variable address in memory. Address is a place in memory
	*/
	fmt.Println(&i)
	fmt.Println(&*p) // the same address

	fmt.Println(*p)

	*p = 100 // i become 100 because we change the value in memory reserved for i through the pointer
	fmt.Println(i)
	fmt.Println(&i)

	p = &j
	*p = *p / 10 // j changed because we operate with the memory reserved for j through the pointer
	fmt.Println(j)
	fmt.Println(*p)
}
