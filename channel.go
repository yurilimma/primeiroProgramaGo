package main

import (
	"fmt"
)

/**
canal é bom para trabalhar com goroutines
já que faz o programa virar sincrono, esperando as informações do canal antes de finalizar o progroma(thread principal - main)
**/
func main() {
	/**
		canal é um tipo, intrisico da linguagem.
		similar ao inteiro, float etc
		dentro dele podem trafegar dados etc

	**/
	/**
	um canal e dentro desse canal
	terá valores inteiros, o segundo parametro
	é referente ao buffer
	**/
	ch := make(chan int, 1)

	//enviando 1 para o ch
	ch <- 1

	//lendo canal, não é obrigado a atribuir a nenhuma variavel
	<-ch

	ch <- 2
	fmt.Println(<-ch)
}
