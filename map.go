package main

import (
	"fmt"
)

func main() {

	//maps devem ser inicializados
	aprovado := make(map[int]string)

	aprovado[123456789] = "maria"
	aprovado[333] = "pedro"
	aprovado[222] = "ana"
	fmt.Println(aprovado)
	/**
	por default primeiro ele joga a chave
	pro primeiro elemento(cpf)
	e o valor do map joga pro segundo
	elemento(nome)
	**/
	for cpf, nome := range aprovado {
		fmt.Printf(" %s (CPF: %d) \n", nome, cpf)
	}

	//imprime só o valor do elemento 222
	fmt.Println(aprovado[222])
	//deleta um valor do map a partir de sua chave
	delete(aprovado, 222)

}
