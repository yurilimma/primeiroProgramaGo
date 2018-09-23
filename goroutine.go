package main

import (
	"fmt"
	"time"
)

/**
concorrencia: concorrencia é a capacidade de administrar multiplas tarefas,
pode ocorrer no mesmo processador. Exemplo: escalonadores;


paralelismo: fazendo duas coisas ao mesmo tempo. Só é possível se
tiver multiplos processadores

Go é concorrente

Paralelismo exige muito mais do S.O e do hardware
**/

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	///fale("Maria", "Porque voce não fala comigo", 3)
	//fale("João", "Só posso falar depois de você", 1)
	/**
	colocando go antes, ele executa uma goroutine
	similar ao conceito de thread
	Obs.: só funciona até o final do thread principal
	no caso, a main, se a thread principal acabar ele para de excutar as novas threads
	**/
	go fale("Maria", "ei...", 500)
	go fale("João", "Opa...", 500)

}
