package main

import "fmt"

func f1() {
	fmt.Println("F1")
}
func f2(p1 string, p2 string) {
	fmt.Printf("F2 %s %s", p1, p2)
}
func f3() string {
	fmt.Printf("F3")
	return "F3"
}

/**se os parâmetros forem do mesmo tipo
	é possível declarar os dois juntos e após
	a declaração o tipo dos mesmos
**/
func f4(p1, p2 string) string {
	//retorna o f4
	return fmt.Sprintf("F4 = %s e %s = ", p1, p2)
}

/**
	funções em go podem retornar mais de um valor
**/
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")
	/**
		Como o retorno da função
		f3 e f4 são strings, as variáveis
		que receberão o retorno das mesmas
		podem ser declaradas assim:
	**/
	r3, r4 := f3(), f4("Param1", "Param2")

	fmt.Println(r3)
	fmt.Println(r4)

	r5, r6 := f5()

	fmt.Println(r5, r6)

	/**
	para ignorar o segundo retorno da função
	não é possível deixar de declarar 2 valores
	já que a função retorna dois valor
	**/
	r7, _ := f5()
	//para ignorar todos os retornos da função
	f5()
	fmt.Println(r7)
}
