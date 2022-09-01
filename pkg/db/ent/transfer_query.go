// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/transfer"
	"github.com/google/uuid"
)

// TransferQuery is the builder for querying Transfer entities.
type TransferQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Transfer
	modifiers  []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TransferQuery builder.
func (tq *TransferQuery) Where(ps ...predicate.Transfer) *TransferQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TransferQuery) Limit(limit int) *TransferQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TransferQuery) Offset(offset int) *TransferQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TransferQuery) Unique(unique bool) *TransferQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *TransferQuery) Order(o ...OrderFunc) *TransferQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// First returns the first Transfer entity from the query.
// Returns a *NotFoundError when no Transfer was found.
func (tq *TransferQuery) First(ctx context.Context) (*Transfer, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{transfer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TransferQuery) FirstX(ctx context.Context) *Transfer {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Transfer ID from the query.
// Returns a *NotFoundError when no Transfer ID was found.
func (tq *TransferQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{transfer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TransferQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Transfer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Transfer entity is found.
// Returns a *NotFoundError when no Transfer entities are found.
func (tq *TransferQuery) Only(ctx context.Context) (*Transfer, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{transfer.Label}
	default:
		return nil, &NotSingularError{transfer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TransferQuery) OnlyX(ctx context.Context) *Transfer {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Transfer ID in the query.
// Returns a *NotSingularError when more than one Transfer ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TransferQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{transfer.Label}
	default:
		err = &NotSingularError{transfer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TransferQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Transfers.
func (tq *TransferQuery) All(ctx context.Context) ([]*Transfer, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TransferQuery) AllX(ctx context.Context) []*Transfer {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Transfer IDs.
func (tq *TransferQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := tq.Select(transfer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TransferQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TransferQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TransferQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TransferQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TransferQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TransferQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TransferQuery) Clone() *TransferQuery {
	if tq == nil {
		return nil
	}
	return &TransferQuery{
		config:     tq.config,
		limit:      tq.limit,
		offset:     tq.offset,
		order:      append([]OrderFunc{}, tq.order...),
		predicates: append([]predicate.Transfer{}, tq.predicates...),
		// clone intermediate query.
		sql:    tq.sql.Clone(),
		path:   tq.path,
		unique: tq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Transfer.Query().
//		GroupBy(transfer.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *TransferQuery) GroupBy(field string, fields ...string) *TransferGroupBy {
	group := &TransferGroupBy{config: tq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.Transfer.Query().
//		Select(transfer.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (tq *TransferQuery) Select(fields ...string) *TransferSelect {
	tq.fields = append(tq.fields, fields...)
	return &TransferSelect{TransferQuery: tq}
}

func (tq *TransferQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !transfer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	if transfer.Policy == nil {
		return errors.New("ent: uninitialized transfer.Policy (forgotten import ent/runtime?)")
	}
	if err := transfer.Policy.EvalQuery(ctx, tq); err != nil {
		return err
	}
	return nil
}

func (tq *TransferQuery) sqlAll(ctx context.Context) ([]*Transfer, error) {
	var (
		nodes = []*Transfer{}
		_spec = tq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Transfer{config: tq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tq *TransferQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.fields
	if len(tq.fields) > 0 {
		_spec.Unique = tq.unique != nil && *tq.unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TransferQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tq *TransferQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transfer.Table,
			Columns: transfer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: transfer.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transfer.FieldID)
		for i := range fields {
			if fields[i] != transfer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TransferQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(transfer.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = transfer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.unique != nil && *tq.unique {
		selector.Distinct()
	}
	for _, m := range tq.modifiers {
		m(selector)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tq *TransferQuery) ForUpdate(opts ...sql.LockOption) *TransferQuery {
	if tq.driver.Dialect() == dialect.Postgres {
		tq.Unique(false)
	}
	tq.modifiers = append(tq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tq *TransferQuery) ForShare(opts ...sql.LockOption) *TransferQuery {
	if tq.driver.Dialect() == dialect.Postgres {
		tq.Unique(false)
	}
	tq.modifiers = append(tq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tq *TransferQuery) Modify(modifiers ...func(s *sql.Selector)) *TransferSelect {
	tq.modifiers = append(tq.modifiers, modifiers...)
	return tq.Select()
}

// TransferGroupBy is the group-by builder for Transfer entities.
type TransferGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TransferGroupBy) Aggregate(fns ...AggregateFunc) *TransferGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *TransferGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tgb *TransferGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransferGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransferGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tgb *TransferGroupBy) StringsX(ctx context.Context) []string {
	v, err := tgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransferGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfer.Label}
	default:
		err = fmt.Errorf("ent: TransferGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tgb *TransferGroupBy) StringX(ctx context.Context) string {
	v, err := tgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransferGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransferGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tgb *TransferGroupBy) IntsX(ctx context.Context) []int {
	v, err := tgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransferGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfer.Label}
	default:
		err = fmt.Errorf("ent: TransferGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tgb *TransferGroupBy) IntX(ctx context.Context) int {
	v, err := tgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransferGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransferGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tgb *TransferGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransferGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfer.Label}
	default:
		err = fmt.Errorf("ent: TransferGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tgb *TransferGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransferGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransferGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tgb *TransferGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransferGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfer.Label}
	default:
		err = fmt.Errorf("ent: TransferGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tgb *TransferGroupBy) BoolX(ctx context.Context) bool {
	v, err := tgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tgb *TransferGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tgb.fields {
		if !transfer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TransferGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TransferSelect is the builder for selecting fields of Transfer entities.
type TransferSelect struct {
	*TransferQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TransferSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.TransferQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ts *TransferSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ts *TransferSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransferSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ts *TransferSelect) StringsX(ctx context.Context) []string {
	v, err := ts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ts *TransferSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfer.Label}
	default:
		err = fmt.Errorf("ent: TransferSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ts *TransferSelect) StringX(ctx context.Context) string {
	v, err := ts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ts *TransferSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransferSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ts *TransferSelect) IntsX(ctx context.Context) []int {
	v, err := ts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ts *TransferSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfer.Label}
	default:
		err = fmt.Errorf("ent: TransferSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ts *TransferSelect) IntX(ctx context.Context) int {
	v, err := ts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ts *TransferSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransferSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ts *TransferSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ts *TransferSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfer.Label}
	default:
		err = fmt.Errorf("ent: TransferSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ts *TransferSelect) Float64X(ctx context.Context) float64 {
	v, err := ts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ts *TransferSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransferSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ts *TransferSelect) BoolsX(ctx context.Context) []bool {
	v, err := ts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ts *TransferSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfer.Label}
	default:
		err = fmt.Errorf("ent: TransferSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ts *TransferSelect) BoolX(ctx context.Context) bool {
	v, err := ts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ts *TransferSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ts *TransferSelect) Modify(modifiers ...func(s *sql.Selector)) *TransferSelect {
	ts.modifiers = append(ts.modifiers, modifiers...)
	return ts
}
