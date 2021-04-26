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

	// Pemanggilan GET Methods pada routes penulis dari database
	e.GET("/author", controllers.FetchAllAuthor)

	return e
}