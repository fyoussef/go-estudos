package main

import "fmt"

func main() {

	// Declaração de var explícito, especificando o tipo de dado
	var variavel string = "Váriavel 1"
	fmt.Println(variavel)

	// Declaração de var implicita, sem especificar o tipo de dado
	// Conhecido tbm como inferência de tipo
	variavel2 := "Váriavel 2"
	fmt.Println(variavel2)

	// No GO tambem é possível declarar mais de uma váriavel explícita
	var (
		primeira string = "Primeira String"
		segunda  string = "Segunda String"
	)

	fmt.Println(primeira, segunda)

	// Tbm é possível declarar mais de uma váriavel usando a inferência de tipos
	terceira, quarta, quinta := "Terceira String", "Quarta String", "Quinta String"

	fmt.Println(terceira, quarta, quinta)

	// Para inverter valores de váriaveis no GO
	terceira, quarta = quarta, terceira
	fmt.Println(terceira, quarta)
}
