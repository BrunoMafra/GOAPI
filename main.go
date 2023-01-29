package main

import (
	"GOAPI/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	//Handle declara o caminho e a função que deve ser chamada
	http.HandleFunc("/", index)
	http.ListenAndServe(":2000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SelectAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}
