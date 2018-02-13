package hasher

import "golang.org/x/crypto/bcrypt"

const defaultCost int = 10

// BCrypt is a Hasher that use the bcrypt algorithm
type BCrypt struct {
	cost int
}

// NewBCrypt allocates a new BCrypt instance
func NewBCrypt() Hasher {
	return &BCrypt{
		cost: defaultCost,
	}
}

// Hash implements the Hasher interface
func (bc BCrypt) Hash(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bc.cost)
}

// Compare implements the Hasher interface
func (bc BCrypt) Compare(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}
