package main

import (
	"fmt"
)

func trocar(p1, p2 int) (segundo int, primeiro int) {
	/**
	não precisou do : antes do = porque as mesmas são variáveis
	de retorno, logo, já foram inicializadas
	**/
	segundo = p1
	primeiro = p2

	/**
	Como já foi atribuido as variáveis de retorno
	ao utilizar só o return, ele por default já
	retorna as mesmas
	**/
	return
}

func main() {
	x, y := trocar(2, 3)

	fmt.Println(x, y)
}
