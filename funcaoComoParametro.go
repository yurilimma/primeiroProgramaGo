package main

import (
	"fmt"
)

func multiplicacao(a, b int) int {
	return a * b
}

/**
recebe uma função chamada função por parametro
a função recebida por parâmetro tem
que declarar os parâmetros de entrada
e saida da mesma [func(int int)int]
e recebe também mais 2 parâmetros adicionais(p1,p2)
retorna um inteiro
**/
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 2)
	fmt.Println(resultado)
}
