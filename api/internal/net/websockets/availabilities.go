package websockets

import (
	"sync"

	"github.com/labstack/echo/v4"

	"github.com/kostaspt/domane/api/internal/app/whois"
)

type AvailabilityRequest struct {
	Domains []string `json:"domains" valid:"required"`
}

type Availability struct {
	Domain      string `json:"domain"`
	IsAvailable bool   `json:"is_available"`
}

func (Handler) Availabilities(ctx echo.Context) error {
	p := whois.NewWhois(whois.NewParser())

	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}

	err = writeWelcomeMessage(ws)
	if err != nil {
		return err
	}

	var mu sync.RWMutex

	for {
		_, domain, err := ws.ReadMessage()
		if err != nil {
			ctx.Logger().Error(err)
			continue
		}

		d := string(domain)

		go func(d string) {
			isAvailable, err := p.IsAvailable(d)
			if err != nil {
				ctx.Logger().Error(err)
				return
			}

			// Avoid concurrent writes to websocket connection
			mu.Lock()
			defer mu.Unlock()

			err = ws.WriteJSON(Availability{
				Domain:      d,
				IsAvailable: isAvailable,
			})
			if err != nil {
				ctx.Logger().Error(err)
			}
		}(d)
	}
}
