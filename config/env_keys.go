package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

var env ENV

const projectDirName = "echo-mongodb-api"

// InitDotEnv init params in .env file
func InitDotEnv() {
	// get env path
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	envPath := string(rootPath) + `/.env`

	// load .env file
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file")
	}

	// ...
	appPort := GetEnvString("APP_PORT")
	database := Database{URI: GetEnvString("DB_URI"), Name: GetEnvString("DB_Name")}

	// jwt
	jwt := JWT{SecretKey: GetEnvString("SECRET_KEY")}

	// ...
	env = ENV{
		AppPort:  appPort,
		Database: database,
		JWT:      jwt,
	}
}

// GetEnvString ...
func GetEnvString(key string) string {
	return os.Getenv(key)
}

// GetEnv return .env data
func GetEnv() *ENV {
	return &env
}
