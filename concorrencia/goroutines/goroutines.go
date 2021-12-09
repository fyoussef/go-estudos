package main

import (
	"fmt"
	"time"
)

func main() {

	// CONCORRÊNCIA != PARALELISMO

	// Paralelismo acontece quando mais de uma tarefa está sendo executada ao mesmo tempo
	// Isso só é possível se o processador tiver mais de um núcleo
	// Pq só assim ele executa mais de uma tarefa dividindo elas em cada nucleo possibilitando serem executadas ao mesmo tempo

	// Já a concorrencia não necessariamente roda simutaneamente com outra tarefa apesar fazer isso em apenas 1 núcleo
	// ela não espera a outra tarefa terminar para ser executada

	// Para ilustrar como funciona a concorrência

	go escrever("TESTE")
	// Essa função irá ficar escrevendo TESTE no terinal infinitamente

	// Em uma linguagem sem concorrencia como de costume ele irá executar primeiro a função que está escrevendo "TESTE"
	// e irá esperar o termino da primeira fução pra começar a execução da proxima
	// Como a função está em um laço infinito ela nunca irá parar de executar
	// E uma forma para contornar isso é com a GO ROUTINE

	// A GO ROUTINE é um metodo que é declarando usando a palavra reservada go
	// Executando a go routine é como se vc tivesse dizendo ao seu código
	// Vai executa o código e não precisa esperar esta função terminar

	escrever("TESTE 2")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
