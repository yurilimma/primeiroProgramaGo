package matematica

import (
	"fmt"
	"strconv"
)

//Media é responsável
func Media(numero ...float64) float64 {
	total := 0.0

	for _, num := range numero {
		total += num
	}

	media := total / float64(len(numero))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
