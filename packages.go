package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	BIG_NUMBER   = 1 << 100         // 100 Bit
	SMALL_NUMBER = BIG_NUMBER >> 99 // 1 bit
)

func need_int(x int) int {
	return x*10 + 1
}

func need_float(x float64) float64 {
	return x * 0.1
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (number_one, number_two int) {
	number_one = sum * 5 / 9
	number_two = sum - number_one
	return
}

func main() {
	const KIPLING_IF = "if you can keep your head when all about you"
	int_number := 43
	float_number := float64(int_number) + 0.0001
	complex_number := 0.1 + 0.12i

	fmt.Println(need_int(10))
	fmt.Println(need_float(10))
	fmt.Println(float_number)
	fmt.Println(complex_number)
	fmt.Println("Some random number: ", rand.Intn(10))
	fmt.Println("Square root from 271 is ", math.Sqrt(271))
	fmt.Println("Math Pi = ", math.Pi)
	fmt.Println(add(55, 100))
	str_one, str_two := "first", "second"
	fmt.Println(swap(str_one, str_two))
	fmt.Println(split(110))

}
