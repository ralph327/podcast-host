// Accesses configuration files

package system

import (
	"github.com/spf13/viper"
	"log"
)

// Creates a new config and loads it
func (s *System) LoadConfig() error {
	s.Conf = viper.New()

	// Initialize and set defaults
	err := s.ConfigInit()

	if err != nil {
		return err
	}

	err = s.Conf.ReadInConfig() // Find and read the config file

	if err != nil {
		return err
	}

	if s.Conf.GetBool("WatchConf") {
		s.Conf.WatchConfig()
	}

	watch := s.Conf.GetBool("WatchConf")
	var ws string
	if watch {
		ws = "TRUE"
	} else {
		ws = "FALSE"
	}
	log.Println("WatchConf: " + ws)

	env := s.Conf.GetString("Environment")
	log.Println("Env: " + env)

	// Set environment of system based on config
	switch env {
	case "development", "dev", "test", "testing", "local", "DEV", "DEVELOPMENT", "TEST", "LOCAL":
		s.Conf.Set("ENV", "development.")
	case "staging", "model", "acceptance", "uat", "remote", "STAGING", "MODEL", "ACCEPTANCE", "UAT", "REMOTE":
		s.Conf.Set("ENV", "staging.")
	case "prod", "production", "live", "PROD", "PRODUCTION", "LIVE":
		s.Conf.Set("ENV", "production.")
	}

	// Setup the URLs used in the system
	webport := s.Conf.GetString(s.Conf.GetString("Env") + "WebPort")
	weburl := s.Conf.GetString(s.Conf.GetString("Env")+"Hostname") + ":" + webport

	dbport := s.Conf.GetString(s.Conf.GetString("Env") + "DBPort")
	dburl := s.Conf.GetString(s.Conf.GetString("Env")+"Hostname") + ":" + dbport

	s.Conf.Set("BaseURL", weburl)
	s.Conf.Set("DBURL", dburl)

	return nil
}

// Initialize the config
func (s *System) ConfigInit() error {
	// config file name
	s.Conf.SetConfigName("podcasthost_conf")

	// config path
	s.Conf.AddConfigPath("./config")
	s.Conf.AddConfigPath("$HOME/.podcasthost")
	s.Conf.AddConfigPath("/etc/podcasthost")

	// config environment variables
	viper.SetEnvPrefix("PH")
	viper.BindEnv("Environment", "PH_Environment")

	// config defaults
	s.Conf.SetDefault("Environment", "development")
	s.SetConfigDefaults()

	return nil
}

// Set defaults for the config
func (s *System) SetConfigDefaults() {
	s.Conf.SetDefault("WatchConf", true)
	s.Conf.SetDefault("development.WebPort", "8080")
	s.Conf.SetDefault("development.DBPort", "8529")
	s.Conf.SetDefault("development.Hostname", "localhost")
	s.Conf.SetDefault("development.SiteName", "Podcast Host")
}
