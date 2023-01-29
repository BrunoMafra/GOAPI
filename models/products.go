package models

import "GOAPI/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quant       int
}

func SelectAllProducts() []Product {
	conn := db.ConnectDB()
	selectProducts, err := conn.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var Id, Quant int
		var Name, Description string
		var Price float64

		err := selectProducts.Scan(&Id, &Name, &Description, &Price, &Quant)

		if err != nil {
			panic(err.Error())
		}

		p.Id = Id
		p.Name = Name
		p.Description = Description
		p.Price = Price
		p.Quant = Quant

		products = append(products, p)
	}
	defer conn.Close()
	return products
}

func InsertProduct(name string, description string, price float64, quant int) {
	conn := db.ConnectDB()
	defer conn.Close()

	insertProductDB, err := conn.Prepare("insert into products (name, description, price, quant) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertProductDB.Exec(name, description, price, quant)
}

func DeleteProduct(id int) {
	conn := db.ConnectDB()
	defer conn.Close()

	deleteProductDB, err := conn.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteProductDB.Exec(id)
}
