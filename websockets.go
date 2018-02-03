package remotecollaber

import (
    "net/http"

    //"github.com/gorilla/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func makeMain() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
    }

    http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	go func(conn *websocket.Conn) {
	    for {
		mType, msg, _ := conn.ReadMessage()
	    }
	}(conn)
	})
}
