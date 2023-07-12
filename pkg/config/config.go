package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	ProductSuvUrl string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSuvUrl   string `mapstructure:"ORDER_SVC_URL"`
}

func LoadConfig() (c Config, err error) {

	// AddConfigPath adds the directory where the configuration file is located.
	viper.AddConfigPath("./")

	// SetConfigName sets the name of the configuration file to be read.
	viper.SetConfigName(".env")

	// SetConfigType sets the type of the configuration file.
	viper.SetConfigType("env")

	// AutomaticEnv enables automatic binding of environment variables to configuration values.
	viper.AutomaticEnv()

	// ReadInConfig reads the configuration file with the specified name and type.
	err = viper.ReadInConfig()

	// Check if there was an error reading the configuration file.
	if err != nil {
		return
	}

	// Unmarshal reads the configuration settings into the Config struct.
	if err = viper.Unmarshal(&c); err != nil {
		return
	}
	return
}
