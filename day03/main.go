package main

import (
	"SoftwareGoDay3/routes"
	"SoftwareGoDay3/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {
  godotenv.Load(".env")
  s := server.NewServer()
  routes.ApplyRoutes(s.Router)
  s.Router.Run(os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT"))
}