package main
import ("fmt"
        "database/sql"
      _ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	db, err := sql.Open("mysql", "root:h._qf8B[%N>UyhEXqr,Y>BRT@tcp(localhost:3306)/recettes")

	println("premiere erreur potentielle")
	checkError(err)

	println("deuxieme erreur potentielle")
	err= db.Ping()
	checkError(err)

	// Le code ne fonctionne pas bien à partir de ce point
	// Deux solutions possibles:
	// - soit vérifier si la table existe deja (avec SHOW TABLES)
	// - soit en faisant un try catch de l'erreur renvoyée par l'executable sql
	// Create table.
  _, err = db.Exec("CREATE TABLE IF NOT EXISTS ingredients (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
  checkError(err)
  fmt.Println("Creation de table reussi.")

	_, err = db.Exec("SHOW TABLES LIKE ingredients ;")

	// Insert some data into table.
  	sqlStatement, err := db.Prepare("INSERT INTO ingredients (name, quantity) VALUES (?, ?);")
  	res, err := sqlStatement.Exec("banana", 150)
    res2, err := sqlStatement.Exec(200, 200)
    println(res2)
    checkError(err)
  	rowCount, err := res.RowsAffected()
  	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

}
