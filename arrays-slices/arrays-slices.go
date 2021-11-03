package main

import (
	"fmt"
	"reflect"
)

func main() {

	// ================================== ARRAY ==================================

	// Array no go como em outras linguagens é uma lista de valores

	// NO GO tem 2 maneiras de declarar uma variável de array

	// var array[tamanho do array(1,2,3..)] tipo (int, string, uint...)
	// No GO todos os dados dentro do array devem ser do mesmo tipo
	// Não pode ter por exemplo uma string um int e etc..
	var array1 [5]int // Array de valor 0

	fmt.Println(array1)
	// [0 0 0 0 0]

	// Uma formar de colocar os valores no array é da seguinte maneira

	array1[1] = 5

	fmt.Println(array1)
	// [0 5 0 0 0]

	// Para declarar um array usando a inferência de tipos deve se fazer da seguinte forma

	array2 := [5]string{"Posição1", "Posição2"}

	fmt.Println(array2)
	// [Posição1 Posição2		] os espaços que não receberam valores retornam vazios

	// Caso queira adicionar mais uma posiçao em um array
	// array2[5] = "Posição 6" <= ele irá dar erro pois o array declarado só permite 5 posições

	// Caso queira declaram um array com posiçoes mais "flexiveis" deve se fazer da seguinte forma

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(array3)
	// [1, 2, 3, 4, 5, 6, 7]

	// Caso queira ainda adicionar mais uma posição em um array declarado desta forma

	// array3[7] = 8 <=  irá dar um erro da mesma forma

	// Definir um array desta forma [...] só irá aceira a quantidade de posições que forem atribuidas a variavel inicialmente

	// Ou seja no GO a quantidade de posições são fixas

	// ================================== FIM ARRAY ==================================

	// ================================== SLICE ==================================

	// Slice é como o array  mas NÃO É UM ARRAY
	// Ele funciona como um ponteiro para um array
	// Ele não precisa de um valor fixo ou seja, ele irá suportar a quantidade de posições de acordo com suas necessidades
	// Exemplo de declaraçao do slice

	// Da mesma forma que no array ele tambem é tipado
	slice := []int{4, 456, 456, 887, 7, 87, 78}

	fmt.Println(slice)
	// [4 456 456 887 7 87 78]

	// Para ver os tipos do slice e do array

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))
	// slice  = []int
	// array3 = [7]int

	// Essa diferença significa que são de tipos diferentes
	// Ou seja não se pode fazer operações entre eles pois o GO é fortemente tipado e eles são de tipos diferentes

	// Para adicionar um item no slice usa a função append
	// Ele irá adicionar o item no slice e retornar um novo com o item adicionado

	sliceNovo := append(slice, 777)
	// a func append recebe 2 argumentos. 1º o array 2º o novo item que será adicionado no novo array

	fmt.Println(sliceNovo)
	// [4 456 456 887 7 87 78 777]

	// Caso queira somente adicionar uma nova posição em um slice já existente se faz da mesma forma

	slice = append(slice, 123)

	fmt.Println(slice)
	// [4 456 456 887 7 87 78 123]

	// ================================== FIM SLICE ==================================
}
