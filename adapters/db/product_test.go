package db_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/eskokado/go-hexagonal/adapters/db"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func setUp() {
	// Db, err := sql.Open("sqlite3", ":memory:")
	// Db, err := sql.Open("sqlite3", "../../db.sqlite")
	// if err != nil {
	// 	log.Fatal("**** Erro ao abrir - " + err.Error())
	// }
	// createTable(Db)
	// createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
			"id" string,
			"name" string,
			"price" float,
			"status" string
			);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal("####### Erro create table - " + err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc","Product Test",0,"disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	// setUp()
	Db, err := sql.Open("sqlite3", "../../db.sqlite")
	if err != nil {
		log.Fatal("########## Erro ao abrir sqlite3: " + err.Error())
	}
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	if err != nil {
		fmt.Println("*********************ex***************************")
		fmt.Println(product.GetName())
		fmt.Println("************************************************")
	}
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}
