package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	_ "github.com/lib/pq"
)

type Message struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn *websocket.Conn
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "<password>"
	dbname   = "todos"
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

func Insert(value Message) {
	res, err := db.Exec("INSERT INTO todos(title, done) VALUES($1, $2)", value.Title, value.Done)
	CheckError(err)
	resCounter, _ := res.RowsAffected()
	fmt.Printf("Rows Affected: %d", resCounter)
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		// check the http.Request
		// make sure it's OK to access
		return true
	}
	var err error
	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade: %s\n", err.Error())
		return
	}

	defer wsConn.Close()

	// event loop
	for {
		var msg Message

		err := wsConn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error reading JSON: %s\n", err.Error())
			break
		}

		fmt.Println("Message Received")
		Insert(msg)
		SendMessage("Hello, Client!")
	}
}

func SendMessage(msg string) {
	err := wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Printf("error sending message: %s\n", err.Error())
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/socket", WsEndpoint)

	log.Fatal(http.ListenAndServe(":9100", router))

}
