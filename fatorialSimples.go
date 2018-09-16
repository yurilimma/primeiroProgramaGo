package main

import (
	"fmt"
)

/**
O uint define que o parâmetro de entrada
necessáriamente precisa ser um inteiro
**/
func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)

	}
}

func main() {
	//dois retornos
	resultado := fatorial(3)
	fmt.Println(resultado)
	/**não compila porque o fatorial
	está definido como parâmetro de entrada
	o uint(inteiro positivo)
	**/
	resultado2 := fatorial(-4)

	fmt.Println(resultado2)

}
