package main

import "fmt"

func somar(num1 int8, num2 int8) int8 { // <-  Se a função tiver um retorno, vc pode declarar o retonro entre os parenteses e a chave
	return num1 + num2
}

// =========================== ATENÇÃO ===========================

// No GO as funções podem ter mais de 1 retorno

// Exemplo:

// No GO caso os parâmetros sejam do mesmo tipo eles podem ser declarados dessa forma
// Para declarar os tipos dos valores que serão retornado (caso seja mais de um retorno) abrir novamente o parenteses e declara o tipo dos retorno
func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {

	soma := somar(10, 10)
	fmt.Println(soma)

	// Tambem é possivel declarar funções em variaveis

	var funcOla = func(str string) {
		fmt.Println(str)
	}

	funcOla("Olá da função")

	// Se eu quiser o retorno dessa função

	var funcOla1 = func(str string) string { // <- Devo especificar oque ela irá retornar
		return str
	}

	retornoDaFunc := funcOla1("Valor que quero retornar")

	fmt.Println(retornoDaFunc)

	// Para pegar o retorno de uma função que tem mais de 1 retorno deve fazer da seguinte forma

	vlrSoma, vlrSubtracao := calculos(10, 10)

	fmt.Println("O valor da soma dos numeros é:", vlrSoma)
	fmt.Println("O valor da subtração dos numeros é:", vlrSubtracao)

	// Se a função retornar mais de 1 valor e não tiver a quantidade certa para "pegar" o retorno desses valores irá dar erro

	// Caso vc queira apenas utilizar 1 retorno da função, vc deve passar o UNDERLINE na posição do retorno que vc não quer utilizar
	// Por Exemplo:

	vlrSoSoma, _ := calculos(15, 20)
	fmt.Println("Somente o valor dá soma é:", vlrSoSoma)
}
