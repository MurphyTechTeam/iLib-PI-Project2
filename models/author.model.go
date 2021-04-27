package models

import (
	"mhilmi999/project-2-mhilmi999/db"
	"net/http"
)

type Author struct {
	Id_author   int    `json:"id"`
	Nama_author string `json:"nama"`
}

func FetchAllAuthor() (Response, error) { // Fungsi untuk mendapatkan Get Methods dari tabel author
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

func StoreAuthor(Nama_author string) (Response, error){ // Fungsi untuk meng-inputkan data Post Method dari tabel author
	var res Response // variable untuk menampung response nya

	con := db.CreateCon() // inisialisasi koneksi database

	sqlStatement := "INSERT INTO author (nama_author) VALUES (?)" // query insert data
 
	stmt, err := con.Prepare(sqlStatement) // Variable persiapan awal sebelum eksekusi insert data

	if err != nil {
		return res,err // lakukan cek apakah ada error atau tidak
	}

	result, err := stmt.Exec(Nama_author) // Eksekusi insert data query nya

	if err != nil{
		return res, err // cek lagi apakah saat insert data terdapat error
	}

	lastInsertedId, err := result.LastInsertId()

	if err !=  nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"Id yang terakhir dimasukan adalah" : lastInsertedId,
	}

	return res, nil

}