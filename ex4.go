package main

func palindromoValidation (word string) bool{
	var reverse string;
	for i := len(word) -1 ; i >= 0; i--{
		reverse = reverse + string(word[i])
	}
	println("Palavra Original: ", word)
	println("Palavra Ao contrário: ", reverse)
	if reverse == word {
		return true
	} else {
		return false
	}
}

func main() {
	palindromo := "natan"
	var isPalindromo = palindromoValidation(palindromo)
	println(isPalindromo)
}
