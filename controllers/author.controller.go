package controllers

import (
	"mhilmi999/project-2-mhilmi999/models"
	"net/http"
	"strconv"

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

func UpdateAuthor(c echo.Context) error{ // Fungsi untuk Put Method dalam error handlingnya dari tabel author
	// Variable untuk menerima inputan dari user 
	Id_author := c.FormValue("Id_author")
	Nama_author := c.FormValue("Nama_author")


	// Konversi id string to int
	conv_id, err := strconv.Atoi(Id_author)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error()) // Jika error dalam konversi terjadi tampilkan apa errornya
	}

	result, err := models.UpdateAuthor(conv_id, Nama_author) // Parsingkan inputan dari user ke models untuk diteruskan ke database
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // Jika saat parsing terjadi error maka tampilkan errornya
	}

	return c.JSON(http.StatusOK, result) // Jika sukses maka berikan status 200
	
}

func HapusAuthor(c echo.Context) error{
	// Variable untuk menerima inputan dari user
	Id_author := c.FormValue("Id_author")

	// Konversi id string to int
	conv_id, err := strconv.Atoi(Id_author)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // cek saat konversi apakah ada error
	}

	result, err := models.HapusAuthor(conv_id) // parsing data Id_author ke models untuk di eksekusi
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error()) // cek apakah saat parsing terjadi error
	}

	return c.JSON(http.StatusOK, result) // jika semua sukses, maka set status menjadi 200 
}