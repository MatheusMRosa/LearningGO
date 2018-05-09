package main

import "fmt"

func nextNumber (number int) int {
	return number + 1
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	next := nextNumber(num)
	fmt.Println("Next: ", next)
}