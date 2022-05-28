package utils

import (
	"database/sql"
	"fmt"
	"log"
)

// for now only one city
// TODO add so that you can init to many cities
func CreateItem(db *sql.DB, name string, cost float64, desc string, city string, quan int64) {
	// insert into Inventory
	var query string = fmt.Sprintf("INSERT INTO Inventory (name, price, description) VALUES(%s, %f, %s)", name, cost, desc)
	// insert into inital City
	CreateItemWarehouse(db, city, name, quan)
	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateItemWarehouse(db *sql.DB, city string, name string, quan int64) {
	var query string = fmt.Sprintf("INSERT INTO Warehouse (name, quantity, city) VALUES(%s, %d, %s)", name, quan, city)
	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateQuantity(db *sql.DB, item_name string, city string, new_quan int64) {
	var query string = fmt.Sprintf("UPDATE %s SET quantity=%d WHERE name=%s", city, new_quan, item_name)
	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatePrice(db *sql.DB, item_name string, new_price float64) {
	var query string = fmt.Sprintf("UPDATE Inventory SET price=%f WHERE name=%s", new_price, item_name)
	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadWarehouse(db *sql.DB, city string) string {
	var query string = fmt.Sprintf("SELECT * FROM %s", city)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var res string = ""

	for rows.Next() {
		var item City
		rows.Scan(&item.name, &item.quan, &item.city)
		res = fmt.Sprintf("%s %s %d %s,", res, item.name, item.quan, item.city)
	}
	return res
}

func ReadInventory(db *sql.DB) string {
	var query string = "SELECT * FROM Inventory"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var res string = ""

	for rows.Next() {
		var item Inventory
		rows.Scan(&item.name, &item.cost, &item.desc)
		res = fmt.Sprintf("%s %s %f %s,", res, item.name, item.cost, item.desc)
	}

	return res
}

func DeleteItem(db *sql.DB, item_name string) {
	// TODO figure out this query
	// TODO does this query work?
	var query string = "DELETE FROM Inventory WHERE name=" + item_name

	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}
