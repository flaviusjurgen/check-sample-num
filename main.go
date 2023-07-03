package main

import (
	"fmt"
	"math"
)

func main() {
	number := 17
	if isPrime(number) {
		fmt.Println(number, "является простым числом.")
	} else {
		fmt.Println(number, "не является простым числом.")
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
