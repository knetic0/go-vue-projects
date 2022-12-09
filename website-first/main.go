package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

type Info struct {
	Message string `json:"message"`
	Mail    string `json:"mail"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	var err error

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err = wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Printf("Could not upgrade...%s\n", err.Error())
		return
	}

	for {
		var msg Info

		err := wsConn.ReadJSON(&msg)

		if err != nil {
			fmt.Printf("error reading json...%s\n", err.Error())
			break
		}

		fmt.Println("200")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", WsEndpoint)
	log.Fatal(http.ListenAndServe(":9100", router))
}
