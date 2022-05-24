package utils

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateItem(db *sql.DB, name string, cost float32, desc string) {
	var query string = fmt.Sprintf("INSERT %s INTO Inventory VALUES(%s, %f, %s)", name, name, cost, desc)

	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateQuantity(db *sql.DB, item_name string, city string, new_quan int) {
	var query string = fmt.Sprintf("UPDATE %s SET quantity=%d WHERE name=%s", city, new_quan, item_name)

	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatePrice(db *sql.DB, item_name string, new_price float32) {
	var query string = fmt.Sprintf("UPDATE Inventory SET price=%f", new_price)

	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadWarehouse(db *sql.DB, city string) {
	var query string = "SELECT * FROM " + city

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var item City
		rows.Scan(&item.name, &item.quan, &item.city)
		item.print()
	}
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
		res += " " + item.name
	}

	return res
}

func DeleteItem(db *sql.DB, item_name string) {
	var query string = "DELETE _ _"

	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted item: ", item_name)
}
