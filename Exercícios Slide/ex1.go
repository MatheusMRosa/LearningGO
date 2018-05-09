package Exercícios_Slide

import "fmt"

func sliceLast (word []string) string{
	var size = len(word) -1
	var letter = word[size]
	return letter
}

func main() {
	var someSlice []string = []string{"a","b","c","d","e","f"}
	var result = sliceLast(someSlice)
	fmt.Println(result)
}