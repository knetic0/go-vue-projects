package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	registermodels "backend/registermodels"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

type Book struct {
	ID         int    `json:'id'`
	BookName   string `json:'bookname'`
	BookType   string `json:'booktype'`
	Author     string `json:'author'`
	Popularity int    `json:'popularity'`
	TotalBook  int    `json:'totalbook'`
}

var user_info registermodels.User
var books []*Book

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
	res, err := db.Query("SELECT * FROM chartdatas")
	CheckError(err)
	defer res.Close()

	for res.Next() {
		book_datas := &Book{}
		err = res.Scan(&book_datas.ID, &book_datas.BookName, &book_datas.BookType, &book_datas.Author, &book_datas.Popularity, &book_datas.TotalBook)
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
		fmt.Println(user_info)
		registermodels.Initialize()
		registermodels.GetFromRegister(user_info)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dashboard", WsEndpoint)
	router.HandleFunc("/register", RegisterEndpoint)
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:9100"}))
	log.Fatal(http.ListenAndServe(":9100", ch(router)))
}
