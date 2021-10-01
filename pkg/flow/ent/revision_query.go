// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/vorteil/direktiv/pkg/flow/ent/instance"
	"github.com/vorteil/direktiv/pkg/flow/ent/predicate"
	"github.com/vorteil/direktiv/pkg/flow/ent/ref"
	"github.com/vorteil/direktiv/pkg/flow/ent/revision"
	"github.com/vorteil/direktiv/pkg/flow/ent/workflow"
)

// RevisionQuery is the builder for querying Revision entities.
type RevisionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Revision
	// eager-loading edges.
	withWorkflow  *WorkflowQuery
	withRefs      *RefQuery
	withInstances *InstanceQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RevisionQuery builder.
func (rq *RevisionQuery) Where(ps ...predicate.Revision) *RevisionQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *RevisionQuery) Limit(limit int) *RevisionQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *RevisionQuery) Offset(offset int) *RevisionQuery {
	rq.offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RevisionQuery) Unique(unique bool) *RevisionQuery {
	rq.unique = &unique
	return rq
}

// Order adds an order step to the query.
func (rq *RevisionQuery) Order(o ...OrderFunc) *RevisionQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryWorkflow chains the current query on the "workflow" edge.
func (rq *RevisionQuery) QueryWorkflow() *WorkflowQuery {
	query := &WorkflowQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(revision.Table, revision.FieldID, selector),
			sqlgraph.To(workflow.Table, workflow.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, revision.WorkflowTable, revision.WorkflowColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRefs chains the current query on the "refs" edge.
func (rq *RevisionQuery) QueryRefs() *RefQuery {
	query := &RefQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(revision.Table, revision.FieldID, selector),
			sqlgraph.To(ref.Table, ref.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, revision.RefsTable, revision.RefsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInstances chains the current query on the "instances" edge.
func (rq *RevisionQuery) QueryInstances() *InstanceQuery {
	query := &InstanceQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(revision.Table, revision.FieldID, selector),
			sqlgraph.To(instance.Table, instance.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, revision.InstancesTable, revision.InstancesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Revision entity from the query.
// Returns a *NotFoundError when no Revision was found.
func (rq *RevisionQuery) First(ctx context.Context) (*Revision, error) {
	nodes, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{revision.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RevisionQuery) FirstX(ctx context.Context) *Revision {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Revision ID from the query.
// Returns a *NotFoundError when no Revision ID was found.
func (rq *RevisionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{revision.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RevisionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Revision entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Revision entity is not found.
// Returns a *NotFoundError when no Revision entities are found.
func (rq *RevisionQuery) Only(ctx context.Context) (*Revision, error) {
	nodes, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{revision.Label}
	default:
		return nil, &NotSingularError{revision.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RevisionQuery) OnlyX(ctx context.Context) *Revision {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Revision ID in the query.
// Returns a *NotSingularError when exactly one Revision ID is not found.
// Returns a *NotFoundError when no entities are found.
func (rq *RevisionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{revision.Label}
	default:
		err = &NotSingularError{revision.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RevisionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Revisions.
func (rq *RevisionQuery) All(ctx context.Context) ([]*Revision, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *RevisionQuery) AllX(ctx context.Context) []*Revision {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Revision IDs.
func (rq *RevisionQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := rq.Select(revision.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RevisionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RevisionQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RevisionQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RevisionQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RevisionQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RevisionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RevisionQuery) Clone() *RevisionQuery {
	if rq == nil {
		return nil
	}
	return &RevisionQuery{
		config:        rq.config,
		limit:         rq.limit,
		offset:        rq.offset,
		order:         append([]OrderFunc{}, rq.order...),
		predicates:    append([]predicate.Revision{}, rq.predicates...),
		withWorkflow:  rq.withWorkflow.Clone(),
		withRefs:      rq.withRefs.Clone(),
		withInstances: rq.withInstances.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithWorkflow tells the query-builder to eager-load the nodes that are connected to
// the "workflow" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RevisionQuery) WithWorkflow(opts ...func(*WorkflowQuery)) *RevisionQuery {
	query := &WorkflowQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withWorkflow = query
	return rq
}

// WithRefs tells the query-builder to eager-load the nodes that are connected to
// the "refs" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RevisionQuery) WithRefs(opts ...func(*RefQuery)) *RevisionQuery {
	query := &RefQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withRefs = query
	return rq
}

// WithInstances tells the query-builder to eager-load the nodes that are connected to
// the "instances" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RevisionQuery) WithInstances(opts ...func(*InstanceQuery)) *RevisionQuery {
	query := &InstanceQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withInstances = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Revision.Query().
//		GroupBy(revision.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rq *RevisionQuery) GroupBy(field string, fields ...string) *RevisionGroupBy {
	group := &RevisionGroupBy{config: rq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Revision.Query().
//		Select(revision.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (rq *RevisionQuery) Select(fields ...string) *RevisionSelect {
	rq.fields = append(rq.fields, fields...)
	return &RevisionSelect{RevisionQuery: rq}
}

func (rq *RevisionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rq.fields {
		if !revision.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RevisionQuery) sqlAll(ctx context.Context) ([]*Revision, error) {
	var (
		nodes       = []*Revision{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [3]bool{
			rq.withWorkflow != nil,
			rq.withRefs != nil,
			rq.withInstances != nil,
		}
	)
	if rq.withWorkflow != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, revision.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Revision{config: rq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rq.withWorkflow; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Revision)
		for i := range nodes {
			if nodes[i].workflow_revisions == nil {
				continue
			}
			fk := *nodes[i].workflow_revisions
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(workflow.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workflow_revisions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Workflow = n
			}
		}
	}

	if query := rq.withRefs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Revision)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Refs = []*Ref{}
		}
		query.withFKs = true
		query.Where(predicate.Ref(func(s *sql.Selector) {
			s.Where(sql.InValues(revision.RefsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.revision_refs
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "revision_refs" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "revision_refs" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Refs = append(node.Edges.Refs, n)
		}
	}

	if query := rq.withInstances; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Revision)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Instances = []*Instance{}
		}
		query.withFKs = true
		query.Where(predicate.Instance(func(s *sql.Selector) {
			s.Where(sql.InValues(revision.InstancesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.revision_instances
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "revision_instances" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "revision_instances" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Instances = append(node.Edges.Instances, n)
		}
	}

	return nodes, nil
}

func (rq *RevisionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RevisionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (rq *RevisionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   revision.Table,
			Columns: revision.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: revision.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if unique := rq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, revision.FieldID)
		for i := range fields {
			if fields[i] != revision.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RevisionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(revision.Table)
	columns := rq.fields
	if len(columns) == 0 {
		columns = revision.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RevisionGroupBy is the group-by builder for Revision entities.
type RevisionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RevisionGroupBy) Aggregate(fns ...AggregateFunc) *RevisionGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rgb *RevisionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rgb *RevisionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *RevisionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RevisionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rgb *RevisionGroupBy) StringsX(ctx context.Context) []string {
	v, err := rgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *RevisionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{revision.Label}
	default:
		err = fmt.Errorf("ent: RevisionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rgb *RevisionGroupBy) StringX(ctx context.Context) string {
	v, err := rgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *RevisionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RevisionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rgb *RevisionGroupBy) IntsX(ctx context.Context) []int {
	v, err := rgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *RevisionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{revision.Label}
	default:
		err = fmt.Errorf("ent: RevisionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rgb *RevisionGroupBy) IntX(ctx context.Context) int {
	v, err := rgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *RevisionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RevisionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rgb *RevisionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *RevisionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{revision.Label}
	default:
		err = fmt.Errorf("ent: RevisionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rgb *RevisionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *RevisionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RevisionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rgb *RevisionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *RevisionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{revision.Label}
	default:
		err = fmt.Errorf("ent: RevisionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rgb *RevisionGroupBy) BoolX(ctx context.Context) bool {
	v, err := rgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rgb *RevisionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rgb.fields {
		if !revision.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *RevisionGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql.Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
		for _, f := range rgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rgb.fields...)...)
}

// RevisionSelect is the builder for selecting fields of Revision entities.
type RevisionSelect struct {
	*RevisionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RevisionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	rs.sql = rs.RevisionQuery.sqlQuery(ctx)
	return rs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rs *RevisionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (rs *RevisionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RevisionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rs *RevisionSelect) StringsX(ctx context.Context) []string {
	v, err := rs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (rs *RevisionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{revision.Label}
	default:
		err = fmt.Errorf("ent: RevisionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rs *RevisionSelect) StringX(ctx context.Context) string {
	v, err := rs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (rs *RevisionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RevisionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rs *RevisionSelect) IntsX(ctx context.Context) []int {
	v, err := rs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (rs *RevisionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{revision.Label}
	default:
		err = fmt.Errorf("ent: RevisionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rs *RevisionSelect) IntX(ctx context.Context) int {
	v, err := rs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (rs *RevisionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RevisionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rs *RevisionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (rs *RevisionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{revision.Label}
	default:
		err = fmt.Errorf("ent: RevisionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rs *RevisionSelect) Float64X(ctx context.Context) float64 {
	v, err := rs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (rs *RevisionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RevisionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rs *RevisionSelect) BoolsX(ctx context.Context) []bool {
	v, err := rs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (rs *RevisionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{revision.Label}
	default:
		err = fmt.Errorf("ent: RevisionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rs *RevisionSelect) BoolX(ctx context.Context) bool {
	v, err := rs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rs *RevisionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rs.sql.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
