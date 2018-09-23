package main

import "fmt"

/**
	informa que pra alguma
	struct ser imprimivel,
	deve conter o método toString()
**/
type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

/**
implementação do método toString()
para cada interface
**/
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f ", p.nome, p.preco)
}

/**
como quem chama o toString é o imprimir
o mesmo faz a chamada x.toString
chama o toString para o tipo de struct
que o está solicitando, pessoa ou produto
**/
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}
func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())

	imprimir(coisa)
	//similar ao polimorfismo
	coisa = produto{"Calça Jeans", 79.90}
	imprimir(coisa)
}
