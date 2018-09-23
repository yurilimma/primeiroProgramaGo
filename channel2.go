package main

import (
	"fmt"
	"time"
)

/**
channel(canal) - é a forma de comunicação entre goroutines

**/
func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)

	c <- 3 * base

	time.Sleep(3 * time.Second)

	c <- 4 * base
}

func main() {
	//pra criar canal tem que criar com make()
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c
	fmt.Println(a, b)
	//só printa o c quando o dado estiver pronto, no caso a ultima operacao da fnção acima, 8
	fmt.Println(<-c)

	//fmt.Println(<-c) não consegue executar duas vezes, gera deadlock já que não está esperando receber nada
}
