package models

import "github.com/lucasfukuhara/db"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SearchAllProducts() []Product {
	db := db.DbConnect()

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

		err = allProducts.Scan(&id, &name, &desc, &price, &qtt) // will map the values from slice to variables
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

	defer db.Close()
	return products
}

func AddNewProduct(name, desc string, price float64, quantity int) {
	db := db.DbConnect()

	insertNewProduct, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertNewProduct.Exec(name, desc, price, quantity)

	defer db.Close()
}
