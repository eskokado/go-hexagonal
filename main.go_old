package main

import (
	"database/sql"
	"log"

	"github.com/eskokado/go-hexagonal/adapters/db"
	"github.com/eskokado/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	Db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal("########## Erro ao abrir sqlite3: " + err.Error())
	}
	defer Db.Close()
	productDbAdapter := db.NewProductDb(Db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product example", 15.0)

	productService.Enable(product)
}
