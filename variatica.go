package main

import (
	"fmt"
)

/** ... define que a quantidade de parâmetros de entrada não está definida,
	porém todos tem que ser do tipo float
**/
func media(numeros ...float64) float64 {
	total := 0.0

	//ignorando o primeiro valor pois não queremos o indice
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	//pode adicionar quantos valores quiser
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 3.1))
}
