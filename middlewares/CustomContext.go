package middlewares

import (
	"encoding/json"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}

func (c CustomContext) GetFlash() map[string]interface{} {
	session, _ := session.Get("flash", c)
	if flash := session.Flashes(); len(flash) > 0 {
		flashes := make(map[string]interface{})
		json.Unmarshal([]byte(flash[0].(string)), &flashes)
		return flashes
	}
	return map[string]interface{}{}
}
