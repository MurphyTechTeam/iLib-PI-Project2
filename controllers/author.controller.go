package controllers

import (
	"mhilmi999/project-2-mhilmi999/models"
	"net/http"

	"github.com/labstack/echo"
)

func FetchAllAuthor(c echo.Context) error { // Fungsi untuk Get Method dalam error handling nya dari tabel author
	// Handle error dan hasilnya
	result, err := models.FetchAllAuthor()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messsage": err.Error()}) // Maka akan ditampilkan pesan error sebagaimana error yang terjadi pada modelsnya
	}

	return c.JSON(http.StatusOK, result)
}

func StoreAuthor(c echo.Context) error{ // Fungsi untuk Post Method dalam error handlingnya dari tabel author
	// Tampung parameter yang diinputkan aplikasi eksternal
	Nama_author := c.FormValue("Nama_author")

	result, err := models.StoreAuthor(Nama_author)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK,result)
}