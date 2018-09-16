package main

import (
	"fmt"
)

func main() {
	/**
	um map dentro do outro
	um definindo a letra dos funcionarios
	o outro definindo o nome e salario dos funcionários daquela letra
	**/
	funcsESalarios := map[string]map[string]float64{
		"J": {
			"Jose": 111.20,
			"João": 112.20,
		},
		"P": {
			"Pedro ": 122.30,
			"Paulo":  112.20,
		},
	}

	delete(funcsESalarios, "K")

	for letra, funcs := range funcsESalarios {
		fmt.Println(letra, funcs)
	}
}
