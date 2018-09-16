package main

import (
	"fmt"
)

func tipo(i interface{}) string {

	switch i.(type) {
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case int:
		return "Inteiro"
	case func():
		return "Função"
	default:
		return "Tipo não definido"
	}
}

func main() {

	fmt.Println(tipo(2.1))
	fmt.Println(tipo(2))
	fmt.Println(tipo("Epa"))
	fmt.Println(tipo(func() {}))
}
