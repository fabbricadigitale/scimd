package event

import "github.com/olebedev/emitter"

// Dispatcher is ...
type Dispatcher interface {
	Emitter() *emitter.Emitter
}

type d struct {
	emitter *emitter.Emitter
}

// NewDispatcher creates a new `Dispatcher`
func NewDispatcher(capacity uint) Dispatcher {
	return &d{
		emitter: emitter.New(capacity),
	}
}

func (d *d) Emitter() *emitter.Emitter {
	return d.emitter
}

var _ Dispatcher = (*d)(nil)
