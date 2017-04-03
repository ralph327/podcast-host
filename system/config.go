// Accesses configuration files
package system

import (
	_ "github.com/spf13/viper"
)

func LoadConfig() (*Viper, error) {
	v := New()

	err := v.init()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.ReadInConfig() // Find and read the config file

	if err != nil {
		panic(fmt.Errorf("Fatal error: config file not found: %s \n", err))
	}

	if GetBool("WatchConf") {
		v.WatchConfig()
	}

}

func (v *Viper) init() error {
	// config file name
	viper.SetConfigName("podcasthost_conf")

	// config path
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.podcasthost")
	viper.AddConfigPath("/etc/podcasthost")

	// config environment variables
	SetEnvPrefix("PH")
	BindEnv("Environment", "PH_Environment")

	// config defaults
	viper.SetDefault("Environment", "development")
	v.SetDefaults()

	return nil
}

func (v *Viper) SetDefaults() {
	viper.SetDefault("WatchConf", "true")
	viper.SetDefault("WebPort", "8080")
	viper.SetDefault("DBPort", "8529")
	viper.SetDefault("Hostname", "localhost")
}
