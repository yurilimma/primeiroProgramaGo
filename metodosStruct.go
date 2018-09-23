package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

/**
retorna o p.nome + p.sobrenome
Obs.: Como não estamos trabalhando com ponteiro
qualquer alteração realizada em p nesse bloco
não será refletida na main ou na função original
que chamou a mesma.
**/
func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

/**
trabalha em cima de um ponteiro de pessoa
e recebe uma string com o nome completo por
parâmetro
**/
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}
func main() {
	p1 := pessoa{"pedro", "silva"}
	fmt.Println(p1.getNomeCompleto())
	/**
	altera o valor do struct por referencia
	via ponteiro
	**/
	p1.setNomeCompleto("Maria Pereira")
	fmt.Println(p1.getNomeCompleto())

}
