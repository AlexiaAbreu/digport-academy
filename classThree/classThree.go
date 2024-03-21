package main

import "fmt"

func main() {
	type Coordenada struct {
		latitude  float64
		longitude float64
	}

	coordenada := Coordenada{latitude: 32.98, longitude: -23.98}

	fmt.Println(coordenada)
	fmt.Println(coordenada.latitude)
	fmt.Printf("%+v", coordenada)
}
