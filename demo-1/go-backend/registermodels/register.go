package registermodels

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:'id'`
	Name     string `json:'name'`
	Email    string `json:'email'`
	Password string `json:'password'`
}

var Db *sql.DB
var err error

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "12345"
	dbname   = "register"
)

func CheckError(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func Initialize() {
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Db, err = sql.Open(user, ConnectionString)
	CheckError(err)

}

func GetFromRegister(user_info User) {
	res, err := Db.Exec("INSERT INTO register(name, email, password) VALUES($1, $2, $3)", user_info.Name, user_info.Email, user_info.Password)
	CheckError(err)
	resAffected, _ := res.RowsAffected()
	fmt.Printf("Rows affected -> %d", resAffected)
}
