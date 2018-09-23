package main

import (
	"fmt"
)

func rotina(ch chan int) {

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou!")
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)

	go rotina(ch)
	/**
	o buffer está definido como 3
	ao entrar mais de 3 vezes com informação no canal
	ele só aceitará o 4 quando ler alguma informação do mesmo,
	e nesse caso ele só lê aqui no main,
	não imprimirá ("Executou") pois quando ele libera o buffer
	já está na função main
	**/
	fmt.Println(<-ch)

}
