package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("email"),
		field.String("phone"),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("artist", Artist.Type).Unique().Ref("contact"),
	}
}
