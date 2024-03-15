package main

import (
	"fmt"
)

func calculadora(valor1 int, valor2 int, operacao int) int {
	var resultado int

	if operacao == 1 {
		resultado = valor1 + valor2
	} else if operacao == 2 {
		resultado = valor1 - valor2
	} else if operacao == 3 && valor2 != 0 {
		resultado = valor1 / valor2
	} else if operacao == 3 && valor2 == 0 {
		fmt.Println("Dividendo não pode ser zero, tente novamente")
		return 0
	} else if operacao == 4 {
		resultado = valor1 * valor2
	}
	return resultado
}

func main() {
	var valor1, valor2, operacao int

	fmt.Print("Digite o primeiro valor ")
	fmt.Scanf("%d", &valor1)
	fmt.Print("Digite o segundo valor ")
	fmt.Scanf("%d", &valor2)
	fmt.Print("Digite a operação desejada: 1 - adição, 2 - subtração, 3 - divisão, 4 - multiplicação: ")
	fmt.Scanf("%d", &operacao)

	resultado := calculadora(valor1, valor2, operacao)
	fmt.Println(resultado)
}
