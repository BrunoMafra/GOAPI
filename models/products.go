package models

import "github.com/BrunoMafra/GOAPI/db"

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
