package models

type Response struct {
	Status int `json:"status"`// akan diisi http 200 kalo sukses, kalo not found 404, dll
	Message string `json:"message"` // Menampung pesan data berhasil atau tidak ditemukan
	Data interface{} `json:"data"`
}

