package main

import "fmt"

func verifyPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func showPrimes () {
	var numbersPrimes []int
	for i:= 2; i <= 100; i++ {//start in 2 because 0 and 1 is not use for calculate a prime, because all number to divide for 1
		if verifyPrime(i) {
			numbersPrimes = append(numbersPrimes, i)
		}
	}
	fmt.Println(numbersPrimes)
}

func main() {
	showPrimes()
}
