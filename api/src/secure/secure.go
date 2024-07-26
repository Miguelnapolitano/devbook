package secure

import "golang.org/x/crypto/bcrypt"

//Transform a password into a hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//Compare a hash and a password and returns if that's equals
func Verify(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}