package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

// godoc.org is a repository to check packages we can use

func dbConnect() *sql.DB {
	connection := "user=postgres dbname=go_store password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	} else {
		return db
	}
}

type Product struct {
	ID          int
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
	db := dbConnect()

	allProducts, err := db.Query("select * from products order by name asc")
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, qtt int
		var name, desc string
		var price float64

		err = allProducts.Scan(&id, &name, &desc, &price, &qtt)
		if err != nil {
			panic(err.Error())
		}
		p.ID = id
		p.Name = name
		p.Description = desc
		p.Price = price
		p.Quantity = qtt

		products = append(products, p)
	}

	/* will send the template index as response, the last parameter is to send
	some information to the template.
	The Index is the anotation {{define}} we added on index.html
	*/
	templateList.ExecuteTemplate(w, "Index", products)

	defer db.Close()
}
