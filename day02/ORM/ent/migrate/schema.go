// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArtistsColumns holds the columns for the "artists" table.
	ArtistsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "nationality", Type: field.TypeString},
	}
	// ArtistsTable holds the schema information for the "artists" table.
	ArtistsTable = &schema.Table{
		Name:       "artists",
		Columns:    ArtistsColumns,
		PrimaryKey: []*schema.Column{ArtistsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArtistsTable,
	}
)

func init() {
}
