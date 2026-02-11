package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func Load() (Config, error) {

	// godotenv.load reads .env and sets them into process env
	// os.getenv -> read those values

	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("filed to load .env")
	}

	mongoURI, err := extactEnv("MONGO_URI")

	if err != nil {
		return Config{}, err
	}

	mongoDB, err := extactEnv("MONGO_NAME")

	if err != nil {
		return Config{}, err
	}
	port, err := extactEnv("PORT")

	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoURI:   mongoURI,
		MongoDB:    mongoDB,
		ServerPort: port,
	}, nil

}

func extactEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("missing req env")
	}
	return val, nil

}
