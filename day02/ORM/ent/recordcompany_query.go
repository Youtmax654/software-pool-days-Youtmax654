// Code generated by ent, DO NOT EDIT.

package ent

import (
	"SoftwareGoDay2/ent/artist"
	"SoftwareGoDay2/ent/predicate"
	"SoftwareGoDay2/ent/recordcompany"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RecordCompanyQuery is the builder for querying RecordCompany entities.
type RecordCompanyQuery struct {
	config
	ctx         *QueryContext
	order       []recordcompany.OrderOption
	inters      []Interceptor
	predicates  []predicate.RecordCompany
	withArtists *ArtistQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RecordCompanyQuery builder.
func (rcq *RecordCompanyQuery) Where(ps ...predicate.RecordCompany) *RecordCompanyQuery {
	rcq.predicates = append(rcq.predicates, ps...)
	return rcq
}

// Limit the number of records to be returned by this query.
func (rcq *RecordCompanyQuery) Limit(limit int) *RecordCompanyQuery {
	rcq.ctx.Limit = &limit
	return rcq
}

// Offset to start from.
func (rcq *RecordCompanyQuery) Offset(offset int) *RecordCompanyQuery {
	rcq.ctx.Offset = &offset
	return rcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rcq *RecordCompanyQuery) Unique(unique bool) *RecordCompanyQuery {
	rcq.ctx.Unique = &unique
	return rcq
}

// Order specifies how the records should be ordered.
func (rcq *RecordCompanyQuery) Order(o ...recordcompany.OrderOption) *RecordCompanyQuery {
	rcq.order = append(rcq.order, o...)
	return rcq
}

// QueryArtists chains the current query on the "artists" edge.
func (rcq *RecordCompanyQuery) QueryArtists() *ArtistQuery {
	query := (&ArtistClient{config: rcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recordcompany.Table, recordcompany.FieldID, selector),
			sqlgraph.To(artist.Table, artist.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, recordcompany.ArtistsTable, recordcompany.ArtistsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RecordCompany entity from the query.
// Returns a *NotFoundError when no RecordCompany was found.
func (rcq *RecordCompanyQuery) First(ctx context.Context) (*RecordCompany, error) {
	nodes, err := rcq.Limit(1).All(setContextOp(ctx, rcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{recordcompany.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rcq *RecordCompanyQuery) FirstX(ctx context.Context) *RecordCompany {
	node, err := rcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RecordCompany ID from the query.
// Returns a *NotFoundError when no RecordCompany ID was found.
func (rcq *RecordCompanyQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rcq.Limit(1).IDs(setContextOp(ctx, rcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{recordcompany.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rcq *RecordCompanyQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RecordCompany entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RecordCompany entity is found.
// Returns a *NotFoundError when no RecordCompany entities are found.
func (rcq *RecordCompanyQuery) Only(ctx context.Context) (*RecordCompany, error) {
	nodes, err := rcq.Limit(2).All(setContextOp(ctx, rcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{recordcompany.Label}
	default:
		return nil, &NotSingularError{recordcompany.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rcq *RecordCompanyQuery) OnlyX(ctx context.Context) *RecordCompany {
	node, err := rcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RecordCompany ID in the query.
// Returns a *NotSingularError when more than one RecordCompany ID is found.
// Returns a *NotFoundError when no entities are found.
func (rcq *RecordCompanyQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rcq.Limit(2).IDs(setContextOp(ctx, rcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{recordcompany.Label}
	default:
		err = &NotSingularError{recordcompany.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rcq *RecordCompanyQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RecordCompanies.
func (rcq *RecordCompanyQuery) All(ctx context.Context) ([]*RecordCompany, error) {
	ctx = setContextOp(ctx, rcq.ctx, "All")
	if err := rcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RecordCompany, *RecordCompanyQuery]()
	return withInterceptors[[]*RecordCompany](ctx, rcq, qr, rcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rcq *RecordCompanyQuery) AllX(ctx context.Context) []*RecordCompany {
	nodes, err := rcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RecordCompany IDs.
func (rcq *RecordCompanyQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if rcq.ctx.Unique == nil && rcq.path != nil {
		rcq.Unique(true)
	}
	ctx = setContextOp(ctx, rcq.ctx, "IDs")
	if err = rcq.Select(recordcompany.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rcq *RecordCompanyQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rcq *RecordCompanyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rcq.ctx, "Count")
	if err := rcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rcq, querierCount[*RecordCompanyQuery](), rcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rcq *RecordCompanyQuery) CountX(ctx context.Context) int {
	count, err := rcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rcq *RecordCompanyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rcq.ctx, "Exist")
	switch _, err := rcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rcq *RecordCompanyQuery) ExistX(ctx context.Context) bool {
	exist, err := rcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RecordCompanyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rcq *RecordCompanyQuery) Clone() *RecordCompanyQuery {
	if rcq == nil {
		return nil
	}
	return &RecordCompanyQuery{
		config:      rcq.config,
		ctx:         rcq.ctx.Clone(),
		order:       append([]recordcompany.OrderOption{}, rcq.order...),
		inters:      append([]Interceptor{}, rcq.inters...),
		predicates:  append([]predicate.RecordCompany{}, rcq.predicates...),
		withArtists: rcq.withArtists.Clone(),
		// clone intermediate query.
		sql:  rcq.sql.Clone(),
		path: rcq.path,
	}
}

// WithArtists tells the query-builder to eager-load the nodes that are connected to
// the "artists" edge. The optional arguments are used to configure the query builder of the edge.
func (rcq *RecordCompanyQuery) WithArtists(opts ...func(*ArtistQuery)) *RecordCompanyQuery {
	query := (&ArtistClient{config: rcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rcq.withArtists = query
	return rcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RecordCompany.Query().
//		GroupBy(recordcompany.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rcq *RecordCompanyQuery) GroupBy(field string, fields ...string) *RecordCompanyGroupBy {
	rcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RecordCompanyGroupBy{build: rcq}
	grbuild.flds = &rcq.ctx.Fields
	grbuild.label = recordcompany.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.RecordCompany.Query().
//		Select(recordcompany.FieldName).
//		Scan(ctx, &v)
func (rcq *RecordCompanyQuery) Select(fields ...string) *RecordCompanySelect {
	rcq.ctx.Fields = append(rcq.ctx.Fields, fields...)
	sbuild := &RecordCompanySelect{RecordCompanyQuery: rcq}
	sbuild.label = recordcompany.Label
	sbuild.flds, sbuild.scan = &rcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RecordCompanySelect configured with the given aggregations.
func (rcq *RecordCompanyQuery) Aggregate(fns ...AggregateFunc) *RecordCompanySelect {
	return rcq.Select().Aggregate(fns...)
}

func (rcq *RecordCompanyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rcq); err != nil {
				return err
			}
		}
	}
	for _, f := range rcq.ctx.Fields {
		if !recordcompany.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rcq.path != nil {
		prev, err := rcq.path(ctx)
		if err != nil {
			return err
		}
		rcq.sql = prev
	}
	return nil
}

func (rcq *RecordCompanyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RecordCompany, error) {
	var (
		nodes       = []*RecordCompany{}
		_spec       = rcq.querySpec()
		loadedTypes = [1]bool{
			rcq.withArtists != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RecordCompany).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RecordCompany{config: rcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rcq.withArtists; query != nil {
		if err := rcq.loadArtists(ctx, query, nodes,
			func(n *RecordCompany) { n.Edges.Artists = []*Artist{} },
			func(n *RecordCompany, e *Artist) { n.Edges.Artists = append(n.Edges.Artists, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rcq *RecordCompanyQuery) loadArtists(ctx context.Context, query *ArtistQuery, nodes []*RecordCompany, init func(*RecordCompany), assign func(*RecordCompany, *Artist)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*RecordCompany)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(recordcompany.ArtistsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.record_company_artists
		if fk == nil {
			return fmt.Errorf(`foreign-key "record_company_artists" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "record_company_artists" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rcq *RecordCompanyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rcq.querySpec()
	_spec.Node.Columns = rcq.ctx.Fields
	if len(rcq.ctx.Fields) > 0 {
		_spec.Unique = rcq.ctx.Unique != nil && *rcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rcq.driver, _spec)
}

func (rcq *RecordCompanyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(recordcompany.Table, recordcompany.Columns, sqlgraph.NewFieldSpec(recordcompany.FieldID, field.TypeUUID))
	_spec.From = rcq.sql
	if unique := rcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rcq.path != nil {
		_spec.Unique = true
	}
	if fields := rcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recordcompany.FieldID)
		for i := range fields {
			if fields[i] != recordcompany.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rcq *RecordCompanyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rcq.driver.Dialect())
	t1 := builder.Table(recordcompany.Table)
	columns := rcq.ctx.Fields
	if len(columns) == 0 {
		columns = recordcompany.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rcq.sql != nil {
		selector = rcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rcq.ctx.Unique != nil && *rcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rcq.predicates {
		p(selector)
	}
	for _, p := range rcq.order {
		p(selector)
	}
	if offset := rcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RecordCompanyGroupBy is the group-by builder for RecordCompany entities.
type RecordCompanyGroupBy struct {
	selector
	build *RecordCompanyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rcgb *RecordCompanyGroupBy) Aggregate(fns ...AggregateFunc) *RecordCompanyGroupBy {
	rcgb.fns = append(rcgb.fns, fns...)
	return rcgb
}

// Scan applies the selector query and scans the result into the given value.
func (rcgb *RecordCompanyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rcgb.build.ctx, "GroupBy")
	if err := rcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RecordCompanyQuery, *RecordCompanyGroupBy](ctx, rcgb.build, rcgb, rcgb.build.inters, v)
}

func (rcgb *RecordCompanyGroupBy) sqlScan(ctx context.Context, root *RecordCompanyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rcgb.fns))
	for _, fn := range rcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rcgb.flds)+len(rcgb.fns))
		for _, f := range *rcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RecordCompanySelect is the builder for selecting fields of RecordCompany entities.
type RecordCompanySelect struct {
	*RecordCompanyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rcs *RecordCompanySelect) Aggregate(fns ...AggregateFunc) *RecordCompanySelect {
	rcs.fns = append(rcs.fns, fns...)
	return rcs
}

// Scan applies the selector query and scans the result into the given value.
func (rcs *RecordCompanySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rcs.ctx, "Select")
	if err := rcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RecordCompanyQuery, *RecordCompanySelect](ctx, rcs.RecordCompanyQuery, rcs, rcs.inters, v)
}

func (rcs *RecordCompanySelect) sqlScan(ctx context.Context, root *RecordCompanyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rcs.fns))
	for _, fn := range rcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
