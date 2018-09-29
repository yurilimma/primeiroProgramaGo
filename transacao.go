package main

import (
	"database/sql"
	"log"

	//o _ quer dizer que está implicito
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:0123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	/**
	abre uma transação
	enquanto a transação tiver aberta,
	caso alguma query der erro, insert,delete,update
	ele desfaz todas as outras que estiverem no bloco de transações

	**/
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")
	/*
	apesar de ser autoincrement não tem problema passar um id
	*/
	stmt.Exec(4000, "Bia")
	stmt.Exec(4001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // chave duplicada
	/**
	tratamento do erro, se der erro, realizar o rollback
	assim ao executar o commit, ele não executará nenhuma
	query das transações
	**/
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
