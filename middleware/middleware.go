package middleware

import (
	"github.com/labstack/echo/middleware"
)

var IsAuth = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey : []byte("UIDToken"),
})