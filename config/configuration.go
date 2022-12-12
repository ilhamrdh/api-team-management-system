package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server   SetupServer
	Database SetupDatabase
}

type SetupServer struct {
	Port int `mapstructure:"PORT"`
}
type SetupDatabase struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBUser   string `mapstructure:"DB_USER"`
	DBPass   string `mapstructure:"DB_PASS"`
	DBName   string `mapstructure:"DB_NAME"`
	DBHost   string `mapstructure:"DB_HOST"`
	DBPort   int    `mapstructure:"DB_PORT"`
}

var Config *Configuration

func SetupConfiguration() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".development")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file ", err)
		return
	}

	Config = new(Configuration)
	err := viper.Unmarshal(&Config)

	if err != nil {
		log.Fatal("Unable to decode into struct ", err)
		return
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
