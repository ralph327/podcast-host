// Configures the complete server (http, fs, database, etc)
package system

import (
	"github.com/fvbock/endless"
	"github.com/solher/arangolite"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

type System struct {
	Server *gin.Engine
	DB     *arangolite.DB
	Conf   *viper.Viper
	Env    string
}

func NewSystem() (*System, error) {
	s := new(System)

	err := s.init()

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *System) Start() {
	endless.ListenAndServe(":"+s.Conf.GetString(s.Env+"WebPort"), s.Server)
}

func (s *System) init() error {
	var err error

	// HTTP Server
	s.Server = gin.Default()

	// Config
	err = s.LoadConfig()

	if err != nil {
		return err
	}

	return nil
}
