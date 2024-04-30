package server

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port     string
	Host     string
	HelloMsg string
	SecretKey string
}

func NewConfig() *Env {
	godotenv.Load("../.env")
	return &Env{
		Port: os.Getenv("SERVER_PORT"),
		Host: os.Getenv("SERVER_HOST"),
		HelloMsg: os.Getenv("HELLO_MESSAGE"),
		SecretKey: os.Getenv("JWT_SECRET"),
	}
}

var Config *Env = NewConfig()