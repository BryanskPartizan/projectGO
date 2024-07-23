package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func summator(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		sum += i
	}
	return sum
}

func only_condition_loop() int {
	sum := 0
	for sum < 1000 { // init and post statements are empty
		sum += sum
	}
	return sum
}

func while_loop() int {
	sum := 1
	for sum < 1000 { // there are no init and post statements, only condition
		sum *= sum
	}
	return sum
}

func infinite_loop() int {
	for {
		fmt.Println("Infinite loop my dear")
	}
}

func sqrt(x float64) float64 {
	if x < 0 {
		return sqrt(-x)
	}
	return math.Sqrt(x)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v // v variable declared only in scope until the end of the if statement
	}
	return lim
}

func modified_pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func newton_sqrt(x float64) float64 {
	z := x / 2
	epsilon := 0.000000001
	for math.Abs(z*z-x) > epsilon {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func switcher_test() {
	fmt.Print("OS is ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("linux")
	case "darwin":
		fmt.Println("OS X")
	default:
		fmt.Println("чел установи нормальную ось кринж")
	}
}

func switcher_evaluation() {
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("ПЯТНИЦА ЕБАТЬ ЕГО В РОТ")
	case today + 1:
		fmt.Println("Страдать один день")
	case today + 2:
		fmt.Println("Страдать два дня")
	default:
		fmt.Println("И претерпевший до конца - спасётся")
	}
}

func switcher_with_no_condition() {
	now_time := time.Now()
	switch {
	case now_time.Hour() < 12:
		fmt.Println("Утро")
	case now_time.Hour() < 20:
		fmt.Println("День")
	default:
		fmt.Println("Ночь")
	}
}

func fibonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	// fmt.Println(summator(100))
	// fmt.Println(only_condition_loop())
	// fmt.Println(while_loop())
	// fmt.Println(sqrt(-10))
	fmt.Println(newton_sqrt(16))
	switcher_test()
	switcher_with_no_condition()
	fmt.Println(fibonacci(15))
}
