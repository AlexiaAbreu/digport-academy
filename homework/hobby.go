package main

import "fmt"

type DadosPessoais struct {
	nome  string
	hobby string
}

func main() {

	dados := []DadosPessoais{
		{"Alexia", "Dançar"},
		{"Rafaela", "Desenhar"},
		{"Viviane", "Patinar"},
	}

	for _, dado := range dados {
		fmt.Println("O hobby da " + dado.nome + " é " + dado.hobby)
	}
	// for i := 0; i < len(dados); i++ {
	// 	fmt.Println("O hobby da " + dados[i].nome + " é " + dados[i].hobby)
	// } outra forma de fazer

}
