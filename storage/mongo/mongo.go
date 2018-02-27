package mongo

import (
	"fmt"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Driver repository adaptee
type Driver struct {
	db         string
	collection string
	session    *mgo.Session
	keys       [][]string
}

// SetIndexes is ...
func (d *Driver) SetIndexes(keys [][]string) {
	d.keys = keys
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
func (d *Driver) getCollection() (*mgo.Collection, func(), error) {
	s := d.session.Copy()
	c, err := d.ensureIndexes(s.DB(d.db).C(d.collection))
	if err != nil {
		return nil, nil, err
	}
	return c, s.Close, nil
}

func (d *Driver) ensureIndexes(c *mgo.Collection) (*mgo.Collection, error) {

	for _, key := range d.keys {

		index := mgo.Index{
			Key:        key,
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}

		err := c.EnsureIndex(index)
		if err != nil {
			fmt.Printf("Error => %+v", err)

			return nil, d.errorWrapper(err)
		}
	}

	return c, nil
}

// Create is the driver method for Create
func (d *Driver) Create(doc *document) error {
	// Not yet implemented
	c, close, err := d.getCollection()
	if err != nil {
		return err
	}
	defer close()

	return d.errorWrapper(c.Insert(doc))
}

// Update is the driver method for Update
func (d *Driver) Update(query bson.M, doc *document) error {
	c, close, err := d.getCollection()
	if err != nil {
		return err
	}
	defer close()

	err = c.Update(query, doc)
	return d.errorWrapper(err, (*doc)["id"])
}

// Delete is the driver method for Delete
func (d *Driver) Delete(query bson.M) error {

	c, close, err := d.getCollection()
	if err != nil {
		return err
	}
	defer close()

	err = c.Remove(query)
	if err != nil {
		return d.errorWrapper(err)
	}

	return nil
}

// Find is the driver method for Find
func (d *Driver) Find(q bson.M) (*mgo.Query, func(), error) {
	c, close, err := d.getCollection()
	if err != nil {
		return nil, nil, err
	}

	query := c.Find(q)
	return query, close, nil
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
