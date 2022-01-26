package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templateList = template.Must(template.ParseGlob("templates/*.html")) // load all files with .html from templates folder

func main() {
	http.HandleFunc("/", index)       // will link the address / with the func index
	http.ListenAndServe(":8000", nil) //start the server on port 8000
}

func index(w http.ResponseWriter, r *http.Request) {
	//creating a slice of products to send to index
	products := []Product{
		{Name: "T-shirt", Description: "The Withcer 3", Price: 30.50, Quantity: 10},
		{Name: "T-shirt", Description: "Blue with white lines", Price: 40.50, Quantity: 5},
		{"Tenis", "Confortable", 89.99, 3},
		{"Headset", "Very good", 49.99, 5},
	}

	/* will send the template index as response, the last parameter is to send
	some information to the template.
	The Index is the anotation {{define}} we added on index.html
	*/
	templateList.ExecuteTemplate(w, "Index", products)
}
