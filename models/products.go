package models

import (
	"github.com/lucasfukuhara/db"
)

type Product struct {
	Id          int
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
		p.Id = id
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

func DeleteProduct(id string) {
	db := db.DbConnect()

	deleteQuery, err := db.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err.Error())
	}
	deleteQuery.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.DbConnect()

	queryResult, err := db.Query("select * from products where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for queryResult.Next() {
		var id, qtt int
		var name, desc string
		var price float64

		err = queryResult.Scan(&id, &name, &desc, &price, &qtt)
		if err != nil {
			panic(err.Error())
		}
		productToUpdate.Id = id
		productToUpdate.Description = desc
		productToUpdate.Name = name
		productToUpdate.Price = price
		productToUpdate.Quantity = qtt
	}

	defer db.Close()

	return productToUpdate
}

func UpdateProduct(id int, name, desc string, price float64, quantity int) {
	db := db.DbConnect()

	updateProductQuery, err := db.Prepare("update products set name= $1, description =$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateProductQuery.Exec(name, desc, price, quantity, id)

	defer db.Close()
}
