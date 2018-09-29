package main

import (
	"database/sql"
	//o _ quer dizer que está implicito
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*
	abre conexão com o mysql
	usuario:senha@/nomeBanco
	como não queremos entrar em banco nenhum e sim criar um, não colocamos nada
	*/
	db, err := sql.Open("mysql", "root:0123456@/")
	if err != nil {
		panic(err)
	}
	//fecha a conexão quando acaba o método main
	defer db.Close()
	//cria branco cursogo
	exec(db, "create database if not exists cursogo")

	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	//tem que ser o `´ se não não reconhece
	exec(db, `create table usuarios(
				id integer auto_increment,
				nome varchar(80),
				PRIMARY KEY (id)
			)`)
}

/**
recebe a conexão com o banco, e a string a ser executada.
retorna um objeto do tipo sql.Result em result
**/
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
	return result
}
