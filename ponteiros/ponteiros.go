package main

import "fmt"

func main() {

	var n1 int = 10

	var n2 int = n1

	fmt.Println(n1, n2)
	// 10 10

	n1++
	fmt.Println(n1, n2)
	// 11 10

	// Após incrementar +1 na váriavel n1 ele irá alterar somente o valor da váriavel n1

	// Isso acontece pq quando vc atribui um valor a uma váriavel esse valor se torna uma CÓPIA
	// Consequentemente o valor da váriavel n2 é uma CÓPIA da váriavel n1

	// ============================ PONTEIRO ============================

	// O ponteiro é uma referência de memória
	// Ou seja ele não será a cópia de um valor como na variável, ele irá guardar o endereço de um valor na memoria do computador

	// Para declarar um ponteiro

	var numero1 int  // <= Váriavel
	var numero2 *int // <= Ponteiro
	fmt.Println(numero1, numero2)

	// numero1 = 0
	// numero2 = <nil> (null)

	numero1 = 100
	// numero2 = numero1 // <= Declarar um ponteiro dessa forma irá causar um erro de tipo int que é o tipo dá váriavel que está sendo atribuida ao ponteiro
	// Para atribuir valor a um ponteirno deve ser declarado da seguinte forma

	numero2 = &numero1 // <= Assim se atribui um valor a um ponteiro

	fmt.Println(numero1, numero2)

	// numero1 = 100
	// numero2 = 0xc0000140c0

	// Ele retorna esse numero pois o ponteiro irá guardar o endereço de memoria do valor atribuido a ele que no caso é a var numero 1
	// Basicamente o valor retornado nele é o endereço da variavel numero1
	// OU seja dentro desse endereço => 0xc0000140c0 está salvo o valor da var numero1 (100)

	// Caso queira saber o valor que está guardado no endereço do ponteiro deve se fazer através da desreferenciação

	fmt.Println(*numero2) // Dessa forma que é feita a desreferenciação

	// E ele irá retornar o valor que está no endereço (100)

	// MOTIVO DE USAR O PONTEIRO

	// Digamos que eu queira trocar o valor atribuido a váriavel numero1

	numero1 = 200

	fmt.Println(numero1, numero2)
	// O valor retornado é
	// numero1 = 200
	// numero2 = 0xc0000140c0

	// O endereço continua imutável ou seja mesmo mudando o valor da váriavel de referência o endereço do ponteirno continua o mesmo
	// Para tirar a prova real

	fmt.Println(*numero2)

	// numero2 = 200
}
