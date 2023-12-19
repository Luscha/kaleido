// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.pitagora/pkg/ent/predicate"
	"github.pitagora/pkg/ent/procedure"
)

// ProcedureQuery is the builder for querying Procedure entities.
type ProcedureQuery struct {
	config
	ctx        *QueryContext
	order      []procedure.OrderOption
	inters     []Interceptor
	predicates []predicate.Procedure
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProcedureQuery builder.
func (pq *ProcedureQuery) Where(ps ...predicate.Procedure) *ProcedureQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *ProcedureQuery) Limit(limit int) *ProcedureQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *ProcedureQuery) Offset(offset int) *ProcedureQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProcedureQuery) Unique(unique bool) *ProcedureQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *ProcedureQuery) Order(o ...procedure.OrderOption) *ProcedureQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// First returns the first Procedure entity from the query.
// Returns a *NotFoundError when no Procedure was found.
func (pq *ProcedureQuery) First(ctx context.Context) (*Procedure, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{procedure.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProcedureQuery) FirstX(ctx context.Context) *Procedure {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Procedure ID from the query.
// Returns a *NotFoundError when no Procedure ID was found.
func (pq *ProcedureQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{procedure.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProcedureQuery) FirstIDX(ctx context.Context) int64 {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Procedure entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Procedure entity is found.
// Returns a *NotFoundError when no Procedure entities are found.
func (pq *ProcedureQuery) Only(ctx context.Context) (*Procedure, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{procedure.Label}
	default:
		return nil, &NotSingularError{procedure.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProcedureQuery) OnlyX(ctx context.Context) *Procedure {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Procedure ID in the query.
// Returns a *NotSingularError when more than one Procedure ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProcedureQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{procedure.Label}
	default:
		err = &NotSingularError{procedure.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProcedureQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Procedures.
func (pq *ProcedureQuery) All(ctx context.Context) ([]*Procedure, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Procedure, *ProcedureQuery]()
	return withInterceptors[[]*Procedure](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProcedureQuery) AllX(ctx context.Context) []*Procedure {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Procedure IDs.
func (pq *ProcedureQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(procedure.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProcedureQuery) IDsX(ctx context.Context) []int64 {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProcedureQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*ProcedureQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProcedureQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProcedureQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProcedureQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProcedureQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProcedureQuery) Clone() *ProcedureQuery {
	if pq == nil {
		return nil
	}
	return &ProcedureQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]procedure.OrderOption{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Procedure{}, pq.predicates...),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
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
//	client.Procedure.Query().
//		GroupBy(procedure.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *ProcedureQuery) GroupBy(field string, fields ...string) *ProcedureGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProcedureGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = procedure.Label
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
//	client.Procedure.Query().
//		Select(procedure.FieldName).
//		Scan(ctx, &v)
func (pq *ProcedureQuery) Select(fields ...string) *ProcedureSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &ProcedureSelect{ProcedureQuery: pq}
	sbuild.label = procedure.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProcedureSelect configured with the given aggregations.
func (pq *ProcedureQuery) Aggregate(fns ...AggregateFunc) *ProcedureSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ProcedureQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !procedure.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *ProcedureQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Procedure, error) {
	var (
		nodes = []*Procedure{}
		_spec = pq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Procedure).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Procedure{config: pq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pq *ProcedureQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProcedureQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(procedure.Table, procedure.Columns, sqlgraph.NewFieldSpec(procedure.FieldID, field.TypeInt64))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, procedure.FieldID)
		for i := range fields {
			if fields[i] != procedure.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProcedureQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(procedure.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = procedure.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProcedureGroupBy is the group-by builder for Procedure entities.
type ProcedureGroupBy struct {
	selector
	build *ProcedureQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProcedureGroupBy) Aggregate(fns ...AggregateFunc) *ProcedureGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ProcedureGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProcedureQuery, *ProcedureGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ProcedureGroupBy) sqlScan(ctx context.Context, root *ProcedureQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProcedureSelect is the builder for selecting fields of Procedure entities.
type ProcedureSelect struct {
	*ProcedureQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ProcedureSelect) Aggregate(fns ...AggregateFunc) *ProcedureSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProcedureSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProcedureQuery, *ProcedureSelect](ctx, ps.ProcedureQuery, ps, ps.inters, v)
}

func (ps *ProcedureSelect) sqlScan(ctx context.Context, root *ProcedureQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}