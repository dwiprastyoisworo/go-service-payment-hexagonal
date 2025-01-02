package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() (*Config, error) {
	var configFile Config
	var apps AppsConfig
	var mongodb MongoDBConfig

	// Set default lokasi file .env
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	// Membaca file .env
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return &configFile, err
	}

	// Membaca environment variables
	viper.AutomaticEnv()

	// Memetakan ke struct
	if err := viper.Unmarshal(&apps); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
		return &configFile, err
	}

	// Memetakan ke struct
	if err := viper.Unmarshal(&mongodb); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
		return &configFile, err
	}

	configFile.Apps = apps
	configFile.MongoDB = mongodb

	return &configFile, nil
}
