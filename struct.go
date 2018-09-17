package main

import (
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// método: função com receiver
//similar ao passar por parâmetro
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}
func main() {
	//receptor
	var produto1 produto
	/*ao instanciar struts é necessário separar
	os atributos com virgula */
	produto1 = produto{
		nome:     "lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	var produto2 produto
	produto2 = produto{nome: "caneta", preco: 2.79, desconto: 0.05}
	/** para chamar o método precoComDesconto()
		é necessário colocar o tipo que o referencia
		o que está evidênciado antes do parenteses
		no caso um struct do tipo produto
	**/
	fmt.Println(produto1)
	fmt.Println(produto1.precoComDesconto())
	fmt.Println(produto2)
	fmt.Println(produto2.precoComDesconto())

}
