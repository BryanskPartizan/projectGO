package main

import "fmt"

type person struct {
	name   string
	age    int
	height float64
	weight float64
}

func main() {
	ivan := person{
		name:   "Ivan",
		age:    20,
		height: 150,
		weight: 100,
	}

	fmt.Println("Ivan", ivan)
	fmt.Println("Ivan.age", ivan.age)
	fmt.Println("Ivan.height", ivan.height)

	ivan.height = 100
	fmt.Println("Ivan.height", ivan.height)

	var ivanPointer *person = &ivan
	fmt.Println("IvanPointer", ivanPointer)

	ivanPointer.age = 99
	fmt.Println("IvanPointer.age", ivanPointer.age)
	fmt.Println("Ivan age", ivan.age)

	(ivanPointer).height = 100 // разыменовывани
	fmt.Println("ivan height", ivan.height)

	var bob *person = new(person) // инициализация указателяна структуру через присваивание адрес безымянного объекта
	fmt.Println("bob", bob)
}
