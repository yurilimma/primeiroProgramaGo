package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct para json
	p1 := produto{1, "notebook", 1899.90, []string{"promoção", "eletrônico"}}
	//retorna um slice de bytes e um erro
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 produto
	jsonString := `{"id":2, "nome": "Caneta", "preco": 8.90, "tags": ["Papelaria","Promoção"]}`
	/**converte para bytes o json e informa o tipo de struct
	que o mesmo pertence*/
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
