package main
import (
	"database/sql"
	"log"

	//o _ quer dizer que está implicito
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:0123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//ao fechar ele realiza o commit automaticamente
	stmt, _ := db.Prepare("update usuarios set nome= ? where id = ?")
	stmt.Exec("uuuu", 1)
	stmt.Exec("vvvv", 2)
	
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(1)




}