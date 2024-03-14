package main

import "fmt"

func main() {
	// var nomeDoUsuario string

	// fmt.Print("Digite seu nome:")
	// fmt.Scanln(&nomeDoUsuario)
	// fmt.Println("bem vinda", nomeDoUsuario)

	var valor int // declaração
	valor2 := 1   // atribuição

	fmt.Printf("Digite um número: ")
	fmt.Scanf("%d", &valor)
	soma := valor + valor2

	fmt.Println("soma é ", soma)
}
