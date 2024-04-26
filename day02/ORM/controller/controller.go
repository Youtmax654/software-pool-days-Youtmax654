package controller

import "SoftwareGoDay2/database"

type Controller struct {
	*database. Database
	// Add some fields if necessary
}

func NewController(db *database.Database) *Controller {
	// Return a pointer to a `Controller` struct filled with the database
	return &Controller{Database: db}
}