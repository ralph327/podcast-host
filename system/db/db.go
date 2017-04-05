// Domain for the database package

package db

/*************************
********          ********
********    DB    ********
********          ********
*************************/

type Database struct {
	URL  string
	Name string
	User string
	Pass string
}

type DatabaseService interface {
	InitConnect(url string, name string, user string, pass string) error
}
