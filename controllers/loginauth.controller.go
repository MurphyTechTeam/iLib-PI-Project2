package controllers

import (
	"mhilmi999/project-2-mhilmi999/helpers"
	"mhilmi999/project-2-mhilmi999/models"
	"net/http"
	"time"
	"github.com/labstack/echo-contrib/session"
	

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
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

	// return c.String(http.StatusOK, "Sukses Login")

	session,_ := session.Get("session",c)

	// Pembuatan Generate Token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims) // Setup payload 
	claims["username"] = username
	claims["level"] = "Application"
	claims["exp_date"] = time.Now().Add(time.Hour * 72).Unix()

	// Sekarang baru generate token yang ter-encode dan kirim sebagai responsenya
	t, err := token.SignedString([]byte("UIDToken"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	session.Values["token"] = t
	session.Values["username"] = username
	session.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, "/dashboard")
}

func NewRegister(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	_, err := models.RegisMember(username, password)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": "Tolong periksa kembali inputan username atau password Anda",
		})
	}

	return c.Redirect(http.StatusMovedPermanently, "/login")
	
}

