package controllers

import (
	"mhilmi999/project-2-mhilmi999/models"
	"net/http"

	"github.com/labstack/echo"
)

func FetchAllAuthor(c echo.Context) error {
	// Handle error dan hasilnya
	result, err := models.FetchAllAuthor()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messsage": err.Error()}) // Maka akan ditampilkan pesan error sebagaimana error yang terjadi pada modelsnya
	}

	return c.JSON(http.StatusOK, result)
}