package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"loganlieou.xyz/omelon_backend/utils"
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
	// TODO set in elastic beanstalk ec2 instance
	db, err = sql.Open("mysql", os.Getenv("DB_HOST"))

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/createItem", createItem)
	http.HandleFunc("/createItemWarehouse", createItemWarehouse)
	http.HandleFunc("/getInventory", getInventory)
	http.HandleFunc("/getWarehouse", getWarehouse)
	http.HandleFunc("/updatePrice", updatePrice)
	http.HandleFunc("/updateQuantity", updateQuantity)
	http.HandleFunc("/deleteItem", deleteItem)

	// TODO in prod port is 5000
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// create item
func createItem(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	cost, _ := strconv.ParseFloat(r.FormValue("cost"), 32)
	desc := r.FormValue("desc")
	city := r.FormValue("City")
	quan, _ := strconv.ParseInt(r.FormValue("quan"), 10, 0)

	utils.CreateItem(db, name, cost, desc, city, quan)
}

func createItemWarehouse(w http.ResponseWriter, r *http.Request) {
	quan, _ := strconv.ParseInt(r.FormValue("quan"), 10, 0)
	utils.CreateItemWarehouse(db, r.FormValue("name"), r.FormValue("city"), quan)
}

// ideally we should return a json not a string but this works atm
func getInventory(w http.ResponseWriter, r *http.Request) {
	res := utils.ReadInventory(db)
	fmt.Fprintf(w, res)
}

func getWarehouse(w http.ResponseWriter, r *http.Request) {
	// atm passing as form data which is fine
	var city string
	city = r.FormValue("city")
	res := utils.ReadWarehouse(db, city)
	fmt.Fprintf(w, res)
}

func updatePrice(w http.ResponseWriter, r *http.Request) {
	price, _ := strconv.ParseFloat(r.FormValue("price"), 32)
	utils.UpdatePrice(db, r.FormValue("name"), price)
}

func updateQuantity(w http.ResponseWriter, r *http.Request) {
	quan, _ := strconv.ParseInt(r.FormValue("quan"), 10, 0)
	utils.UpdateQuantity(db, r.FormValue("name"), r.FormValue("city"), quan)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	utils.DeleteItem(db, r.FormValue("name"))
}
