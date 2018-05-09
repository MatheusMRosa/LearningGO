package main

import "fmt"

 func verifyBigger (v1, v2 int){
 	if v1 == v2 {
 		fmt.Println(v1, "=", v2)
	} else if v1 > v2 {
		fmt.Println(v1)
	} else {
		fmt.Println(v2)
	}
 }

func main() {
	var v1, v2 int
	fmt.Scanf("%d", &v1)
	fmt.Scanf("%d", &v2)
	verifyBigger(v1, v2)
}