package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"MONGO_HOST"`
	DBUserName     string `mapstructure:"MONGO_USER"`
	DBUserPassword string `mapstructure:"MONGO_PASSWORD"`
	DBName         string `mapstructure:"MONGO_DB"`
	DBPort         string `mapstructure:"MONGO_PORT"`
	ServerPort     string `mapstructure:"PORT"`
	KCBaseURL      string `mapstructure:"KC_BASE_URL"`
	KCClientID     string `mapstructure:"KC_CLIENT_ID"`
	KCClientSecret string `mapstructure:"KC_CLIENT_SECRET"`
	KCRealm        string `mapstructure:"KC_REALM"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
