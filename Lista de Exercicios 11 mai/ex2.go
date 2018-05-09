package main

import "fmt"

func media(v1, v2, v3 float32) {
	media:= (v1 + v2 + v3) / 3
	switch media != 0 {
	case media >= 7:
		fmt.Println("Aprovado")
	case media >= 5:
		fmt.Println("Recuperação")
	default:
		fmt.Println("Reprovado")
	}
}

func main() {
	var v1, v2, v3 float32
	fmt.Scanf("%g", &v1)
	fmt.Scanf("%g", &v2)
	fmt.Scanf("%g", &v3)
	media(v1, v2, v3)
}
