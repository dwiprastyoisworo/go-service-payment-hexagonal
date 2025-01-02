package config

import "time"

type Config struct {
	Apps    AppsConfig
	MongoDB MongoDBConfig
}

type AppsConfig struct {
	Name    string        `mapstructure:"APP_NAME"`
	Version string        `mapstructure:"VERSION"`
	Port    string        `mapstructure:"PORT"`
	Address string        `mapstructure:"ADDRESS"`
	Timeout time.Duration `mapstructure:"TIMEOUT"`
}

type MongoDBConfig struct {
	Url      string `mapstructure:"MONGO_DB_URL"`
	Port     string `mapstructure:"MONGO_DB_PORT"`
	Username string `mapstructure:"MONGO_DB_USERNAME"`
	Password string `mapstructure:"MONGO_DB_PASSWORD"`
	Database string `mapstructure:"MONGO_DB_DATABASE"`
}
