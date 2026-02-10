package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	mongoURI   string
	mongoDB    string
	serverPort string
}

func Load() (config, error) {

	// godotenv.load reads .env and sets them into process env
	// os.getenv -> read those values

	if err := godotenv.Load(); err != nil {
		return config{}, fmt.Errorf("filed to load .env")
	}

	mongoURI, err := extactEnv("MONGO_URI")

	if err != nil {
		return config{}, err
	}

	mongoDB, err := extactEnv("MONGO_NAME")

	if err != nil {
		return config{}, err
	}
	port, err := extactEnv("PORT")

	if err != nil {
		return config{}, err
	}

	return config{
		mongoURI:   mongoURI,
		mongoDB:    mongoDB,
		serverPort: port,
	}, nil

}

func extactEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("missing req env")
	}
	return val, nil

}
