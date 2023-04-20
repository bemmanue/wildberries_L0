package server

import (
	"errors"
	"os"
	"strconv"
)

// Config ...
type Config struct {
	BindAddr string
	Database *DatabaseConfig
}

// DatabaseConfig ...
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

// NewConfig ...
func NewConfig() (*Config, error) {
	bindAddr, exists := os.LookupEnv("BIND_ADDR")
	if !exists {
		return nil, errors.New("BIND_ADDR not found")
	}

	databaseConfig, err := NewDatabaseConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		BindAddr: bindAddr,
		Database: databaseConfig,
	}, nil
}

// NewDatabaseConfig ...
func NewDatabaseConfig() (*DatabaseConfig, error) {
	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		return nil, errors.New("DB_HOST not found")
	}

	portS, exists := os.LookupEnv("DB_PORT")
	if !exists {
		return nil, errors.New("DB_PORT not found")
	}

	port, err := strconv.Atoi(portS)
	if err != nil {
		return nil, errors.New("PORT is invalid")
	}

	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		return nil, errors.New("DB_USER not found")
	}

	password, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		return nil, errors.New("DB_PASSWORD not found")
	}

	name, exists := os.LookupEnv("DB_NAME")
	if !exists {
		return nil, errors.New("DB_NAME not found")
	}

	sslMode, exists := os.LookupEnv("DB_SSLMODE")
	if !exists {
		return nil, errors.New("DB_SSLMODE not found")
	}

	return &DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
		SSLMode:  sslMode,
	}, nil
}
