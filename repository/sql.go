package repository

type ISQLRepository interface {
	GetOne(dest interface{}, query string, args ...interface{}) error
	// dest function should create a new instance everytime it's called
	// it is called for every row in the result
	GetAll(dest func() interface{}, query string, args ...interface{}) error
	Clean()
}
