package hasher

// Hasher can generates hashes from plaintext password and compare hashes with plaintext passwords
type Hasher interface {

	// Hash returns the hash of the password
	Hash(password []byte) ([]byte, error)

	// Compare compares the hashed password with a plaintext password
	// Return true in case of successfull comparation, false otherwise
	Compare(hash, password []byte) bool
}
