package main

import (
	"fmt"
)

/**
	Jogando o retorno de uma função
	declarando e escrevendo a mesma
	diretamente	em uma variável

	obs: funciona mesmo a variavel
	estando fora da função main
**/
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
