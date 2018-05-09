package main

/*
4. Crie uma função para calcular a quantidade dinheiro gasta por um fumante.
a função deve receber por parâmetro, o número de cigarros fumados diariamente,
a quantidade de anos que a pessoa fuma e o preço da carteira de cigarro e retorne
a quantidade de dinheiro gasto.
 */

import "fmt"

func bobMarley (cigarettes, years, price float32) float32 {
	priceByCigarettes := price / 20 // because a Wallet of cigarettes have a 20 cigarettes
	priceByDay := priceByCigarettes * cigarettes // price by day
	bobMarleySmoking := priceByDay * (years * 365) // price of all cigarettes, but not calculating the bisex year
	return bobMarleySmoking
}

func main() {
	var cig, year, price float32
	fmt.Println("Welcome to a game Show, how many cigarettes Bob Marley Smoking")
	fmt.Scanf("%g", &cig)
	fmt.Println("Years: ")
	fmt.Scanf("%g", &year)
	fmt.Println("Price:")
	fmt.Scanf("%g", &price)
	result := bobMarley(cig, year, price)
	fmt.Println("The Bob Marley pay",result, "$")
}
