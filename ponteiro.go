package main

import (
	"fmt"
)

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// * da acesso ao valor do ponteiro
	// o valor é jogado em n
	*n++
}

func main() {
	n := 1
	//passagem por valor
	inc1(n)
	fmt.Println(n)

	//passagem por referencia
	inc2(&n)
	fmt.Println(n)
}
