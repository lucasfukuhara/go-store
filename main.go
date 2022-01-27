package main

import (
	"net/http"

	"github.com/lucasfukuhara/routes"
)

// godoc.org is a repository to check packages we can use

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil) //start the server on port 8000
}
