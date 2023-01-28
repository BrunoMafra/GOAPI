package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func connectDB() *sql.DB {
	conn := "user=postgres dbname=store password=admin host=localhost port=3001 sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quant       int
}

func main() {
	//Handle declara o caminho e a função que deve ser chamada
	http.HandleFunc("/", index)
	http.ListenAndServe(":2000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	conn := connectDB()
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

	temp.ExecuteTemplate(w, "Index", products)
	defer conn.Close()
}
