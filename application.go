package main

import (
	//"os"
	//"log"
	"fmt"
	"database/sql"
	"net/http"

	//"loganlieou.xyz/omelon_backend/utils"
	_ "github.com/go-sql-driver/mysql"
)

/*
Enviornment Variables
DB_HOST
is of the form
user:pass@tcp(ip:port)/dbname
*/

var db *sql.DB
var err error

func main() {

	/*
	I'm so sad rn
	db, err = sql.Open("mysql", os.Getenv("DB_HOST"))

	if err != nil {
		log.Fatal(err)
	}
	*/
	

	// HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test")
	})
	http.ListenAndServe(":5000", nil)
}

