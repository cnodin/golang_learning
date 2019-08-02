package auth

import "golang.org/x/crypto/bcrypt"

func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func Compare(hashedBytes, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedBytes), []byte(password))
}
