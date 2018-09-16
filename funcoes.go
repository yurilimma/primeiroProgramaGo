package main

import (
	"fmt"
	"strconv"
)

/**
	retorno das funções é declarado
	após os parâmetros de entrada.
**/
func somar(a int, b int) int {
	c := a + b
	return c
}

func imprimir(valor int) {
	fmt.Println(strconv.Itoa(valor))
}

func main() {
	var c = somar(2, 2)

	imprimir(c)
}
