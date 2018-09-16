package main

import (
	"fmt"
	"time"
)

func notaParaConceito(n float64) string {

	var nota = int(n)

	switch nota {
	case 10:
		//forca entrada no proximo case
		fallthrough
	case 9:
		return "A"
	/**pode colocar mais de um case
	de uma vez**/
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func caseHour() string {
	t := time.Now()
	switch {
	case t.Hour() < 18:
		return "Boa tarde"
	case t.Hour() > 18:
		return "Boa Noite"
	default:
		return "Bom dia"
	}
}

func main() {
	fmt.Printf(notaParaConceito(10))
	fmt.Println(caseHour())
}
