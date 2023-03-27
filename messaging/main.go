package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Alfeenn/messaging/handler"
	"github.com/Alfeenn/messaging/model"
	"github.com/gorilla/websocket"
)

type M map[string]interface{}

func main() {
	connections := model.NewSocketConnections().ConnectionSocket()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("web/index.html")
		if err != nil {
			http.Error(w, "Could not open requested file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		currentGorillaConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}

		username := r.URL.Query().Get("username")
		currentConn := model.WebSocketConnection{Conn: currentGorillaConn, Username: username}
		connections = append(connections, &currentConn)

		go handler.HandleIO(&currentConn, connections)

	})

	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}
