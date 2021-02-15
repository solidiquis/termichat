package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	router   *mux.Router
	upgrader websocket.Upgrader
)

func init() {
	router = mux.NewRouter()
	upgrader = websocket.Upgrader{}

	router.HandleFunc("/", logRequest(home))
	router.HandleFunc("/ws", logRequest(serveWebSocket))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}

func serveWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		errLog.Println(err)
		http.Error(
			w,
			"Failed to establish websocket.",
			http.StatusInternalServerError,
		)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		infoLog.Println(msg)
	}

	infoLog.Println("Exiting")
}
