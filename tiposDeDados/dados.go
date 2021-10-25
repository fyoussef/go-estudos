package main

import (
	"errors"
	"fmt"
)

func main() {

	// ============================= NUMEROS INT =============================

	// Existem 8 tipos de números inteiros no GO
	// int8, int16, int32, int64
	// int8 iŕa suportat números com 8bits
	// int16 iŕa suportat números com 16bits
	// ....

	var numero1 int8 = 10

	fmt.Println(numero1)

	// Tmb é possível usar somente o tipo int
	var numero2 int = -200
	fmt.Println(numero2)

	// o tipo int irá usar os bits do seu próprio computador como base
	// Por Ex: computardor 64bits -> int64
	//         computador  32bits -> int32

	// O tipo int é assumido quando é declarado uma var implicita
	numero3 := 50
	// var numero3 int = 50 <- Seria a mesma coisa
	fmt.Println(numero3)

	// No GO tamem existe o uint
	// Significa unsygnedint = um tipo de dado int mas sem sinal
	// uint8, uint16, uint32, uint64

	// Ele só aceitará ser declarado caso o numero não possua um sinal
	// EX: Valor negativo (-100)
	var numero4 uint = 200
	fmt.Println(numero4)

	// Alguns desses tipos de dados aceitam tambem um ALIAS (apelido)

	// INT32 = RUNE
	var numero6 rune = 23
	fmt.Println(numero6)

	// UINT8 = BYTE
	var numero7 byte = 5
	fmt.Println(numero7)

	// ============================= FIM NUMEROS INT =============================

	// ============================= NUMEROS REAIS (FLOAT) =============================

	// Numero float segue a mesma lógica mas só possuem 2 tipos
	// float32, float64
	// A diferença é que esse tipo aceita numeros quebrados (com ponto)

	var numeroReal1 float32 = 12.5678
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 15.55555666777
	fmt.Println(numeroReal2)

	// Com inferencia de tipos tambem funciona
	numeroReal3 := 15.55
	fmt.Println(numeroReal3)
	// Ele irá ter o espaço em bits de a cordo com o seu computador

	// ============================= FIM NUMEROS REAIS (FLOAT) =============================

	// ============================= STRINGS =============================

	// Váriaveis do tipo strings devem ser declaradas no GO somente com "" (Áspas Duplas)

	var str1 string = "String 1"
	fmt.Println(str1)

	// Declaração com inferência
	str2 := "String 2"
	fmt.Println(str2)

	// Declaração de string por '' (Áspas simples)
	// irá representar um CHAR

	char := 'B'
	fmt.Println(char)

	// Como resultado do Println ele irá retornar 66
	// Que representa 1 numero da TABELA ASCII
	// OBS: Este char é do tipo INT

	// ============================= FIM STRINGS =============================

	// ============================= VÁRIAVEL 0 =============================

	// Váriavel 0 no go é quando vc inicializa uma variavel sem atribuir um valor a ela
	// E elá irá retornar o valor inicial
	// Por Ex:

	var varStr0 string
	fmt.Println(varStr0)

	// Irá retornar vazio

	var varNum0 int64
	fmt.Println(varNum0)

	// Irá retornar 0

	// Ná declaração por inferencia não é possivel fazer isso pois vc deve OBRIGATORIAMENTE atribuir 1 valor a essa váriavel

	// ============================= FIM VÁRIAVEL 0 =============================

	// ============================= BOOLEANOS =============================

	// Booleanos no GO (true, false)
	// bool

	var boolean1 bool = true
	fmt.Println(boolean1)

	var boolean2 bool = false
	fmt.Println(boolean2)

	// Se declarar váriavel 0 no booleano o valor padrão é false
	var boolean bool
	fmt.Println(boolean)

	// ============================= FIM BOOLEANOS =============================

	// ============================= ERRO =============================

	// Erro é um tipo de dado no GO

	var error error
	fmt.Println(error)

	// Á váriavel 0 do tipo erro irá retornar <nil>
	// nil = zero ou nada

	// Para declarar um erro em uma váriavel no GO é necessário utilizar um pacote (errors)

	var errinho = errors.New("Error com o pacote Errors")
	fmt.Println(errinho)

	// Esse tipo é muito utilizado em GO
	// errors New <- Serve para criar um erro
	// OBS: Não é do tipo string

	// ============================= FIM ERRO =============================

}
