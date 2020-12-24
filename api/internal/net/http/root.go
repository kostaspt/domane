package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func (Handler) RootIndex(ctx echo.Context) error {
	resp := map[string]interface{}{
		"message": fmt.Sprintf("Domane API %s", viper.GetString("VERSION")),
		"datetime": time.Now(),
	}

	return ctx.JSON(http.StatusOK, resp)
}
