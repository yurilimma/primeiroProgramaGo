package main

import (
	"fmt"
)

func main() {
	/**make : cria um array baseado em um slice
	primeiro parâmetro é o tipo do slice
	segundo é o tamanho do slice
	**/
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	/**cria um array com 10 elementos,
		associa 10 das 20 posições do slice ao array
		cap = capacidade do slice, imprime 20
	**/
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	/**
	adicionanado 3 elementos ao array
	funciona pois a capacidade máxima do slice submetido
	ao array é 20 e 10 + 3 é igual a 13
	**/
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	/**
	Se atingir a capacidade máxima do slice,
	no caso 20, ele internamente vai dobrar o tamanho do slice interno do array
	**/
}
