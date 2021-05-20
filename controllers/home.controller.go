package controllers

import (
	. "mhilmi999/project-2-mhilmi999/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeView(c echo.Context) error {
	flash := c.(*CustomContext).GetFlash()
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"flash": flash,
	})
}

func LoginView(c echo.Context) error {
	flash := c.(*CustomContext).GetFlash()
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"flash": flash,
	})
}

func RegisterView(c echo.Context) error {
	flash := c.(*CustomContext).GetFlash()
	return c.Render(http.StatusOK, "register.html", map[string]interface{}{
		"flash": flash,
	})
}

func DashboardView(c echo.Context) error {
	auth := c.(*CustomContext).Auth()
	return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
		"auth": auth,
	})
}