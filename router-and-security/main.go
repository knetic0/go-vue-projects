package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"

	"golang.org/x/crypto/bcrypt"
)

type Info struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn *websocket.Conn
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "12345"
	dbname   = "information"
)

var db *sql.DB

func init() {
	var err error
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", ConnectionString)
	CheckError(err)
}

func Insert(info Info) {
	res, err := db.Exec("INSERT INTO information(email, password) VALUES($1, $2)", info.Email, info.Password)
	CheckError(err)
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Rows affected : %d\n", rowsAffected)
}

func SigninEndpoint(w http.ResponseWriter, r *http.Request) {
	var err error

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade : %s\n", err.Error())
		return
	}

	Signin(wsConn)
}

func EncEndpoint(w http.ResponseWriter, r *http.Request) {
	var err error

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade : %s\n", err.Error())
		return
	}

	GetDatasByEmail(wsConn)
}

func Signin(conn *websocket.Conn) {
	for {
		var info Info

		err := conn.ReadJSON(&info)
		if err != nil {
			fmt.Printf("error reading json : %s\n", err.Error())
			break
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(info.Password), 8)
		CheckError(err)
		info.Password = string(hashedPassword)

		fmt.Println("Informations recevied!")
		Insert(info)
	}
}

func GetDatasByEmail(conn *websocket.Conn) {
	for {
		var info Info
		var getInfo string
		err := conn.ReadJSON(&info)
		if err != nil {
			fmt.Printf("error reading json : %s\n", err.Error())
			break
		}

		err = db.QueryRow("SELECT password FROM information WHERE email=$1", info.Email).Scan(&getInfo)

		switch {
		case err == sql.ErrNoRows:
			log.Printf("No information with that email")
		case err != nil:
			log.Fatal(err)
		default:
			fmt.Printf("Password with encrypted is %s\n", getInfo)
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	var err error

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade : %s\n", err.Error())
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", wsEndpoint)
	router.HandleFunc("/Signin", SigninEndpoint)
	router.HandleFunc("/Decrypt", EncEndpoint)
	log.Fatal(http.ListenAndServe(":9100", router))
}
