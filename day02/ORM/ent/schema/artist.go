package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Artist holds the schema definition for the Artist entity.
type Artist struct {
	ent.Schema
}

// Fields of the Artist.
func (Artist) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.String("nationality"),
	}
}

// Edges of the Artist.
func (Artist) Edges() []ent.Edge {
	return nil
}
