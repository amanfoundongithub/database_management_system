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


// Verifies the hash of the password
func VerifyHash(source string, hashed string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(source)) ; err != nil {
		return false
	} else {
		return true
	}
}

