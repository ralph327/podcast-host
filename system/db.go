// ArangoDB implementation of podhost

package system

import (
	"codenex.us/ralph/podcast-host"
	"github.com/solher/arangolite"

	"encoding/json"
	"errors"
)

/*************************
********          ********
********    DB    ********
********          ********
*************************/

func (s *System) DBConnect() {
	s.DB = arangolite.New().
		LoggerOptions(false, false, false).
		Connect("http://"+s.Conf.GetString("URL"), "podcast-host", "devpod", "**DB!!podhost")
}

/*************************
********          ********
********   USER   ********
********          ********
*************************/

type UserService struct {
	DB *arangolite.DB
}

func (s *UserService) User(key string) (*podhost.User, error) {
	q := arangolite.NewQuery(`
		FOR user IN users
			FILTER user._key == "%s"
			RETURN user
	  	`, key)

	users := []*podhost.User{}

	r, err := s.DB.Run(q)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(r, &users)

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("No matches")
	}

	if len(users) > 1 {
		return nil, errors.New("Too many matches")
	}

	user := users[0]

	return user, nil
}

func (s *UserService) CreateUser(u *podhost.User) error {
	user, err := json.Marshal(u)

	if err != nil {
		return err
	}

	q := arangolite.NewQuery(`
		INSERT 
		%s
		INTO users
		`, user)

	_, err = s.DB.Run(q)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(key string) error {
	q := arangolite.NewQuery(`
		UPDATE { _key: "%s" } 
		WITH { Active: false }
		IN users
		`, key)

	_, err := s.DB.Run(q)

	if err != nil {
		return err
	}

	return nil
}
