// Configures the complete server (http, fs, database, etc)
package system

import (
	"codenex.us/ralph/podcast-host/system/db"
	"github.com/fvbock/endless"
	"github.com/olahol/melody"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"os"
)

// Group routing, websockets, db and config
type System struct {
	Server *gin.Engine
	WS     *melody.Melody
	DB     *db.ArangoDB
	Conf   *viper.Viper
	Env    string
}

// Creates the system
func NewSystem() (*System, error) {
	s := new(System)

	err := s.init()

	if err != nil {
		return nil, err
	}

	return s, nil
}

// Runs the system with graceful restarts
func (s *System) Start() {
	err := endless.ListenAndServe(":"+s.Conf.GetString(s.Env+"WebPort"), s.Server)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// Initializes the system
func (s *System) init() error {
	var err error

	// HTTP Server
	s.Server = gin.Default()

	// Websocket
	s.WS = melody.New()

	// Config
	err = s.LoadConfig()

	if err != nil {
		return err
	}

	// Database
	s.DB = new(db.ArangoDB)
	err = s.DB.InitConnect(s.Conf.GetString("DBURL"), s.Conf.GetString(s.Env+"DBName"), s.Conf.GetString(s.Env+"DBUser"), s.Conf.GetString(s.Env+"DBPass"))

	if err != nil {
		return err
	}

	err = s.DB.ModelCheck()

	if err != nil {
		return err
	}

	// Routes
	s.AddRoutes()

	return nil
}
