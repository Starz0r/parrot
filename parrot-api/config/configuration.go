package config

import (
	"os"
)

type AppConfig struct {
	Port           string
	DBName         string
	DBConn         string
	AuthIssuer     string
	AuthSigningKey string
}

func IsExistingConfig() bool {
	if _, err := os.Stat("settings.ini"); os.IsNotExist(err) {
		return false
	}
	return true
}

// TODO: add config from conf file function
func FromEnv() (*AppConfig, error) {

	port, ok := os.LookupEnv("PARROT_API_PORT")
	if !ok {
		port = "9990"
	}
	dbName, ok := os.LookupEnv("PARROT_DB_NAME")
	if !ok {
		dbName = "postgres"
	}
	dbConn, ok := os.LookupEnv("PARROT_DB_CONN")
	if !ok {
		dbConn = "postgres://postgres@localhost:5432/parrot?sslmode=disable"
	}
	authIssuer, ok := os.LookupEnv("PARROT_AUTH_ISSUER")
	if !ok {
		authIssuer = "parrot@localhost"
	}
	authSigningKey, ok := os.LookupEnv("PARROT_AUTH_SIGNING_KEY")
	if !ok {
		authSigningKey = "secret"
	}
	return &AppConfig{
		Port:           port,
		DBName:         dbName,
		DBConn:         dbConn,
		AuthIssuer:     authIssuer,
		AuthSigningKey: authSigningKey,
	}, nil
}
