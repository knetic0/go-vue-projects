package registermodels

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	loginmodels "backend/loginmodels"
	models "backend/models"

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

func Initialize() {
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open(user, ConnectionString)
	CheckError(err)

}

func GetFromRegister(user_info models.User) {
	res, err := db.Exec("INSERT INTO register(name, email, password, birthyear) VALUES($1, $2, $3, $4)", user_info.Name, user_info.Email, user_info.Password, user_info.Birthyear)
	CheckError(err)
	resAffected, _ := res.RowsAffected()
	fmt.Printf("Rows affected -> %d", resAffected)
}

func TakePasswordWithEmail(login_info models.LoginUser) {
	var pass string
	var year string

	err = db.QueryRow("SELECT password, birthyear FROM register WHERE email=$1", string(login_info.Email)).Scan(&pass, &year)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No users with that Email")

	case err != nil:
		log.Fatal(err)

	default:
		fmt.Println("Success!")
	}

	if login_info.Password == pass {

		fmt.Println("Giris basarili.")

		switch {
		case err == sql.ErrNoRows:
			log.Printf("No users with that Email")

		case err != nil:
			log.Fatal(err)

		default:
			fmt.Println("Success!")
		}

		integer_year, _ := strconv.Atoi(year)
		age := time.Now().Year() - integer_year
		fmt.Println(age)

		loginmodels.Initialize()
		loginmodels.Insert(login_info, age)
		loginmodels.GetAge()

	} else {
		fmt.Println("Sifreniz yanlistir.")
	}
}
