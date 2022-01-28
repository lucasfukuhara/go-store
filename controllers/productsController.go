package controllers

import (
	"log"
	"net/http"
	"strconv"
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

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error during price convertion: ", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error during quantity convertion: ", err)
		}

		models.AddNewProduct(name, description, priceConverted, quantityConverted)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
