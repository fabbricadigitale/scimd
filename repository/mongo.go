package repository

import mgo "gopkg.in/mgo.v2"

//Mongo repository adaptee
type repository struct {
	db         string
	collection string
	session    *mgo.Session
}

// CreateRepository factory
func CreateRepository(url, db, collection string) (Repository, error) {
	repo := &repository{}

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
func (r *repository) getCollection() (*mgo.Collection, func()) {
	s := r.session.Copy()
	return s.DB(r.db).C(r.collection), func() { s.Close() }
}

func (r *repository) Create() error {
	// Not yet implemented
	/* c, close := r.getCollection()
	defer close()

	var fakeInput interface{}
	return r.errorWrapper(c.Insert(fakeInput)) */
}

func (r *repository) Get() error {
	// Not yet implemented
	return nil
}

func (r *repository) GetAll() error {
	// Not yet implemented
	return nil
}

func (r *repository) Count() error {
	// Not yet implemented
	return nil
}

func (r *repository) Update() error {
	// Not yet implemented
	return nil
}

func (r *repository) Delete() error {
	// Not yet implemented
	return nil
}

func (r *repository) Search() error {
	// Not yet implemented
	return nil
}

// mongoErrorWrapper translates mongo errors in specific domain errors
func (r *repository) errorWrapper(e error) error {
	// Not yet implemented
	return e
}
