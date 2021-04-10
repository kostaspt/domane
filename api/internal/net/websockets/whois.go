package websockets

import (
	"github.com/labstack/echo/v4"

	"github.com/kostaspt/domane/api/internal/app/whois"
)

func (Handler) Whois(ctx echo.Context) error {
	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}

	err = writeWelcomeMessage(ws)
	if err != nil {
		return err
	}

	whois.NewEmitter(ws).Start()
	return nil
}
