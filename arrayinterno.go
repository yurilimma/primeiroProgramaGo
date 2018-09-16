package main

import (
	"fmt"
)

func main() {

	s1 := make([]int, 10, 20)

	s2 := append(s1, 2, 3)
	fmt.Println(s1, s2)
	/**
	Como o array s2 está associado
	ao slice s1, ao alterar algum valor
	do slice s1 é refletido no array s2
	**/
	s1[0] = 7
	fmt.Println(s1, s2)

}
