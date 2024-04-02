package main

import "fmt"

func main() {

	dados := map[string]string{
		"Alexia":  "Dançar",
		"Paula":   "Correr",
		"Viviane": "Cantar",
	}

	for key, value := range dados {
		fmt.Println("O hobby da " + key + " é " + value)
	}
}
