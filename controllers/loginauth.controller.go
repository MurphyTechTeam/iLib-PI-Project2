package controllers

import (
	"mhilmi999/project-2-mhilmi999/helpers"
	"mhilmi999/project-2-mhilmi999/models"
	"net/http"

	"github.com/labstack/echo"
)

func GenerateHashPassword(c echo.Context) error { // Fungsi untuk handling controller dalam peng-hashan password dari inputan user
	password := c.FormValue("password") // tangkap inputan password user (plain)

	hash, _ := helpers.PasswordhHash(password) // parsing password plain ke fungsi ini agar dapat di enkripsi secara decrypt

	return c.JSON(http.StatusOK, hash) // set status 200
}

func CheckingLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckingLogin(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	return c.String(http.StatusOK, "Sukses Login")
}
