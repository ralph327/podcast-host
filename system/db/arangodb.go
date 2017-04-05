// ArangoDB implementation of db package

package db

import (
	"codenex.us/ralph/podcast-host"
	"encoding/json"
	"errors"
	"github.com/solher/arangolite"
)

// A map of podcast-host Structs and their corresponding collection(table) name
var Clxns = map[interface{}]string{
	podhost.User{}:               "users",
	podhost.UserRole{}:           "user_roles",
	podhost.UserHasRole{}:        "user_has_role",
	podhost.Podcast{}:            "podcasts",
	podhost.Episode{}:            "episodes",
	podhost.PodcastHasEpisode{}:  "podcast_has_episode",
	podhost.RssFeed{}:            "rss_feed",
	podhost.PodcastHasRssFeed{}:  "podcast_has_rss_feed",
	podhost.PodcastCategory{}:    "podcast_category",
	podhost.PodcastHasCategory{}: "podcast_has_category",
	podhost.UserHasPodcast{}:     "user_has_podcast",
}

// Extends db.Database with arangolite connection
type ArangoDB struct {
	Database
	DB *arangolite.DB
}

// Establishes a connection and saves DB info
func (d *ArangoDB) InitConnect(url string, name string, user string, pass string) error {
	// Connect
	db := arangolite.New().
		LoggerOptions(false, false, false).
		Connect("http://"+url, name, user, pass)

	// Check Connection
	q := arangolite.NewQuery(`
		RETURN "connected"
	  	`)

	_, err := db.Run(q)

	if err != nil {
		return err
	}

	d.DB = db
	d.URL = url
	d.Name = name
	d.User = user
	d.Pass = pass

	// No Error
	return nil
}

// Connects to the database
func (d *ArangoDB) Connect() (*arangolite.DB, error) {
	// Connect
	d.DB = arangolite.New().
		LoggerOptions(false, false, false).
		Connect("http://"+d.URL, d.Name, d.User, d.Pass)

	// Check Connection
	q := arangolite.NewQuery(`
		RETURN "connected"
	  	`)

	_, err := d.DB.Run(q)

	if err != nil {
		return nil, err
	}

	// No Error
	return d.DB, nil
}

// Check that the model exists in the database
func (d *ArangoDB) ModelCheck() error {
	var err error
	var r []byte

	var cols *[]arangolite.CollectionInfo

	if err != nil {
		return err
	}

	// iterate over structs/collections
	// check for existence and create if not there
	for _, clxn := range Clxns {
		r, _ = d.DB.Run(&arangolite.GetCollectionInfo{
			CollectionName: clxn,
			IncludeSystem:  false},
		)

		json.Unmarshal(r, &cols)

		if cols == nil {
			_, err = d.DB.Run(&arangolite.CreateCollection{Name: clxn})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

/*************************
********          ********
********   USER   ********
********          ********
*************************/
func (d *ArangoDB) GetUser(key string) (*podhost.User, error) {

	q := arangolite.NewQuery(`
		FOR user IN users
			FILTER user._key == "%s"
			RETURN user
	  	`, key)

	users := []*podhost.User{}

	r, err := d.DB.Run(q)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(r, &users)

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	if len(users) > 1 {
		return nil, errors.New("Too many matches")
	}

	user := users[0]

	return user, nil
}

func (d *ArangoDB) CreateUser(u *podhost.User) error {
	user, err := json.Marshal(u)

	if err != nil {
		return err
	}

	q := arangolite.NewQuery(`
		INSERT 
		%s
		INTO users
		`, user)

	_, err = d.DB.Run(q)

	if err != nil {
		return err
	}

	return nil
}

func (d *ArangoDB) DeleteUser(key string) error {
	q := arangolite.NewQuery(`
		UPDATE { _key: "%s" } 
		WITH { Active: false }
		IN users
		`, key)

	_, err := d.DB.Run(q)

	if err != nil {
		return err
	}

	return nil
}
