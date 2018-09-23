package main

import (
	"fmt"
)

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	//um struct dentro do outro
	itens []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	//percorre os itens do pedido
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}
func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 3, 22.10},
			item{3, 4, 22.10},
		},
	}

	fmt.Printf("O valor total do pedido é: %.2f ", pedido.valorTotal())
}
