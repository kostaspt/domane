package v1

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (Handler) RootIndex(ctx echo.Context) error {
	resp := map[string]interface{}{
		"message": "Domane API v1.0",
		"datetime": time.Now(),
	}

	return ctx.JSON(http.StatusOK, resp)
}
