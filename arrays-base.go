package main

import (
	"fmt"
)

func main() {
	var arrayOne [10]int

	arrayOne[0] = 10
	arrayOne[2] = 10

	fulledArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(fulledArray)
	fmt.Println(arrayOne)

	primes := fulledArray[0:3]
	fmt.Println(primes)

	lastValue := primes[len(primes)-1]
	firstValue := primes[0]

	fmt.Println(firstValue)
	fmt.Println(lastValue)
	fmt.Println(primes[len(primes)-1])

	planetList := [...]string{
		"PlanetOne",
		"PlanetTwo",
		"PlanetThree",
	}
	fmt.Println(len(planetList))
	arrayIteration(planetList)
	arrayRange(planetList)
}

func arrayIteration(arrayToIterate [3]string) int {
	indexSum := 0
	for i := 0; i < len(arrayToIterate); i++ {
		fmt.Println(i, arrayToIterate[i])
		indexSum += i
	}
	return indexSum
}

func arrayRange(arrayToIterate [3]string) int {
	indexSum := 0
	for i, arrayElement := range arrayToIterate {
		indexSum += i
		fmt.Println(i, arrayElement)
	}
	return indexSum
}
