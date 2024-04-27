// Code generated by ent, DO NOT EDIT.

package artist

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the artist type in the database.
	Label = "artist"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNationality holds the string denoting the nationality field in the database.
	FieldNationality = "nationality"
	// EdgeContact holds the string denoting the contact edge name in mutations.
	EdgeContact = "contact"
	// EdgeRecordcompanies holds the string denoting the recordcompanies edge name in mutations.
	EdgeRecordcompanies = "recordcompanies"
	// Table holds the table name of the artist in the database.
	Table = "artists"
	// ContactTable is the table that holds the contact relation/edge.
	ContactTable = "contacts"
	// ContactInverseTable is the table name for the Contact entity.
	// It exists in this package in order to avoid circular dependency with the "contact" package.
	ContactInverseTable = "contacts"
	// ContactColumn is the table column denoting the contact relation/edge.
	ContactColumn = "artist_contact"
	// RecordcompaniesTable is the table that holds the recordcompanies relation/edge.
	RecordcompaniesTable = "artists"
	// RecordcompaniesInverseTable is the table name for the RecordCompany entity.
	// It exists in this package in order to avoid circular dependency with the "recordcompany" package.
	RecordcompaniesInverseTable = "record_companies"
	// RecordcompaniesColumn is the table column denoting the recordcompanies relation/edge.
	RecordcompaniesColumn = "record_company_artists"
)

// Columns holds all SQL columns for artist fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldNationality,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "artists"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"record_company_artists",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Artist queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByNationality orders the results by the nationality field.
func ByNationality(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNationality, opts...).ToFunc()
}

// ByContactField orders the results by contact field.
func ByContactField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContactStep(), sql.OrderByField(field, opts...))
	}
}

// ByRecordcompaniesField orders the results by recordcompanies field.
func ByRecordcompaniesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRecordcompaniesStep(), sql.OrderByField(field, opts...))
	}
}
func newContactStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContactInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ContactTable, ContactColumn),
	)
}
func newRecordcompaniesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RecordcompaniesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RecordcompaniesTable, RecordcompaniesColumn),
	)
}
