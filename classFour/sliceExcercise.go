package main

import "fmt"

func main() {
	const hello = "Hello"

	slice := make([]string, 0, 4)

	slice = append(slice, hello, "Alexia", "Viviane", "Carol")

	secondNameIs := slice[2]

	fmt.Println(slice[0], secondNameIs)
}
