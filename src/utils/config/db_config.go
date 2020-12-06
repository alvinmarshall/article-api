package config

import "os"

type dbConfig struct {
	Driver   string
	Host     string
	Name     string
	User     string
	Password string
	Port     string
}

func GetDBConfig() *dbConfig {
	return &dbConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
	}
}
