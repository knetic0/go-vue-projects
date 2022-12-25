package loginmodels

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"

	models "backend/models"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "12345"
	dbname   = "login"
)

var db *sql.DB
var err error
var Email string

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Initialize() {
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open(user, ConnectionString)
	CheckError(err)
}

func Insert(login_information models.LoginUser, age int) {
	fmt.Println("datalar = ", login_information)
	res, err := db.Exec("INSERT INTO login(email, password, age) VALUES($1, $2, $3)", login_information.Email, login_information.Password, age)
	CheckError(err)
	Email = login_information.Email
	fmt.Println(Email)
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Rows Affected -> %d", rowsAffected)
}

func GetAge() int {
	var Age int

	fmt.Println("kontrol noktasi")
	err = db.QueryRow("SELECT age FROM login").Scan(&Age)
	CheckError(err)
	fmt.Println(Age)

	return Age
}

func Signout() {
	_, err := db.Exec("DELETE FROM login")
	CheckError(err)
	fmt.Println("Success! Signout.")
}
