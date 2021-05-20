package models

import (
	"database/sql"
	"fmt"
	"mhilmi999/project-2-mhilmi999/db"
	"mhilmi999/project-2-mhilmi999/helpers"
	
	
	"net/http"
	"golang.org/x/crypto/bcrypt"
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

func RegisMember(username, password string) (Response, error){
	var res Response

	con := db.CreateCon()
	
	sqlStatement := "INSERT INTO users SET username=?, password=?"

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil{
		fmt.Println("error abis hashing")
		return res, err
	}

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		fmt.Println("con prepare error")
		return res,err
	}
	
	result, err := stmt.Exec(username, hashedPass)
	
	if err != nil {
		fmt.Println("exec error")
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil{
		return res,err 
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"Id yang terakhir dimasukan adalah" : lastInsertedId,
	}

	return res, nil


}