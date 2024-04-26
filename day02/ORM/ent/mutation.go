// Code generated by ent, DO NOT EDIT.

package ent

import (
	"SoftwareGoDay2/ent/artist"
	"SoftwareGoDay2/ent/contact"
	"SoftwareGoDay2/ent/predicate"
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeArtist  = "Artist"
	TypeContact = "Contact"
)

// ArtistMutation represents an operation that mutates the Artist nodes in the graph.
type ArtistMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	name           *string
	nationality    *string
	clearedFields  map[string]struct{}
	contact        *uuid.UUID
	clearedcontact bool
	done           bool
	oldValue       func(context.Context) (*Artist, error)
	predicates     []predicate.Artist
}

var _ ent.Mutation = (*ArtistMutation)(nil)

// artistOption allows management of the mutation configuration using functional options.
type artistOption func(*ArtistMutation)

// newArtistMutation creates new mutation for the Artist entity.
func newArtistMutation(c config, op Op, opts ...artistOption) *ArtistMutation {
	m := &ArtistMutation{
		config:        c,
		op:            op,
		typ:           TypeArtist,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withArtistID sets the ID field of the mutation.
func withArtistID(id uuid.UUID) artistOption {
	return func(m *ArtistMutation) {
		var (
			err   error
			once  sync.Once
			value *Artist
		)
		m.oldValue = func(ctx context.Context) (*Artist, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Artist.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withArtist sets the old Artist of the mutation.
func withArtist(node *Artist) artistOption {
	return func(m *ArtistMutation) {
		m.oldValue = func(context.Context) (*Artist, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ArtistMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ArtistMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Artist entities.
func (m *ArtistMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ArtistMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ArtistMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Artist.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *ArtistMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ArtistMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Artist entity.
// If the Artist object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArtistMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ArtistMutation) ResetName() {
	m.name = nil
}

// SetNationality sets the "nationality" field.
func (m *ArtistMutation) SetNationality(s string) {
	m.nationality = &s
}

// Nationality returns the value of the "nationality" field in the mutation.
func (m *ArtistMutation) Nationality() (r string, exists bool) {
	v := m.nationality
	if v == nil {
		return
	}
	return *v, true
}

// OldNationality returns the old "nationality" field's value of the Artist entity.
// If the Artist object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArtistMutation) OldNationality(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNationality is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNationality requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNationality: %w", err)
	}
	return oldValue.Nationality, nil
}

// ResetNationality resets all changes to the "nationality" field.
func (m *ArtistMutation) ResetNationality() {
	m.nationality = nil
}

// SetContactID sets the "contact" edge to the Contact entity by id.
func (m *ArtistMutation) SetContactID(id uuid.UUID) {
	m.contact = &id
}

// ClearContact clears the "contact" edge to the Contact entity.
func (m *ArtistMutation) ClearContact() {
	m.clearedcontact = true
}

// ContactCleared reports if the "contact" edge to the Contact entity was cleared.
func (m *ArtistMutation) ContactCleared() bool {
	return m.clearedcontact
}

// ContactID returns the "contact" edge ID in the mutation.
func (m *ArtistMutation) ContactID() (id uuid.UUID, exists bool) {
	if m.contact != nil {
		return *m.contact, true
	}
	return
}

// ContactIDs returns the "contact" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ContactID instead. It exists only for internal usage by the builders.
func (m *ArtistMutation) ContactIDs() (ids []uuid.UUID) {
	if id := m.contact; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetContact resets all changes to the "contact" edge.
func (m *ArtistMutation) ResetContact() {
	m.contact = nil
	m.clearedcontact = false
}

// Where appends a list predicates to the ArtistMutation builder.
func (m *ArtistMutation) Where(ps ...predicate.Artist) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ArtistMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ArtistMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Artist, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ArtistMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ArtistMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Artist).
func (m *ArtistMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ArtistMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, artist.FieldName)
	}
	if m.nationality != nil {
		fields = append(fields, artist.FieldNationality)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ArtistMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case artist.FieldName:
		return m.Name()
	case artist.FieldNationality:
		return m.Nationality()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ArtistMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case artist.FieldName:
		return m.OldName(ctx)
	case artist.FieldNationality:
		return m.OldNationality(ctx)
	}
	return nil, fmt.Errorf("unknown Artist field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArtistMutation) SetField(name string, value ent.Value) error {
	switch name {
	case artist.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case artist.FieldNationality:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNationality(v)
		return nil
	}
	return fmt.Errorf("unknown Artist field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ArtistMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ArtistMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArtistMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Artist numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ArtistMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ArtistMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ArtistMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Artist nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ArtistMutation) ResetField(name string) error {
	switch name {
	case artist.FieldName:
		m.ResetName()
		return nil
	case artist.FieldNationality:
		m.ResetNationality()
		return nil
	}
	return fmt.Errorf("unknown Artist field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ArtistMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.contact != nil {
		edges = append(edges, artist.EdgeContact)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ArtistMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case artist.EdgeContact:
		if id := m.contact; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ArtistMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ArtistMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ArtistMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcontact {
		edges = append(edges, artist.EdgeContact)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ArtistMutation) EdgeCleared(name string) bool {
	switch name {
	case artist.EdgeContact:
		return m.clearedcontact
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ArtistMutation) ClearEdge(name string) error {
	switch name {
	case artist.EdgeContact:
		m.ClearContact()
		return nil
	}
	return fmt.Errorf("unknown Artist unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ArtistMutation) ResetEdge(name string) error {
	switch name {
	case artist.EdgeContact:
		m.ResetContact()
		return nil
	}
	return fmt.Errorf("unknown Artist edge %s", name)
}

// ContactMutation represents an operation that mutates the Contact nodes in the graph.
type ContactMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	email         *string
	phone         *string
	clearedFields map[string]struct{}
	artist        *uuid.UUID
	clearedartist bool
	done          bool
	oldValue      func(context.Context) (*Contact, error)
	predicates    []predicate.Contact
}

var _ ent.Mutation = (*ContactMutation)(nil)

// contactOption allows management of the mutation configuration using functional options.
type contactOption func(*ContactMutation)

// newContactMutation creates new mutation for the Contact entity.
func newContactMutation(c config, op Op, opts ...contactOption) *ContactMutation {
	m := &ContactMutation{
		config:        c,
		op:            op,
		typ:           TypeContact,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withContactID sets the ID field of the mutation.
func withContactID(id uuid.UUID) contactOption {
	return func(m *ContactMutation) {
		var (
			err   error
			once  sync.Once
			value *Contact
		)
		m.oldValue = func(ctx context.Context) (*Contact, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Contact.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withContact sets the old Contact of the mutation.
func withContact(node *Contact) contactOption {
	return func(m *ContactMutation) {
		m.oldValue = func(context.Context) (*Contact, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ContactMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ContactMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Contact entities.
func (m *ContactMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ContactMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ContactMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Contact.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEmail sets the "email" field.
func (m *ContactMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *ContactMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Contact entity.
// If the Contact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContactMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *ContactMutation) ResetEmail() {
	m.email = nil
}

// SetPhone sets the "phone" field.
func (m *ContactMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the value of the "phone" field in the mutation.
func (m *ContactMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// OldPhone returns the old "phone" field's value of the Contact entity.
// If the Contact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContactMutation) OldPhone(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPhone is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPhone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhone: %w", err)
	}
	return oldValue.Phone, nil
}

// ResetPhone resets all changes to the "phone" field.
func (m *ContactMutation) ResetPhone() {
	m.phone = nil
}

// SetArtistID sets the "artist" edge to the Artist entity by id.
func (m *ContactMutation) SetArtistID(id uuid.UUID) {
	m.artist = &id
}

// ClearArtist clears the "artist" edge to the Artist entity.
func (m *ContactMutation) ClearArtist() {
	m.clearedartist = true
}

// ArtistCleared reports if the "artist" edge to the Artist entity was cleared.
func (m *ContactMutation) ArtistCleared() bool {
	return m.clearedartist
}

// ArtistID returns the "artist" edge ID in the mutation.
func (m *ContactMutation) ArtistID() (id uuid.UUID, exists bool) {
	if m.artist != nil {
		return *m.artist, true
	}
	return
}

// ArtistIDs returns the "artist" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ArtistID instead. It exists only for internal usage by the builders.
func (m *ContactMutation) ArtistIDs() (ids []uuid.UUID) {
	if id := m.artist; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetArtist resets all changes to the "artist" edge.
func (m *ContactMutation) ResetArtist() {
	m.artist = nil
	m.clearedartist = false
}

// Where appends a list predicates to the ContactMutation builder.
func (m *ContactMutation) Where(ps ...predicate.Contact) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ContactMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ContactMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Contact, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ContactMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ContactMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Contact).
func (m *ContactMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ContactMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.email != nil {
		fields = append(fields, contact.FieldEmail)
	}
	if m.phone != nil {
		fields = append(fields, contact.FieldPhone)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ContactMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case contact.FieldEmail:
		return m.Email()
	case contact.FieldPhone:
		return m.Phone()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ContactMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case contact.FieldEmail:
		return m.OldEmail(ctx)
	case contact.FieldPhone:
		return m.OldPhone(ctx)
	}
	return nil, fmt.Errorf("unknown Contact field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ContactMutation) SetField(name string, value ent.Value) error {
	switch name {
	case contact.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case contact.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	}
	return fmt.Errorf("unknown Contact field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ContactMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ContactMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ContactMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Contact numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ContactMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ContactMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ContactMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Contact nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ContactMutation) ResetField(name string) error {
	switch name {
	case contact.FieldEmail:
		m.ResetEmail()
		return nil
	case contact.FieldPhone:
		m.ResetPhone()
		return nil
	}
	return fmt.Errorf("unknown Contact field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ContactMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.artist != nil {
		edges = append(edges, contact.EdgeArtist)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ContactMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case contact.EdgeArtist:
		if id := m.artist; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ContactMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ContactMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ContactMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedartist {
		edges = append(edges, contact.EdgeArtist)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ContactMutation) EdgeCleared(name string) bool {
	switch name {
	case contact.EdgeArtist:
		return m.clearedartist
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ContactMutation) ClearEdge(name string) error {
	switch name {
	case contact.EdgeArtist:
		m.ClearArtist()
		return nil
	}
	return fmt.Errorf("unknown Contact unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ContactMutation) ResetEdge(name string) error {
	switch name {
	case contact.EdgeArtist:
		m.ResetArtist()
		return nil
	}
	return fmt.Errorf("unknown Contact edge %s", name)
}
