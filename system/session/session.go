// Implements user sessions on the server

package session

import (
	"github.com/ralph327/sessions-1"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
)

type Session struct {
	SHandler gin.HandlerFunc
	RStore   sessions.RedisStore
}

var SecretKey = []byte("secret")

func NewRedisStore(url string, secret []byte) (sessions.RedisStore, error) {
	store, err := sessions.NewRedisStore(10, "tcp", url, "", secret)
	if err != nil {
		log.Println("Failing here")
		return nil, err
	}

	return store, nil
}

func New(url string, secret []byte) (*Session, error) {
	var err error
	s := new(Session)

	s.RStore, err = NewRedisStore(url, secret)
	if err != nil {
		return nil, err
	}

	s.SHandler = sessions.Sessions("podcast-host", s.RStore)

	return s, nil
}

func (s *Session) Get() gin.HandlerFunc {
	return s.SHandler
}
