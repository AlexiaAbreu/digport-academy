package main

import "fmt"

func main() {
	var listaDespesas = []string{"mercado", "roupas", "casa", "água", "luz", "streamming", "faculdade"}

	fmt.Print("Digite a despesa que você deseja verificar: ")

	var conta string
	fmt.Scanln(&conta)

	for _, despesa := range listaDespesas {
		if despesa == conta {
			fmt.Println("sua depesa '" + despesa + "' já está na lista")
		}
	}

	fmt.Printf("O total de despesas a pagar é de %d itens \n", len(listaDespesas))

	mapDespesas := make(map[string]int)

	mapDespesas["mercado"] = 600
	mapDespesas["luz"] = 100
	mapDespesas["água"] = 200
	mapDespesas["casa"] = 1500
	mapDespesas["roupas"] = 250

	var totalDespesas int

	for _, despesa := range mapDespesas {
		totalDespesas += despesa
	}
	if totalDespesas > 2000 {
		fmt.Println("o valor total de despesas é maior do que o orçamento de R$2000")
	}

	fmt.Printf("O valor total de despensas é de %d", totalDespesas)

}
