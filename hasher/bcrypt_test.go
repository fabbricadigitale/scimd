package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBCrypt_Hash(t *testing.T) {
	password := []byte("password")
	expected := []byte("$2a$10$nnKMazO/AsNRayhioXqb1.WNNcEkDPjj3/ownOU3jIil7aXZRnXNC")
	hasher := NewBCrypt()
	hashedPassword, err := hasher.Hash(password)

	assert.Nil(t, err)
	assert.True(t, len(expected) == len(hashedPassword))
}

func TestBCryptHasher_Compare(t *testing.T) {
	password := []byte("password")
	hasher := NewBCrypt()
	hashedPassword, err := hasher.Hash(password)
	assert.Nil(t, err)
	assert.True(t, hasher.Compare(hashedPassword, password))
}
