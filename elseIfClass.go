package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite seu nome:")
	fmt.Scanf("%d", &numero)

	if numero > 0 {
		fmt.Println("positivo")
	} else if numero < 0 {
		fmt.Println("negativo")
	} else {
		fmt.Println("zero")
	}
}
