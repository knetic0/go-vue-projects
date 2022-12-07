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
	ID    int    `json:"id"`
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
	password = "12345"
	dbname   = "todos"
)

var db *sql.DB
var todos []*Message

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
	GetDatas()
}

func Insert(value Message) {
	res, err := db.Exec("INSERT INTO todos(title, done) VALUES($1, $2)", value.Title, value.Done)
	CheckError(err)
	resCounter, _ := res.RowsAffected()
	fmt.Printf("Rows Affected: %d", resCounter)
}

func GetDatas() {
	res, err := db.Query("SELECT * FROM todos")
	CheckError(err)
	defer res.Close()

	for res.Next() {
		tds := &Message{}
		err := res.Scan(&tds.ID, &tds.Title, &tds.Done)
		CheckError(err)
		todos = append(todos, tds)
	}
	if err = res.Err(); err != nil {
		log.Fatal(err)
	}
}

/*
func DeleteVal() {
	var delmsg Message
	err := wsConn.ReadJSON(&delmsg)
	CheckError(err)
	del, err := db.Exec("DELETE FROM todos WHERE id=$1", delmsg.ID)
	CheckError(err)
	delAffected, err := del.RowsAffected()
	CheckError(err)
	fmt.Println(delAffected)
}
*/

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	var err error

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		// check the http.Request
		// make sure it's OK to access
		return true
	}

	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade: %s\n", err.Error())
		return
	}

	for _, key := range todos {
		err = wsConn.WriteMessage(websocket.TextMessage, []byte(key.Title))
		CheckError(err)
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

		if msg.Title != "" && msg.Title != " " {
			Insert(msg)
		}
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
