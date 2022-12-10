package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost           string `mapstructure:"MONGO_HOST"`
	DBUserName       string `mapstructure:"MONGO_USER"`
	DBUserPassword   string `mapstructure:"MONGO_PASSWORD"`
	DBName           string `mapstructure:"MONGO_DB"`
	DBPort           string `mapstructure:"MONGO_PORT"`
	ServerPort       string `mapstructure:"PORT"`
	CasbinPolicyFile string `mapstructure:"CASBIN_POLICY_FILE"`
	CasbinConfFile   string `mapstructure:"CASBIN_CONFIG_FILE"`
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
