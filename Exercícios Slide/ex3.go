package Exercícios_Slide

import "fmt"

func sliceN (array []int, numberSlice int) int{
	return array[numberSlice -1]
}

func main() {
	var someIntSlice []int = []int{1,2,3,4,5,6,7,8}
	var numberSlice = 3
	var result = sliceN(someIntSlice, numberSlice)
	fmt.Println(result)
}
