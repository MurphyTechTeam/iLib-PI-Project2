package routes

import (
	"mhilmi999/project-2-mhilmi999/controllers"
	"mhilmi999/project-2-mhilmi999/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo{
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ini percobaan Echo nya dah jalan belum")
	}) // Mendefinisikan routing awal akses nya

	// Coba template (tampilan html) pada echo framework go
	e.Renderer = controllers.NewRenderer("./view/*.html", true)
	e.GET("/index", func(c echo.Context) error {
		data := controllers.M{"message" : "Halo Halo Bandung!"}
		return c.Render(http.StatusOK, "index.html",data)
	})

	// Pemanggilan GET Methods pada routes author dari database
	e.GET("/author", controllers.FetchAllAuthor, middleware.IsAuth)

	// Pemanggilan Post Methods pada routes author dari database
	e.POST("/author", controllers.StoreAuthor, middleware.IsAuth)
	
	// Pemanggilan Put Methods pada routes author untuk ke database
	e.PUT("/author", controllers.UpdateAuthor, middleware.IsAuth)

	// Pemanggilan Delete Methods pada routes author untuk ke database
	e.DELETE("/author", controllers.HapusAuthor, middleware.IsAuth)

	// Melakukan validasi user secara sederhana (login) dengan metode bcrypt
	e.GET("/make-hash/:password", controllers.GenerateHashPassword)

	// Melakukan authorization user
	e.POST("/login", controllers.CheckingLogin)

	return e
}