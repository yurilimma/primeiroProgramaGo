package main

import (
	"fmt"
	"reflect"
)

func main() {
	//array
	a1 := [3]int{1, 2, 3}
	/**slice (array que não tem um tamanho
	previamente definido. Basicamente é um trecho/pedaço de um array)
	**/
	s1 := []int{1, 2, 3}

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:2]
	fmt.Println(a2, s2)

	/**slice que pega a partir da posição 0
	até um elemento antes da posição 2 do array a2
	**/
	s3 := a2[:2]
	fmt.Println(a2, s3)
	/**slice que pega a partir da posição 0
	até um elemento antes da posição 1 do slice s2
	**/
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
