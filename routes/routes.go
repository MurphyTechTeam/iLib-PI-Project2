package routes

import (
	"mhilmi999/project-2-mhilmi999/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo{
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ini percobaan Echo nya dah jalan belum")
	}) // Mendefinisikan routing awal akses nya

	// Pemanggilan GET Methods pada routes author dari database
	e.GET("/author", controllers.FetchAllAuthor)

	// Pemanggilan Post Methods pada routes author dari database
	e.POST("/author", controllers.StoreAuthor)
	
	// Pemanggilan Put Methods pada routes author untuk ke database
	e.PUT("/author", controllers.UpdateAuthor)

	// Pemanggilan Delete Methods pada routes author untuk ke database
	e.DELETE("/author", controllers.HapusAuthor)

	// Melakukan validasi user secara sederhana (login) dengan metode bcrypt
	e.GET("/make-hash/:password", controllers.GenerateHashPassword)

	// Melakukan authorization user
	e.POST("/login", controllers.CheckingLogin)

	return e
}