package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	ID                 int
	Title, Description string
	Price              float32
}

const (
	host     = "localhost"
	port     = "3000"
	user     = "postgres"
	password = "<password>"
	dbname   = "products"
)

var db *sql.DB

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connString)
	CheckError(err)
}

func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO products(title, description, price) VALUES($1, $2, $3)", data.Title, data.Description, data.Price)
	CheckError(err)
	rowsAffected, err := result.RowsAffected()
	CheckError(err)
	fmt.Printf("%d Kayit Etkilendi.", rowsAffected)
}

func UpdateProduct(data Product) {
	result, err := db.Exec("UPDATE products SET title=%2 description=$3 price=$4 WHERE id=$1", data.ID, data.Title, data.Description, data.Price)
	CheckError(err)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("%d Kayit Etkilendi.", rowsAffected)
}

func GetProducts() {
	rows, err := db.Query("SELECT * FROM products")
	CheckError(err)
	if err != sql.ErrNoRows {
		fmt.Println("No Records Found!")
		return
	}
	defer rows.Close()
	var products []*Product
	for rows.Next() {
		prd := &Product{}
		err := rows.Scan(&prd.ID, &prd.Title, &prd.Description, &prd.Price)
		CheckError(err)
		products = append(products, prd)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, pr := range products {
		fmt.Printf("%d - %s, %s, %.2f\n", pr.ID, pr.Title, pr.Description, pr.Price)
	}
}

func GetProductsByID(id int) {
	var product string
	err := db.QueryRow("SELECT title FROM products WHERE id=$1", id).Scan(&product)
	CheckError(err)
	if err != sql.ErrNoRows {
		fmt.Println("No Records Found!")
		return
	}
	fmt.Printf("Product is %s\n", product)
}
