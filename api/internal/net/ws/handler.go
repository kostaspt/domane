package ws

import (
	"fmt"
	"net/http"

	"github.com/google/wire"
	"github.com/gorilla/websocket"
	"github.com/kostaspt/domane/api/config"
)

var ProviderSet = wire.NewSet(NewHandler)

type Handler struct {
	config *config.Config
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: false,
}

func NewHandler(c *config.Config) *Handler {
	return &Handler{
		config: c,
	}
}

func (h Handler) writeWelcomeMessage(ws *websocket.Conn) error {
	return ws.WriteJSON(map[string]interface{}{
		"message": fmt.Sprintf("Domane WS %s", h.config.Version),
		"success": true,
	})
}
