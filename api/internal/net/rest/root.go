package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h Handler) RootIndex(ctx echo.Context) error {
	resp := map[string]interface{}{
		"message":  fmt.Sprintf("Domane API %s", h.config.Version),
		"datetime": time.Now(),
	}

	return ctx.JSON(http.StatusOK, resp)
}
