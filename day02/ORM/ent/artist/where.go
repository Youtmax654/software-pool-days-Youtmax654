// Code generated by ent, DO NOT EDIT.

package artist

import (
	"SoftwareGoDay2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Artist {
	return predicate.Artist(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Artist {
	return predicate.Artist(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Artist {
	return predicate.Artist(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Artist {
	return predicate.Artist(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Artist {
	return predicate.Artist(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Artist {
	return predicate.Artist(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Artist {
	return predicate.Artist(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Artist {
	return predicate.Artist(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Artist {
	return predicate.Artist(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Artist {
	return predicate.Artist(sql.FieldEQ(FieldName, v))
}

// Nationality applies equality check predicate on the "nationality" field. It's identical to NationalityEQ.
func Nationality(v string) predicate.Artist {
	return predicate.Artist(sql.FieldEQ(FieldNationality, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Artist {
	return predicate.Artist(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Artist {
	return predicate.Artist(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Artist {
	return predicate.Artist(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Artist {
	return predicate.Artist(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Artist {
	return predicate.Artist(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Artist {
	return predicate.Artist(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Artist {
	return predicate.Artist(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Artist {
	return predicate.Artist(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Artist {
	return predicate.Artist(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Artist {
	return predicate.Artist(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Artist {
	return predicate.Artist(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Artist {
	return predicate.Artist(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Artist {
	return predicate.Artist(sql.FieldContainsFold(FieldName, v))
}

// NationalityEQ applies the EQ predicate on the "nationality" field.
func NationalityEQ(v string) predicate.Artist {
	return predicate.Artist(sql.FieldEQ(FieldNationality, v))
}

// NationalityNEQ applies the NEQ predicate on the "nationality" field.
func NationalityNEQ(v string) predicate.Artist {
	return predicate.Artist(sql.FieldNEQ(FieldNationality, v))
}

// NationalityIn applies the In predicate on the "nationality" field.
func NationalityIn(vs ...string) predicate.Artist {
	return predicate.Artist(sql.FieldIn(FieldNationality, vs...))
}

// NationalityNotIn applies the NotIn predicate on the "nationality" field.
func NationalityNotIn(vs ...string) predicate.Artist {
	return predicate.Artist(sql.FieldNotIn(FieldNationality, vs...))
}

// NationalityGT applies the GT predicate on the "nationality" field.
func NationalityGT(v string) predicate.Artist {
	return predicate.Artist(sql.FieldGT(FieldNationality, v))
}

// NationalityGTE applies the GTE predicate on the "nationality" field.
func NationalityGTE(v string) predicate.Artist {
	return predicate.Artist(sql.FieldGTE(FieldNationality, v))
}

// NationalityLT applies the LT predicate on the "nationality" field.
func NationalityLT(v string) predicate.Artist {
	return predicate.Artist(sql.FieldLT(FieldNationality, v))
}

// NationalityLTE applies the LTE predicate on the "nationality" field.
func NationalityLTE(v string) predicate.Artist {
	return predicate.Artist(sql.FieldLTE(FieldNationality, v))
}

// NationalityContains applies the Contains predicate on the "nationality" field.
func NationalityContains(v string) predicate.Artist {
	return predicate.Artist(sql.FieldContains(FieldNationality, v))
}

// NationalityHasPrefix applies the HasPrefix predicate on the "nationality" field.
func NationalityHasPrefix(v string) predicate.Artist {
	return predicate.Artist(sql.FieldHasPrefix(FieldNationality, v))
}

// NationalityHasSuffix applies the HasSuffix predicate on the "nationality" field.
func NationalityHasSuffix(v string) predicate.Artist {
	return predicate.Artist(sql.FieldHasSuffix(FieldNationality, v))
}

// NationalityEqualFold applies the EqualFold predicate on the "nationality" field.
func NationalityEqualFold(v string) predicate.Artist {
	return predicate.Artist(sql.FieldEqualFold(FieldNationality, v))
}

// NationalityContainsFold applies the ContainsFold predicate on the "nationality" field.
func NationalityContainsFold(v string) predicate.Artist {
	return predicate.Artist(sql.FieldContainsFold(FieldNationality, v))
}

// HasContact applies the HasEdge predicate on the "contact" edge.
func HasContact() predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ContactTable, ContactColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContactWith applies the HasEdge predicate on the "contact" edge with a given conditions (other predicates).
func HasContactWith(preds ...predicate.Contact) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		step := newContactStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Artist) predicate.Artist {
	return predicate.Artist(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Artist) predicate.Artist {
	return predicate.Artist(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Artist) predicate.Artist {
	return predicate.Artist(sql.NotPredicates(p))
}
