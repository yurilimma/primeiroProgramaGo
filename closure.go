package main

import "fmt"

func closure() func() {
	x := 10

	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)
	imprimeX := closure()
	/**imprime o 10 porque o o imprimeX
	recebeu o valor da func funcao e
	na mesma o valor é 10
	**/
	imprimeX()
}
