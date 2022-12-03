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

// PostgreSQL Init Fonksiyonu, Baglanti Saglaniyor
func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
}

// Database Insert Islemleri
func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO products(title, description, price) VALUES($1, $2, $3)", data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Etkilenen kayit sayisi = %d", rowsAffected)
}

// Update Islemleri
func Update(data Product) {
	result, err := db.Exec("UPDATE products SET title =$2 description=$3 price=$4 WHERE id=$1", data.ID, data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Etkilenen kayit sayisi = %d", rowsAffected)
}

// Butun Degerleri Aliyoruz.
func GetProducts() {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No records found!")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		prd := &Product{}
		err := rows.Scan(&prd.ID, &prd.Title, &prd.Description, &prd.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, prd)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, pr := range products {
		fmt.Printf("%d - %s, %s, %.2f\n", pr.ID, pr.Title, pr.Description, pr.Price)
	}
}

// Belirli Bir ID'e Gore Veri Alma
func GetProductsByID(id int) {
	var product string
	err := db.QueryRow("SELECT title FROM products WHERE id=$1", id).Scan(&product)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No product with that ID.")

	case err != nil:
		log.Fatal(err)

	default:
		fmt.Printf("Product is %s\n", product)
	}
}

// ID'e Gore Veri Silme
func DeleteProductsByID(id int) {
	_, err := db.Exec("DELETE FROM products WHERE id=$1;", id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No product with that ID.")
		}
		log.Fatal(err)
	}
	fmt.Printf("Deleted ID=%d", id)
}
