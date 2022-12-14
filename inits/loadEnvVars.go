package inits

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Sucesss get env ")
	}
}
