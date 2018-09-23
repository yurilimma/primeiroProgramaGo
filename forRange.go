package main

import "fmt"

func main() {
	/**cria um array com quantas posições
	forem passadas por parametro
	**/
	numeros := [...]int{1, 2, 3, 4, 5} //compilador conta

	/**retorna o indice e o elemento do array

	**/
	for i, numero := range numeros {
		fmt.Printf("%d) %d \n", i, numero)
	}
	/**
	como o go obriga a utilizar todas as variáveis definidas
	uma dica é colocar o _
	**/
	for _, numero := range numeros {
		fmt.Printf("valor elemento %d \n", numero)
	}
	//percorre todas as possições do array
	for numero := range numeros {
		fmt.Printf("valor elemento %d \n", numero)
	}

}
