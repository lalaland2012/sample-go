package handler

import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword Function return hasspassword and error
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash Function return result of checking password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func main() {
//     password := "duy"
//     hash := "$2a$14$KxcY4oC0UVQVYiNM75LUxOP/zzR71NhEqIuz5f/wDPbCEvOO5D7si" // ignore error for the sake of simplicity

//     fmt.Println("Password:", password)
//     fmt.Println("Hash:    ", hash)

//     match := CheckPasswordHash(password, hash)
//     fmt.Println("Match:   ", match)
// }
