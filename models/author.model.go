package models

import (
	"mhilmi999/project-2-mhilmi999/db"
	"net/http"
)

type Author struct {
	Id_author   int    `json:"id"`
	Nama_author string `json:"nama"`
}

func FetchAllAuthor() (Response, error) {
	var obj Author
	var arrobj []Author
	var res Response
	// inisialisasi variable yang akan digunakan

	con := db.CreateCon() // Inisialisasi awal koneksi model terhadap database

	sqlStatement := "SELECT * FROM author" // query menampilkan seluruh data yang ada pada database iLib pada tabel author

	rows, err := con.Query(sqlStatement) // eksekusi query

	defer rows.Close()                   // tutup koneksi thdp database

	if err != nil {
		return res, err // error akan di handle oleh controller
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_author, &obj.Nama_author) // Scan tiap kolom apakah ada error
		if err != nil {
			return res, err // error akan di handle oleh controller
		}

		arrobj = append(arrobj, obj) // Menampilkan obj yang telah di scan error persatunya dalam bentuk array
	}

	res.Status = http.StatusOK // set status sbg 200 (ok) karena telah dipilah errornya
	res.Message = "Success"    // set pesan berhasil
	res.Data = arrobj          // set data yang ditampilkan pada json dalam bentuk arr yang telah lolos cek error

	return res, nil
}
