package main

import (
	"SoftwareGoDay3/routes"
	"SoftwareGoDay3/server"
)

func main() {
  server := server.NewServer()
  routes.ApplyRoutes(server.Router)
  server.Router.Run("localhost:8080")
}