package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

// godoc.org is a repository to check packages we can use

func dbConnect() *sql.DB {
	connection := "user=postgres dbname=so_store password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	} else {
		return db
	}
}

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templateList = template.Must(template.ParseGlob("templates/*.html")) // load all files with .html from templates folder

func main() {

	db := dbConnect()
	defer db.Close()

	http.HandleFunc("/", index)       // will link the address / with the func index
	http.ListenAndServe(":8000", nil) //start the server on port 8000
}

func index(w http.ResponseWriter, r *http.Request) {
	//creating a slice of products to send to index
	products := []Product{
		{Name: "T-shirt", Description: "The Withcer 3", Price: 30.50, Quantity: 10},
		{Name: "T-shirt", Description: "Blue with white lines", Price: 40.50, Quantity: 5},
		{"Sneakers", "Confortable", 89.99, 3},
		{"Headset", "Very good", 49.99, 5},
	}

	/* will send the template index as response, the last parameter is to send
	some information to the template.
	The Index is the anotation {{define}} we added on index.html
	*/
	templateList.ExecuteTemplate(w, "Index", products)
}
