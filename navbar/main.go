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

var db *sql.DB
var msg []*Navbar

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Navbar struct {
	ID    int
	Title string
}

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "12345"
	dbname   = "navbar"
)

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn *websocket.Conn
)

func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connString)
	CheckError(err)
	defer db.Close()
	err = db.Ping()
	CheckError(err)
	GetDatas()
}

func GetDatas() {
	res, err := db.Query("SELECT * FROM navbar")
	CheckError(err)
	defer res.Close()

	for res.Next() {
		nvb := &Navbar{}
		err := res.Scan(&nvb.ID, &nvb.Title)
		CheckError(err)
		msg = append(msg, nvb)
	}
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	var err error

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	for _, key := range msg {
		err = wsConn.WriteMessage(websocket.TextMessage, []byte(key.Title))
		CheckError(err)
	}

	defer wsConn.Close()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/socket", WsEndpoint)
	log.Fatal(http.ListenAndServe(":9100", router))
}
