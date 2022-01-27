package routes

import (
	"net/http"

	"github.com/lucasfukuhara/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index) // will link the address / with the func index
}