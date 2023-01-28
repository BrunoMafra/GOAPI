package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

type Product struct {
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
	products := []Product{
		{"Xbox One", "Good for gaming", 349, 5},
		{"PS5", "Good for exclusives", 400, 3},
		{"Nintendo", "Mario", 200, 7},
		{"Gaming PC", "GTX 4090", 1099, 6},
	}
	temp.ExecuteTemplate(w, "Index", products)
}
