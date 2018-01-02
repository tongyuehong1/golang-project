package utility

import "golang.org/x/crypto/bcrypt"

func CompareHash(digest []byte, password string) bool {
	hex := []byte(password)
	if err := bcrypt.CompareHashAndPassword(digest, hex); err == nil {
		return true
	}

	return false
}

// bcrypto
func GenerateHash(password string) ([]byte, error) {
	hex := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(hex, 10)

	if err != nil {
		return hashedPassword, err
	}

	return hashedPassword, nil
}
