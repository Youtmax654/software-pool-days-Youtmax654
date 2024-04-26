package database

import (
	"SoftwareGoDay2/ent"
	"SoftwareGoDay2/ent/migrate"
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	Client *ent.Client
}

func NewDatabase(databaseUrl string) (*Database, error) {
	// Open a connection to the database.
	db, err := sql.Open("pgx", databaseUrl)
  if err != nil {
    return nil, err
  }

  // Create an ent.Driver from `db`.
  drv := entsql.OpenDB(dialect.Postgres, db)
	// Create an ent.Client from the ent.Driver.
	client := ent.NewClient(ent.Driver(drv))
	
	// Create the table from the schema with the
	// "DropIndex" and "DropColumn" options set to true.
	if err := client.Schema.Create(
		context.Background(),
		// Drop the existing tables and index
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		); err != nil {
		return nil, err
	}

	return &Database{Client: client}, nil
}