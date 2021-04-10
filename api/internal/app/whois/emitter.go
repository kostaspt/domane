package whois

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

type Emitter struct {
	ws *websocket.Conn
	mu sync.RWMutex
}

func NewEmitter(ws *websocket.Conn) *Emitter {
	return &Emitter{
		ws: ws,
	}
}

func (e *Emitter) Start() {
	for {
		_, domain, err := e.ws.ReadMessage()
		if err != nil {
			log.Err(err).Send()
			continue
		}

		d := string(domain)

		go func() {
			err = e.checkDomain(d)
			if err != nil {
				e.handleError(d, err)
			}
		}()
	}
}

func (e *Emitter) checkDomain(domain string) error {
	info, err := Get(domain)
	if err != nil {
		return err
	}

	return e.write(domain, info)
}

func (e *Emitter) write(domain string, info Info) error {
	// Avoid concurrent writes to websocket connection
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.ws.WriteJSON(Message{
		Domain: domain,
		Info:   info,
	})
}

func (e *Emitter) handleError(domain string, err error) {
	log.Err(err).Send()
	_ = e.write(domain, Info{})
	return
}
