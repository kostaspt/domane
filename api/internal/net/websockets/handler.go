package websockets

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

type Handler struct{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: false,
}

func New() *Handler {
	return &Handler{}
}

func writeWelcomeMessage(ws *websocket.Conn) error {
	return ws.WriteJSON(map[string]interface{}{
		"message": fmt.Sprintf("Domane WS %s", viper.GetString("VERSION")),
		"success": true,
	})
}
