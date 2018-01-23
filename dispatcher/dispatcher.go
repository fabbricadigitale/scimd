package dispatcher

import (
	"github.com/olebedev/emitter"
)

// Dispatcher is a struct containing a topic emitter
type Dispatcher struct {
	emitter *emitter.Emitter
}

var _ Emitter = (*Dispatcher)(nil)

// New creates a new Dispatcher
func New(capacity uint) *Dispatcher {
	return &Dispatcher{
		emitter: emitter.New(capacity),
	}
}

// Emitter returns the underlying topic emitter
func (o *Dispatcher) Emitter() *emitter.Emitter {
	return o.emitter
}
