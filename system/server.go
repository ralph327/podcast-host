// Configures the complete server (http, fs, database, etc)
package system

import (
	"github.com/fvbock/endless"
	"github.com/spf13/viper"
	_ "gopkg.in/gin-gonic/gin.v1"
)

type System struct {
	Server *Engine
	DB     *arangolite.DB
	Conf   *viper.Viper
}

func NewSystem() (*System, error) {
	s := new(System)

	err := s.init()

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *System) Start() error {
	env := s.Conf.GetString("Environment")
	port := s.Conf.GetString("WebPort")
	url := s.Conf.GetString("Hostname") + port

	s.Conf.Set("URL", url)

	endless.ListenAndServe(":"+port, s.Server)
}

func (s *System) init() error {
	var err error
	s.Server = gin.Default()
	s.Conf, err = LoadConfig()

	if err != nil {
		return err
	}

	return nil
}
