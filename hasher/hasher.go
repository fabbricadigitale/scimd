package hasher

// Hasher can generates hashes from plaintext password and compare hashes with plaintext passwords
type Hasher interface {

	// Hash returns the hash of its argument
	Hash(b []byte) ([]byte, error)

	// Compare compares the hashed password with a plaintext password
	// It returns true in case of successfull comparation, false otherwise
	Compare(hash, b []byte) bool
}
