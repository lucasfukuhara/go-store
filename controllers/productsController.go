package controllers

import (
	"net/http"
	"text/template"

	"github.com/lucasfukuhara/models"
)

var templateList = template.Must(template.ParseGlob("templates/*.html")) // load all files with .html from templates folder

func Index(w http.ResponseWriter, r *http.Request) {

	products := models.SearchAllProducts()

	/* will send the template index as response, the last parameter is to send
	some information to the template.
	The Index is the anotation {{define}} we added on index.html
	*/
	templateList.ExecuteTemplate(w, "Index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {

	/* will send the template index as response, the last parameter is to send
	some information to the template.
	The Index is the anotation {{define}} we added on index.html
	*/
	templateList.ExecuteTemplate(w, "NewProducts", nil)
}
