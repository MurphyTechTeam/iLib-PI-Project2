package helpers

import "golang.org/x/crypto/bcrypt"

func PasswordhHash(password string) (string,error){ // Fungsi untuk meng-generate password dari plain ke bcrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error){ // Fungsi mengecek validasi antara hash bcrypt dengan password plain
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil{
		return false,err
	}

	return true, nil
}