package main

import "fmt"

func main() {
	contador := 10

	for contador != 0 {
		fmt.Print("contagem regressiva - ")
		fmt.Println(contador)
		contador--
	}
	fmt.Println("arrasou!")
}
