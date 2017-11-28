package mongo

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Driver repository adaptee
type Driver struct {
	db         string
	collection string
	session    *mgo.Session
}

// CreateDriver factory
func CreateDriver(url, db, collection string) (*Driver, error) {
	repo := &Driver{}

	var session *mgo.Session
	var err error

	if session, err = mgo.Dial(url); err != nil {
		return nil, err
	}

	repo.session = session
	repo.db = db
	repo.collection = collection

	return repo, nil
}

// The new session create with Copy() will share the same cluster information and connection pool.
// Every session created must have its Close method called at the end of its life time.
// This pattern allows to take a full advantage of concurrency.
func (d *Driver) getCollection() (*mgo.Collection, func()) {
	s := d.session.Copy()
	return s.DB(d.db).C(d.collection), func() { s.Close() }
}

// Create is the adapter method for Create
func (d *Driver) Create(res *HResource) error {
	// Not yet implemented
	c, close := d.getCollection()
	defer close()

	return d.errorWrapper(c.Insert(res))
}

// Get is the adapter method for Get
func (d *Driver) Get(id string) (*HResource, error) {
	//not yet implemented
	c, close := d.getCollection()
	defer close()

	data := &HResource{}
	var query bson.M

	query = bson.M{"id": id}

	err := c.Find(query).One(&data)
	if err != nil {
		return nil, d.errorWrapper(err, id)
	}

	return data, nil
}

// Count is the adapter method for Count
func (d *Driver) Count() error {
	// Not yet implemented
	return nil
}

// Update is the adapter method for Update
func (d *Driver) Update(id string, resource *HResource) error {
	c, close := d.getCollection()
	defer close()

	var query bson.M
	query = bson.M{"id": id}

	err := c.Update(query, *resource)
	return d.errorWrapper(err, resource.Data[0]["id"])
}

// Delete is the adapter method for Delete
func (d *Driver) Delete(id string) error {

	c, close := d.getCollection()
	defer close()

	var query bson.M
	query = bson.M{"id": id}
	err := c.Remove(query)
	if err != nil {
		return d.errorWrapper(err, id)
	}

	return nil
}

// Search is the adapter method for Search
func (d *Driver) Search() error {
	// Not yet implemented
	return nil
}

// mongoErrorWrapper translates mongo errors in specific domain errors
func (d *Driver) errorWrapper(e error, args ...interface{}) error {
	// Not yet implemented

	if e == nil {
		return nil
	}
	switch {
	case e.Error() == "not found":
		if len(args) > 1 {
			return ResourceNotFoundError{
				fmt.Sprintf("%v", args[0]),
				fmt.Sprintf("%v", args[1]),
			}
		} else if len(args) > 0 {
			return ResourceNotFoundError{
				fmt.Sprintf("%v", args[0]),
				"",
			}
		}
		return ResourceNotFoundError{
			"",
			"",
		}
	default:
		return e
	}
}

// ResourceNotFoundError is ...
type ResourceNotFoundError struct {
	msg string
	id  string
}

func (r ResourceNotFoundError) Error() string {
	return fmt.Sprintf("%v - id %v", r.msg, r.id)
}
