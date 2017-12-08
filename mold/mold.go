package mold

import (
	m "gopkg.in/go-playground/mold.v2"
)

// Error represents an error occuring during transformations
type Error struct {
	message string
}

func (e Error) Error() string {
	return e.message
}

// Transformer is a singleton instance providing validation functionalities
var Transformer *m.Transformer

func init() {
	Transformer = m.New()
	Transformer.Register("min", min)
	Transformer.Register("max", max)
	Transformer.Register("normurn", normurn)
}
