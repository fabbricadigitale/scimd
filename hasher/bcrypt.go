package hasher

import "golang.org/x/crypto/bcrypt"

const cost int = 10

// BCryptHasher is a Hasher that use bcrypt algorithm
type BCryptHasher struct {
	cost int
}

// NewBCryptHasher allocates new NewBCryptHasher
func NewBCryptHasher() Hasher {
	return &BCryptHasher{}
}

// Hash implements Hasher interface
func (b BCryptHasher) Hash(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

// Compare implements Hasher interface
func (b BCryptHasher) Compare(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}
