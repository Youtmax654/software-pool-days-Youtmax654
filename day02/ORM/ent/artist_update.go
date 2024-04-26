// Code generated by ent, DO NOT EDIT.

package ent

import (
	"SoftwareGoDay2/ent/artist"
	"SoftwareGoDay2/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArtistUpdate is the builder for updating Artist entities.
type ArtistUpdate struct {
	config
	hooks    []Hook
	mutation *ArtistMutation
}

// Where appends a list predicates to the ArtistUpdate builder.
func (au *ArtistUpdate) Where(ps ...predicate.Artist) *ArtistUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *ArtistUpdate) SetName(s string) *ArtistUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *ArtistUpdate) SetNillableName(s *string) *ArtistUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetNationality sets the "nationality" field.
func (au *ArtistUpdate) SetNationality(s string) *ArtistUpdate {
	au.mutation.SetNationality(s)
	return au
}

// SetNillableNationality sets the "nationality" field if the given value is not nil.
func (au *ArtistUpdate) SetNillableNationality(s *string) *ArtistUpdate {
	if s != nil {
		au.SetNationality(*s)
	}
	return au
}

// Mutation returns the ArtistMutation object of the builder.
func (au *ArtistUpdate) Mutation() *ArtistMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArtistUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArtistUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArtistUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArtistUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ArtistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(artist.Table, artist.Columns, sqlgraph.NewFieldSpec(artist.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(artist.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Nationality(); ok {
		_spec.SetField(artist.FieldNationality, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArtistUpdateOne is the builder for updating a single Artist entity.
type ArtistUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArtistMutation
}

// SetName sets the "name" field.
func (auo *ArtistUpdateOne) SetName(s string) *ArtistUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *ArtistUpdateOne) SetNillableName(s *string) *ArtistUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetNationality sets the "nationality" field.
func (auo *ArtistUpdateOne) SetNationality(s string) *ArtistUpdateOne {
	auo.mutation.SetNationality(s)
	return auo
}

// SetNillableNationality sets the "nationality" field if the given value is not nil.
func (auo *ArtistUpdateOne) SetNillableNationality(s *string) *ArtistUpdateOne {
	if s != nil {
		auo.SetNationality(*s)
	}
	return auo
}

// Mutation returns the ArtistMutation object of the builder.
func (auo *ArtistUpdateOne) Mutation() *ArtistMutation {
	return auo.mutation
}

// Where appends a list predicates to the ArtistUpdate builder.
func (auo *ArtistUpdateOne) Where(ps ...predicate.Artist) *ArtistUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArtistUpdateOne) Select(field string, fields ...string) *ArtistUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Artist entity.
func (auo *ArtistUpdateOne) Save(ctx context.Context) (*Artist, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArtistUpdateOne) SaveX(ctx context.Context) *Artist {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArtistUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArtistUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ArtistUpdateOne) sqlSave(ctx context.Context) (_node *Artist, err error) {
	_spec := sqlgraph.NewUpdateSpec(artist.Table, artist.Columns, sqlgraph.NewFieldSpec(artist.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Artist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artist.FieldID)
		for _, f := range fields {
			if !artist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != artist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(artist.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Nationality(); ok {
		_spec.SetField(artist.FieldNationality, field.TypeString, value)
	}
	_node = &Artist{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}