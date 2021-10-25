package main

import "fmt"

func main() {

	// ====================== ARITIMETICOS ======================

	// + <- Soma
	// - <- Subtração
	// / <- Divisão
	// * <- Multiplicação
	// % <- Módulo

	// OBS: Esses operadores não funcionam se o tipo de dados não forem o mesmo
	// Exemplo : int8 + int8, int8 - int8, int8 / int8...

	// ====================== FIM ARITIMETICOS ======================

	// ====================== ATRIBUIÇÃO ======================

	// Operadores de atribuição explicita ou por inferencia

	var str string = "Váriavel de Atribuição Explicita"

	str1 := "Váriavel de Atribuição por Inferência"

	fmt.Println(str)
	fmt.Println(str1)

	// ====================== FIM ATRIBUIÇÃO ======================

	// ====================== RELACIONAIS ======================

	// > <- Maior
	// < <- Menor
	// >= <- Maior que
	// <= <- Menor que
	// != <- Diferente
	// == <- Igual

	// ====================== FIM RELACIONAIS ======================

	// ====================== LÓGICOS ======================

	// No GO só existem 3 operadores Lógicos

	// && <- AND
	// || <- OR
	// !  <- NOT

	// ====================== FIM LÓGICOS ======================

	// ====================== UNÁRIOS ======================

	// ++
	// +=
	// --
	// -=

	incremento := 100
	incremento++

	fmt.Println(incremento)

	incremento1 := 10
	incremento1 += 15

	fmt.Println(incremento1)

	// OBS: No GO não pode fazer os operadores antes de um numero
	// Exemplo: ++1 OU --1
	// O GO não executa esses operadores

	// ====================== FIM UNÁRIOS ======================

	// OBS: NO GO NÃO EXISTE OPERADOR TERNARIO
}
