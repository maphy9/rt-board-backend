package http

import (
	"log"
	"net/http"

	"github.com/maphy9/rt-board-backend/internal/http/websocket"
)

func SetupHttp() {
	websocketManager := websocket.NewManager()

	http.HandleFunc("/ws", websocketManager.ServeWS)
	log.Println("Websockets listening on :8080/ws")
	http.ListenAndServe(":8080", nil)
}