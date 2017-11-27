package mongo

import (
	mgo "gopkg.in/mgo.v2"
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

func (d *Driver) Create(res *HResource) error {
	// Not yet implemented
	c, close := d.getCollection()
	defer close()

	return d.errorWrapper(c.Insert(res))
}

func (d *Driver) Get(id, version string) error {
	//not yet implemented
	_, close := d.getCollection()
	defer close()

	return nil
}

func (d *Driver) Count() error {
	// Not yet implemented
	return nil
}

func (d *Driver) Update() error {
	// Not yet implemented
	return nil
}

func (d *Driver) Delete(id, version string) error {
	return nil
}

func (d *Driver) Search() error {
	// Not yet implemented
	return nil
}

// mongoErrorWrapper translates mongo errors in specific domain errors
func (d *Driver) errorWrapper(e error) error {
	// Not yet implemented
	return e
}
