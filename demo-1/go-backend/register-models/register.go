package registermodels

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB
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

func init() {
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open(user, ConnectionString)
	CheckError(err)
}

func main() {

}
