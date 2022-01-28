package routes

import (
	"net/http"

	"github.com/lucasfukuhara/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index) // will link the address / with the func index
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}
