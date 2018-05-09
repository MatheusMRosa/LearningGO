package main

import "fmt"

func media(v1, v2, v3 float32) (string, float32) {
	media:= (v1 + v2 + v3) / 3
	switch media != 0 {
	case media >= 7:
		return "Aprovado", media
	case media >= 5:
		return "Recuperação", media
	default:
		return "Reprovado", media
	}
}

func main() {
	var v1, v2, v3 float32
	fmt.Scanf("%g", &v1)
	fmt.Scanf("%g", &v2)
	fmt.Scanf("%g", &v3)
	situation, media := media(v1, v2, v3)
	fmt.Println("O aluno está com",situation, "com média:", media)
}
