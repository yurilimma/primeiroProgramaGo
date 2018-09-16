package main

import (
	"fmt"
)

func main() {

	/**
	em go arrays possuem uma estrutura homogênia(mesmo tipo)
	e estática (fixo, não pode alterar o valor)
	**/

	var notas [3]float64
	fmt.Println(notas)
	total := 0.0
	//len = tamanho dos arrays
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	/**
	cast para dividir float / float,
	len retorna um inteiro porque
	logicamente não existe um array e meio
	**/
	media := total / float64(len(notas))
	fmt.Printf("media: %.2f \n ", media)

}
