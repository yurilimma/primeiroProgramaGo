package main

import (
	"fmt"
)

//criação de um tipo, do tipo float
type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

/**
como não está com receiver (definição antes do método)
não da pra acessar o struct nota, apenas
o que foi passado por parâmetro
*/
func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.0) {
		return "B"
	} else if n.entre(5.0, 6.0) {
		return "D"
	} else if n.entre(3.0, 4.0) {
		return "E"
	} else {
		return "F"
	}
}

/**
como não está com receiver não da pra acessar
o struct nota, apenas o que foi passado por
parâmetro
*/
func notaParaConceitoSwitch(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.0):
		return "B"
	case n.entre(5.0, 6.0):
		return "C"
	case n.entre(3.0, 4.0):
		return "D"
	case n.entre(1.0, 2.0):
		return "E"
	default:
		return "Invalida"
	}
}

func main() {
	var nota nota
	nota = 9.5

	fmt.Println(notaParaConceitoSwitch(nota))

}
