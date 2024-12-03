package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port              string
	Dd_Addr           string
	DB_MAX_OPEN_CONNS string
	DB_MAX_IDLE_TIME  string
}

func InitEnv() *Env {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Env{
		Port:              os.Getenv("PORT"),
		Dd_Addr:           os.Getenv("DB_ADDR"),
		DB_MAX_OPEN_CONNS: os.Getenv("DB_MAX_OPEN_CONNS"),
		DB_MAX_IDLE_TIME:  os.Getenv("DB_MAX_IDLE_TIME"),
	}
}
