package controllers

import (
	"mhilmi999/echo-rest/helpers"
	"net/http"

	"github.com/labstack/echo"
)

func GenerateHashPassword(c echo.Context) error{ // Fungsi untuk handling controller dalam peng-hashan password dari inputan user
	password := c.FormValue("password") // tangkap inputan password user (plain)

	hash, _ := helpers.HashPassword(password) // parsing password plain ke fungsi ini agar dapat di enkripsi secara decrypt

	return c.JSON(http.StatusOK, hash) // set status 200
}