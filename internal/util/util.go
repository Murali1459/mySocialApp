package util

import "golang.org/x/crypto/bcrypt"

func IsSamePassword(pw, hashedPw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(pw)) == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
