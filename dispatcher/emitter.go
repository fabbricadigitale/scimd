package dispatcher

import (
	"github.com/olebedev/emitter"
)

// Emitter is the interface implemented by an object capable to return an emitter.Emitter
type Emitter interface {
	Emitter() *emitter.Emitter
}
