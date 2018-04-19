package listeners

import (
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/hasher"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/storage/mongo"
	"github.com/olebedev/emitter"
	set "gopkg.in/fatih/set.v0"
)

// AddListeners for the emitted events
func AddListeners(e *emitter.Emitter) {

	resTypeRepo := core.GetResourceTypeRepository()

	// Create event handler
	e.On("create", func(event *emitter.Event) {
		res, ok := event.Args[0].(*resource.Resource)
		adapter, ok := event.Args[1].(*mongo.Adapter)

		if ok != true {
			return
		}

		resType := resTypeRepo.Pull(res.Meta.ResourceType)

		relations, err := attr.GetRelationships(resType.Schema, resType.ID)
		if err != nil {
			return
		}
		for _, relation := range relations {
			addMembership(relation.RWAttribute, relation.ROAttribute, res, &relation.ROResourceType, *adapter)
		}

		hashPassword(res)
	})

	// Update event handler
	e.On("update", func(event *emitter.Event) {
		res, ok := event.Args[0].(*resource.Resource)
		adapter, ok := event.Args[1].(*mongo.Adapter)

		if ok != true {
			return
		}

		resType := resTypeRepo.Pull(res.Meta.ResourceType)

		relations, err := attr.GetRelationships(resType.Schema, resType.ID)
		if err != nil {
			return
		}
		for _, relation := range relations {
			// stored resource
			mRes, _ := adapter.DoGet(res.ResourceType(), res.ID, "", nil)
			updateMembership(relation.RWAttribute, relation.ROAttribute, mRes, res, &relation.ROResourceType, *adapter)
		}

		hashPassword(res)
	})

	e.On("patch", func(event *emitter.Event) {
		// stored resource
		mResource, ok := event.Args[0].(*resource.Resource)
		if ok != true {
			return
		}
		// current resource
		resource, ok := event.Args[1].(*resource.Resource)
		if ok != true {
			return
		}
		adapter, ok := event.Args[2].(*mongo.Adapter)
		if ok != true {
			return
		}

		resType := resTypeRepo.Pull(resource.Meta.ResourceType)

		relations, err := attr.GetRelationships(resType.Schema, resType.ID)
		if err != nil {
			return
		}
		for _, relation := range relations {
			updateMembership(relation.RWAttribute, relation.ROAttribute, mResource, resource, &relation.ROResourceType, *adapter)
		}
	})

	e.On("patchPassword", func(event *emitter.Event) {
		values, ok := event.Args[0].(*storage.PContainer)
		if ok != true {
			return
		}

		v := values.Value.(string)
		(*values).Value = hashString((datatype.String)(v))
	})

	e.On("delete", func(event *emitter.Event) {
		rt, ok := event.Args[0].(*core.ResourceType)
		id, ok := event.Args[1].(string)
		adapter, ok := event.Args[3].(*mongo.Adapter)

		if ok != true {
			return
		}

		resType := resTypeRepo.Pull(rt.ID)

		relations, err := attr.GetRelationships(resType.Schema, resType.ID)
		if err != nil {
			return
		}
		for _, relation := range relations {
			deleteMembership(relation.RWAttribute, relation.ROAttribute, id, rt, &relation.ROResourceType, *adapter)
		}

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

	hashedPassword, err := hasher.NewBCrypt().Hash(password)

	if err != nil {
		panic(err)
	}

	(*values)["password"] = hashedPassword

	res.SetValues("urn:ietf:params:scim:schemas:core:2.0:User", values)
}

func hashString(pwd datatype.String) string {

	password := []byte(pwd)

	hashedPassword, err := hasher.NewBCrypt().Hash(password)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func addMembership(rw attr.Path, ro attr.Path, res *resource.Resource, roResType *core.ResourceType, adapter mongo.Adapter) error {

	rwValues := res.Values(rw.URI)

	if _, ok := (*rwValues)[rw.Name]; !ok {
		return nil
	}

	members := (*rwValues)[rw.Name].([]datatype.DataTyper)
	displayName := (*rwValues)["displayName"]

	for _, m := range members {

		var value interface{}

		switch m.(type) {
		case datatype.Complex:
			value = m.(datatype.Complex)["value"]
			break
		case *datatype.Complex:
			value = (*(m.(*datatype.Complex)))["value"]
			break
		default:
			break
		}

		roResource, err := adapter.DoGet(roResType, string(value.(datatype.String)), "", nil)
		if err != nil {
			return err
		}
		if roResource == nil {
			return err
		}

		roValues := roResource.Values(ro.URI)

		switch (*roValues)[ro.Name].(type) {
		case []datatype.DataTyper:
			(*roValues)[ro.Name] = addReferenceToList((*roValues)[ro.Name], res.ID, res.Meta.Location, string(displayName.(datatype.String)))
			break
		case datatype.DataTyper:
			(*roValues)[ro.Name] = addReference((*roValues)[ro.Name], res.ID, res.Meta.Location, string(displayName.(datatype.String)))
			break
		case nil:
			if ro.Context(roResType).Attribute.MultiValued {
				(*roValues)[ro.Name] = make([]datatype.DataTyper, 0)
				(*roValues)[ro.Name] = addReferenceToList((*roValues)[ro.Name], res.ID, res.Meta.Location, string(displayName.(datatype.String)))
			} else {
				(*roValues)[ro.Name] = datatype.Complex{}
				(*roValues)[ro.Name] = addReference((*roValues)[ro.Name], res.ID, res.Meta.Location, string(displayName.(datatype.String)))
			}
			break
		default:
			break
		}

		roResource.SetValues(ro.URI, roValues)
		err = adapter.DoUpdate(roResource, roResource.ID, "")
		if err != nil {
			return err
		}

	}

	return nil
}

func updateMembership(rw attr.Path, ro attr.Path, mRes *resource.Resource, res *resource.Resource, roResType *core.ResourceType, adapter mongo.Adapter) error {

	values := mRes.Values(rw.URI)
	// collection of the stored resource
	mCollection := (*values)[rw.Name]

	if mRes == nil || map[string]interface{}(*mRes.Values(rw.URI))[rw.Name] == nil {
		addMembership(rw, ro, res, roResType, adapter)
		return nil
	}

	// collection of the current resource
	cCollection := make([]datatype.DataTyper, 0)
	c := (*res.Values(rw.URI))[rw.Name]
	if c != nil {
		cCollection = c.([]datatype.DataTyper)
	}

	displayName := (*res.Values(rw.URI))["displayName"]
	s := set.New()
	t := set.New()

	for _, item := range mCollection.([]datatype.DataTyper) {
		i := item.(*datatype.Complex)
		s.Add((*i)["value"].(datatype.String))
	}
	for _, item1 := range cCollection {

		switch item1.(type) {
		case datatype.Complex:
			t.Add((item1.(datatype.Complex))["value"].(datatype.String))
			break
		case *datatype.Complex:
			it := item1.(*datatype.Complex)
			t.Add((*it)["value"].(datatype.String))
		}

	}

	// reference to add
	addIDs := set.Difference(set.Union(s, t), s).List()

	// reference to remove
	removeIDs := set.Difference(set.Union(s, t), t).List()

	for _, d := range removeIDs {

		roResource, err := adapter.DoGet(roResType, string(d.(datatype.String)), "", nil)
		if err != nil {
			return err
		}
		if roResource == nil {
			return err
		}

		roValues := roResource.Values(ro.URI)

		switch (*roValues)[ro.Name].(type) {
		case []datatype.DataTyper:
			(*roValues)[ro.Name] = removeReferenceFromList((*roValues)[ro.Name], res.ID)
			break
		case datatype.DataTyper:
			(*roValues)[ro.Name] = removeReference((*roValues)[ro.Name], res.ID)
			break
		default:
			break
		}

		roResource.SetValues(ro.URI, roValues)
		err = adapter.DoUpdate(roResource, roResource.ID, "")
		if err != nil {
			return err
		}
	}

	for _, a := range addIDs {
		roResource, err := adapter.DoGet(roResType, string(a.(datatype.String)), "", nil)
		if err != nil {
			return err
		}
		if roResource == nil {
			return err
		}

		roValues := roResource.Values(ro.URI)

		switch (*roValues)[ro.Name].(type) {
		case []datatype.DataTyper:
			(*roValues)[ro.Name] = addReferenceToList((*roValues)[ro.Name], mRes.ID, mRes.Meta.Location, string(displayName.(datatype.String)))
			break
		case datatype.DataTyper:
			(*roValues)[ro.Name] = addReference((*roValues)[ro.Name], mRes.ID, mRes.Meta.Location, string(displayName.(datatype.String)))
			break
		case nil:

			if ro.Context(roResType).Attribute.MultiValued {
				(*roValues)[ro.Name] = make([]datatype.DataTyper, 0)
				(*roValues)[ro.Name] = addReferenceToList((*roValues)[ro.Name], mRes.ID, mRes.Meta.Location, string(displayName.(datatype.String)))
			} else {
				(*roValues)[ro.Name] = datatype.Complex{}
				(*roValues)[ro.Name] = addReference((*roValues)[ro.Name], mRes.ID, mRes.Meta.Location, string(displayName.(datatype.String)))
			}
			break
		default:
			break
		}

		roResource.SetValues(ro.URI, roValues)
		err = adapter.DoUpdate(roResource, roResource.ID, "")
		if err != nil {
			return err
		}
	}

	return nil
}

func addReferenceToList(parent interface{}, ID, location, display string) interface{} {

	reference := parent.([]datatype.DataTyper)

	found := false
	for _, g := range reference {

		gValue := g.(*datatype.Complex)
		if string((*gValue)["value"].(datatype.String)) == ID {
			found = true
			break
		}
	}

	if !found {

		membership := datatype.Complex{}
		membership["value"] = ID
		membership["$ref"] = location
		membership["display"] = display

		reference = append(reference, datatype.Complex(membership))
		parent = reference
	}

	return parent
}

func addReference(parent interface{}, ID, location, display string) interface{} {
	membership := datatype.Complex{}
	membership["value"] = ID
	membership["$ref"] = location
	membership["display"] = display
	parent = membership
	return parent
}

func removeReferenceFromList(parent interface{}, ID string) interface{} {
	list := parent.([]datatype.DataTyper)
	index := -1
	for i, g := range list {

		gValue := g.(*datatype.Complex)
		if string((*gValue)["value"].(datatype.String)) == ID {
			index = i
			break
		}
	}
	if index != -1 {
		list = append(list[:index], list[index+1:]...)
		parent = list
	}
	return parent
}

func removeReference(parent interface{}, ID string) interface{} {
	obj := parent.(datatype.DataTyper)
	v := obj.(*datatype.Complex)
	if string((*v)["value"].(datatype.String)) == ID {
		parent = nil
	}
	return parent
}

func deleteMembership(rw attr.Path, ro attr.Path, ID string, resType *core.ResourceType, roResType *core.ResourceType, adapter mongo.Adapter) error {

	res, err := adapter.DoGet(resType, ID, "", nil)
	if err != nil {
		return err
	}

	rwValues := res.Values(rw.URI)
	m := (*rwValues)[rw.Name]

	if m == nil {
		return nil
	}

	members := m.([]datatype.DataTyper)

	for _, m := range members {

		v := m.(*datatype.Complex)
		value := (*v)["value"]

		roResource, err := adapter.DoGet(roResType, string(value.(datatype.String)), "", nil)
		if err != nil {
			return err
		}
		if roResource == nil {
			return err
		}

		roValues := roResource.Values(ro.URI)

		switch (*roValues)[ro.Name].(type) {
		case []datatype.DataTyper:
			(*roValues)[ro.Name] = removeReferenceFromList((*roValues)[ro.Name], res.ID)
			break
		case datatype.DataTyper:
			(*roValues)[ro.Name] = removeReference((*roValues)[ro.Name], res.ID)
			break
		default:
			break
		}

		roResource.SetValues(ro.URI, roValues)
		err = adapter.DoUpdate(roResource, roResource.ID, "")
		if err != nil {
			return err
		}
	}
	return nil
}
