package Exercícios_Slide

import "fmt"

func slicepenultimate (word []string) string{
	var size = len(word) - 2
	var letter = word[size]
	return letter
}

func main() {
	var someSlice []string = []string{"a","b","c","d","e","f"}
	var result = slicepenultimate(someSlice)
	fmt.Println(result)
}
