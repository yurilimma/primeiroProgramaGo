package main

import (
	"fmt"

	"github.com/yurilimma/html"
)

//origem, como é <-, é somente leitura
func encaminhar(origem <-chan string, destino chan string) {
	for {
		//pegando o valor de origem e jogando em destino
		destino <- <-origem
	}
}

//multiplexar - misturar(mensagens) num único canal

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://youtube.com", "https://youtube.com"),
		html.Titulo("https://google.com", "https://gmail.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)

}
