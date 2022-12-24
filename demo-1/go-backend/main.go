package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "backend/models"
	registermodels "backend/registermodels"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
var Age *int

var user_info models.User
var books []*models.Book
var login_info models.LoginUser

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "12345"
	dbname   = "chartdatas"
)

func init() {
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open(user, ConnectionString)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	connection *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	connection, err = wsUpgrader.Upgrade(w, r, nil)
	CheckError(err)

	SendDatas(connection)
}

func SendDatas(conn *websocket.Conn) {
	fmt.Println("SendDatas", Age)
	res, err := db.Query("SELECT * FROM chartdatas WHERE agelimit >= $1", Age)
	CheckError(err)
	defer res.Close()

	for res.Next() {
		book_datas := &models.Book{}
		err = res.Scan(&book_datas.ID, &book_datas.BookName, &book_datas.BookType, &book_datas.Author, &book_datas.Popularity, &book_datas.TotalBook, &book_datas.AgeLimit)
		CheckError(err)
		books = append(books, book_datas)
	}

	books_marshal, _ := json.Marshal(books)
	err = conn.WriteMessage(websocket.TextMessage, []byte(books_marshal))
	CheckError(err)

	defer conn.Close()
}

func RegisterEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	connection, err = wsUpgrader.Upgrade(w, r, nil)
	CheckError(err)

	defer connection.Close()

	for {
		err = connection.ReadJSON(&user_info)
		CheckError(err)
		fmt.Println("Success! Informations sended to Database.")

		// time.Now() - int(user_info.Birthyear)

		fmt.Println(user_info)
		registermodels.Initialize()
		registermodels.GetFromRegister(user_info)
	}
}

func LoginEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	connection, err = wsUpgrader.Upgrade(w, r, nil)
	CheckError(err)

	for {
		err = connection.ReadJSON(&login_info)
		CheckError(err)
		fmt.Println("Success! Informations taken!")
		registermodels.Initialize()
		control_datas, age := registermodels.TakePasswordWithEmail(login_info)
		Age = &age
		err = connection.WriteMessage(websocket.TextMessage, []byte(control_datas))
		CheckError(err)
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dashboard", WsEndpoint)
	router.HandleFunc("/register", RegisterEndpoint)
	router.HandleFunc("/users", LoginEndpoint)
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:9100"}))
	log.Fatal(http.ListenAndServe(":9100", ch(router)))
}
