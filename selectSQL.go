package main
import (
	"fmt"
	"database/sql"
	"log"
	//o _ quer dizer que está implicito
	_ "github.com/go-sql-driver/mysql"
)


type usuario struct{
	id int
	nome string
}

func main(){

	db, err := sql.Open("mysql", "root:0123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	rows, _ := db.Query("select id,nome from usuarios where id> ?", 5)
	//é necessário fechar as linhas
	defer rows.Close()
	//percorre o resultado da query
	for rows.Next(){

		var u usuario 
		//mapea o retorno da consulta em &u.id e &u.nome
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
		
	}

}