package main

import (
	"fmt"
)

/**
definindo dois retornos, um int e um error
**/
func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		//formata a mensagem de erro
		return -1, fmt.Errorf("Número inválido")
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil

	}
}

func main() {
	//dois retornos
	resultado, _ := fatorial(3)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

}
