package config

import "github.com/spf13/viper"

type Config struct {
	HOST            string `mapstructure:"HOST"`
	USERNAME        string `mapstructure:"USERNAME"`
	PORT            string `mapstructure:"PORT"`
	PASSWORD        string `mapstructure:"PASSWORD"`
	DBNAME          string `mapstructure:"DBNAME"`
	SSLMODE         string `mapstructure:"SSLMODE"`
	GRPCUSERPORT    string `mapstructure:"GRPCUSERPORT"`
	GRPCSERVICEPORT string `mapstructure:"GRPCSERVICEPORT"`
	GRPCADMINPORT   string `mapstructure:"GRPCADMINPORT"`
}

func LoadConfig() *Config{
	var cfg *Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&cfg)
	return cfg
}