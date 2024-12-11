package security

import "golang.org/x/crypto/bcrypt"

// Encrypts a string 
func Encrypt(source string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost) 
	if err != nil {
		return "", err 
	} else {
		return string(hashedPassword), nil
	}
}