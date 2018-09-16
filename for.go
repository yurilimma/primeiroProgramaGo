package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}
	for j := 0; j <= 10; j += 2 {
		fmt.Print("%d ", j)

	}
	//looping infinito
	for {
		fmt.Print("looping")

	}

}
