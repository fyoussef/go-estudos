package main

import "fmt"

// =============================== STRUCT ===============================

type usuario struct {
	nome      string   // <- O nome será OBRIGATORIAMENTE uma STRING
	sobreNome string   // <- O sobre nome será OBRIGATORIAMENTE uma STRING
	idade     uint8    // <- A idade será OBRIGATORIAMENTE sem sinais (UINT)
	endereco  endereco // <- O endereço seŕa com base no Struct de endereco
}

// Também é permitido um Struct dentro de outro Struct

type endereco struct {
	rua    string
	numero uint8
}

func main() {

	// Declarar a variavel com o Struct sem valor nenhum
	// Ele irá atribuir o valor 0 para todos os campos do Struct
	var user usuario
	fmt.Println(user)

	endereco := endereco{"Rua dos Bobos", 0}

	// Para atribuir valor de acordo com os campos do Struct
	var user1 usuario

	user1.nome = "Filipi"
	user1.sobreNome = "Youssef"
	user1.idade = 20
	user1.endereco = endereco // O campo endereco do Struct de usuario ira receber os valores da variavel endereco

	fmt.Println(user1)

	// Uma outra forma de declarar valores a 1 Struct é usando a inferência de tipos
	user2 := usuario{"Felps", "Youssef", 20, endereco}
	// user2   <- váriavel
	// usuario <- Struct
	// {}      <- Campo que irá receber os valores para os campos do Struct

	fmt.Println(user2)

	// OBS: Com esses formatos já vistos sempre declarar os valores do campo do Struct na ordem correta do Struct e passando todos os valores necessários

	// Caso queira passar somente 1 campo do struct fazer da seguinte forma

	user3 := usuario{nome: "Filipi"}

	fmt.Println(user3)

	// Dessa formar ele irá aceitar somente o campo fornecido e irá atribuir 0 para os outros campos

}
