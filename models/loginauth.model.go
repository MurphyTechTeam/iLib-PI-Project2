package models

import (
	"database/sql"
	"fmt"
	"mhilmi999/project-2-mhilmi999/db"
	"mhilmi999/project-2-mhilmi999/helpers"
)

type Member struct {
	Id int `json:"id"`
	Username  string `json:"username"`
}

func CheckingLogin(username, password string) (bool, error) {
	var obj Member

	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd, 
	)

	if err == sql.ErrNoRows {
		fmt.Println("Pengguna tidak ditemukan")
		return false, err
	}

	if err != nil {
		fmt.Println("Query nya salah")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)

	if !match {
		fmt.Println("Hash dan passwordnya tidak sesuai")
		return false, err
	}

	return true, nil

}
