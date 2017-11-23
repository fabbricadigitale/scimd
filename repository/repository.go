package repository

//Repository is the target interface
type Repository interface {
	Create() error

	Get() error

	GetAll() error

	Count() error

	Update() error

	Delete() error

	Search() error
}

// RepoAdapter is the repository Adapter
type RepoAdapter struct {
	adaptee *Repository
}

// Create is ...
func (a *RepoAdapter) Create() error {
	return (*a.adaptee).Create()
}

// Get is ...
func (a *RepoAdapter) Get() error {
	return (*a.adaptee).Get()
}

// GetAll is ...
func (a *RepoAdapter) GetAll() error {
	return (*a.adaptee).GetAll()
}

// Count ...
func (a *RepoAdapter) Count() error {
	return (*a.adaptee).Count()
}

// Update is ...
func (a *RepoAdapter) Update() error {
	return (*a.adaptee).Update()
}

// Delete is ...
func (a *RepoAdapter) Delete() error {
	return (*a.adaptee).Delete()
}

// Search is ...
func (a *RepoAdapter) Search() error {
	return (*a.adaptee).Search()
}
