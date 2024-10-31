package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	ConnectionString = ""
	Port             = 0
	SecretKey        []byte
)

func Load() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	SecretKey = []byte(os.Getenv("API_SECRET_KEY"))
}
