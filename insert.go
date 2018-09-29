package main
import (
	"fmt"
	"database/sql"
	//o _ quer dizer que está implicito
	_ "github.com/go-sql-driver/mysql"

	
)


func main(){
	//definindo o banco que vamos trabalhar
	db, err := sql.Open("mysql", "root:0123456@/cursogo")
	if err!= nil{
		//se der erro, cuspir o mesmo
		panic(err)
	}
	//fecha a conexão, similar ao finally
	defer db.Close()


	//monta a query
	stmt, _ :=db.Prepare("insert into usuarios(nome) values(?)")
	//executa a query setando o parâmetro em ? com Maria
	stmt.Exec("Maria")
	stmt.Exec("João")

	/**
	stmt retorna uma resposta contendo o id
	do registro que foi inserido no banco d edados
	**/
	res, _ := stmt.Exec("Pedro")
	/**
	pega o ultimo id inserido
	**/
	id, _ := res.LastInsertId()

	fmt.Println(id)

	linhas, _ := res.RowsAffected()

	fmt.Println(linhas)
}