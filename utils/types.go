package utils

import "fmt"

type Inventory struct {
	name string
	cost float64
	desc string
}

type City struct {
	name string
	quan int64
	city string
}

func (item Inventory) print() {
	fmt.Printf("name:%s\nprice:%f\ndescription:%s\n",
		item.name, item.cost, item.desc)
}

func (item City) print() {
	fmt.Printf("name:%s\nquantity:%d\ncity:%s\n",
		item.name, item.quan, item.city)
}
