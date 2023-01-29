package controllers

import (
	"GOAPI/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SelectAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		quant, _ := strconv.Atoi(r.FormValue("quant"))
		models.InsertProduct(name, description, price, quant)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct, _ := strconv.Atoi(r.URL.Query().Get("id"))
	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}
