package main

import "fmt"

type Horse struct {
	height float64
	width  float64
	age    int
}

func main() {
	horseOne := Horse{height: 10, width: 20, age: 30}
	horseTwo := Horse{height: 10, width: 20, age: 30}
	fmt.Println(horseOne)
	fmt.Println(horseTwo)

	ageSum := horseOne.age + horseTwo.age
	fmt.Println(ageSum)

	horseOneAgePointer := &horseOne.age
	horseOne.age = 25
	fmt.Println(*horseOneAgePointer)
	fmt.Println(horseOneAgePointer)
	fmt.Println(horseOne.age)
}
