package main

import (
	r "GOAPI/routes"
	"net/http"
)

func main() {
	//Handle declara o caminho e a função que deve ser chamada
	r.LoadRoutes()
	http.ListenAndServe(":2000", nil)
}
