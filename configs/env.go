package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load() //Read and load env

	if err != nil {
		log.Fatalln("error .env")
	}

	mongoURI := os.Getenv("MONGOURI")
	return mongoURI // takes the MONGOURI and returns
}
