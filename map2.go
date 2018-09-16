package main

import (
	"fmt"
)

func main() {
	//chave é o que está entre [] no map
	funcsESalarios := map[string]float64{
		"Jose":     111.20,
		"Gabriela": 222.22,
		"Pedro ":   122.30, // ultima virgula é necessária
	}

	funcsESalarios["Rafael Filho"] = 360.20
	/**
	Ao tentar excluir um valor que não
	existe no map, ele apenas não exclui
	e segue a sequencia lógica
	**/
	delete(funcsESalarios, "Inexistente")

	for salario := range funcsESalarios {
		fmt.Println(salario)
	}
}
