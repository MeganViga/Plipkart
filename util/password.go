package util

import (

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string)(string, error){
	password_byte, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(password_byte), err
}

func VerifyPassword(password string, hashedpassword string) error{
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword),[]byte(password))
	return err
}