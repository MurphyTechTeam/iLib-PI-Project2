package middlewares

import (
	"encoding/json"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}
type auth struct{
	Token string `json:"token"`
	Username string `json:"username"`
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

func (c CustomContext) Auth() auth {
	var auth auth
	session, _ := session.Get("session", c)
	auth.Token = session.Values["token"].(string)
	auth.Username = session.Values["username"].(string)
	return auth 
}
