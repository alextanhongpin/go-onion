package database

// Client represents all the method available for the database
type Client interface {
	Query(string) string
}

type database struct{}

func (db *database) Query(string) string {
	return "Mock response from DB"
}

// New returns a new instance of connection to the database
func New() Client {
	return &database{}
}
