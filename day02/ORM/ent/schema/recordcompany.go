package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RecordCompany holds the schema definition for the RecordCompany entity.
type RecordCompany struct {
	ent.Schema
}

// Fields of the RecordCompany.
func (RecordCompany) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
	}
}

// Edges of the RecordCompany.
func (RecordCompany) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artists", Artist.Type),
	}
}
