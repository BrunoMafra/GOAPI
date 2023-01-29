package routes

import (
	c "GOAPI/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/new", c.New)
	http.HandleFunc("/insert", c.Insert)
	http.HandleFunc("/delete", c.Delete)
}
