package listeners

import (
	"github.com/fabbricadigitale/scimd/hasher"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/olebedev/emitter"
)

// AddListeners for the emitted events
func AddListeners(e *emitter.Emitter) {

	// Create event handler
	e.On("create", func(event *emitter.Event) {
		res, ok := event.Args[0].(*resource.Resource)

		if ok != true {
			return
		}

		hashPassword(res)

	})

	// Update event handler
	e.On("update", func(event *emitter.Event) {
		res, ok := event.Args[0].(*resource.Resource)

		if ok != true {
			return
		}

		hashPassword(res)

	})
}

// hash the password value if there is the password attribute in the resources
func hashPassword(res *resource.Resource) {
	values := res.Values("urn:ietf:params:scim:schemas:core:2.0:User")
	if values == nil {
		return
	}

	passwordValue, ok := (*values)["password"]
	if ok != true {
		return
	}

	password := []byte(passwordValue.(datatype.String))

	hashedPassword, err := hasher.NewBCryptHasher().Hash(password)

	if err != nil {
		panic(err)
	}

	res.SetValues("urn:ietf:params:scim:schemas:core:2.0:User", &datatype.Complex{
		"password": datatype.String(hashedPassword),
	})

}
