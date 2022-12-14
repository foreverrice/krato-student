// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"finance/internal/data/ent/predicate"
	"finance/internal/data/ent/storebalanceaccount"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StoreBalanceAccountQuery is the builder for querying StoreBalanceAccount entities.
type StoreBalanceAccountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.StoreBalanceAccount
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StoreBalanceAccountQuery builder.
func (sbaq *StoreBalanceAccountQuery) Where(ps ...predicate.StoreBalanceAccount) *StoreBalanceAccountQuery {
	sbaq.predicates = append(sbaq.predicates, ps...)
	return sbaq
}

// Limit adds a limit step to the query.
func (sbaq *StoreBalanceAccountQuery) Limit(limit int) *StoreBalanceAccountQuery {
	sbaq.limit = &limit
	return sbaq
}

// Offset adds an offset step to the query.
func (sbaq *StoreBalanceAccountQuery) Offset(offset int) *StoreBalanceAccountQuery {
	sbaq.offset = &offset
	return sbaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sbaq *StoreBalanceAccountQuery) Unique(unique bool) *StoreBalanceAccountQuery {
	sbaq.unique = &unique
	return sbaq
}

// Order adds an order step to the query.
func (sbaq *StoreBalanceAccountQuery) Order(o ...OrderFunc) *StoreBalanceAccountQuery {
	sbaq.order = append(sbaq.order, o...)
	return sbaq
}

// First returns the first StoreBalanceAccount entity from the query.
// Returns a *NotFoundError when no StoreBalanceAccount was found.
func (sbaq *StoreBalanceAccountQuery) First(ctx context.Context) (*StoreBalanceAccount, error) {
	nodes, err := sbaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{storebalanceaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sbaq *StoreBalanceAccountQuery) FirstX(ctx context.Context) *StoreBalanceAccount {
	node, err := sbaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StoreBalanceAccount ID from the query.
// Returns a *NotFoundError when no StoreBalanceAccount ID was found.
func (sbaq *StoreBalanceAccountQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = sbaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{storebalanceaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sbaq *StoreBalanceAccountQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := sbaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StoreBalanceAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StoreBalanceAccount entity is found.
// Returns a *NotFoundError when no StoreBalanceAccount entities are found.
func (sbaq *StoreBalanceAccountQuery) Only(ctx context.Context) (*StoreBalanceAccount, error) {
	nodes, err := sbaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{storebalanceaccount.Label}
	default:
		return nil, &NotSingularError{storebalanceaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sbaq *StoreBalanceAccountQuery) OnlyX(ctx context.Context) *StoreBalanceAccount {
	node, err := sbaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StoreBalanceAccount ID in the query.
// Returns a *NotSingularError when more than one StoreBalanceAccount ID is found.
// Returns a *NotFoundError when no entities are found.
func (sbaq *StoreBalanceAccountQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = sbaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{storebalanceaccount.Label}
	default:
		err = &NotSingularError{storebalanceaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sbaq *StoreBalanceAccountQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := sbaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StoreBalanceAccounts.
func (sbaq *StoreBalanceAccountQuery) All(ctx context.Context) ([]*StoreBalanceAccount, error) {
	if err := sbaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sbaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sbaq *StoreBalanceAccountQuery) AllX(ctx context.Context) []*StoreBalanceAccount {
	nodes, err := sbaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StoreBalanceAccount IDs.
func (sbaq *StoreBalanceAccountQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := sbaq.Select(storebalanceaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sbaq *StoreBalanceAccountQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := sbaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sbaq *StoreBalanceAccountQuery) Count(ctx context.Context) (int, error) {
	if err := sbaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sbaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sbaq *StoreBalanceAccountQuery) CountX(ctx context.Context) int {
	count, err := sbaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sbaq *StoreBalanceAccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := sbaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sbaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sbaq *StoreBalanceAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := sbaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StoreBalanceAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sbaq *StoreBalanceAccountQuery) Clone() *StoreBalanceAccountQuery {
	if sbaq == nil {
		return nil
	}
	return &StoreBalanceAccountQuery{
		config:     sbaq.config,
		limit:      sbaq.limit,
		offset:     sbaq.offset,
		order:      append([]OrderFunc{}, sbaq.order...),
		predicates: append([]predicate.StoreBalanceAccount{}, sbaq.predicates...),
		// clone intermediate query.
		sql:    sbaq.sql.Clone(),
		path:   sbaq.path,
		unique: sbaq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AccountNo string `json:"account_no,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StoreBalanceAccount.Query().
//		GroupBy(storebalanceaccount.FieldAccountNo).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sbaq *StoreBalanceAccountQuery) GroupBy(field string, fields ...string) *StoreBalanceAccountGroupBy {
	grbuild := &StoreBalanceAccountGroupBy{config: sbaq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sbaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sbaq.sqlQuery(ctx), nil
	}
	grbuild.label = storebalanceaccount.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AccountNo string `json:"account_no,omitempty"`
//	}
//
//	client.StoreBalanceAccount.Query().
//		Select(storebalanceaccount.FieldAccountNo).
//		Scan(ctx, &v)
//
func (sbaq *StoreBalanceAccountQuery) Select(fields ...string) *StoreBalanceAccountSelect {
	sbaq.fields = append(sbaq.fields, fields...)
	selbuild := &StoreBalanceAccountSelect{StoreBalanceAccountQuery: sbaq}
	selbuild.label = storebalanceaccount.Label
	selbuild.flds, selbuild.scan = &sbaq.fields, selbuild.Scan
	return selbuild
}

func (sbaq *StoreBalanceAccountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sbaq.fields {
		if !storebalanceaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sbaq.path != nil {
		prev, err := sbaq.path(ctx)
		if err != nil {
			return err
		}
		sbaq.sql = prev
	}
	return nil
}

func (sbaq *StoreBalanceAccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*StoreBalanceAccount, error) {
	var (
		nodes = []*StoreBalanceAccount{}
		_spec = sbaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*StoreBalanceAccount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &StoreBalanceAccount{config: sbaq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sbaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sbaq *StoreBalanceAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sbaq.querySpec()
	_spec.Node.Columns = sbaq.fields
	if len(sbaq.fields) > 0 {
		_spec.Unique = sbaq.unique != nil && *sbaq.unique
	}
	return sqlgraph.CountNodes(ctx, sbaq.driver, _spec)
}

func (sbaq *StoreBalanceAccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sbaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sbaq *StoreBalanceAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   storebalanceaccount.Table,
			Columns: storebalanceaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: storebalanceaccount.FieldID,
			},
		},
		From:   sbaq.sql,
		Unique: true,
	}
	if unique := sbaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sbaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, storebalanceaccount.FieldID)
		for i := range fields {
			if fields[i] != storebalanceaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sbaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sbaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sbaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sbaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sbaq *StoreBalanceAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sbaq.driver.Dialect())
	t1 := builder.Table(storebalanceaccount.Table)
	columns := sbaq.fields
	if len(columns) == 0 {
		columns = storebalanceaccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sbaq.sql != nil {
		selector = sbaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sbaq.unique != nil && *sbaq.unique {
		selector.Distinct()
	}
	for _, p := range sbaq.predicates {
		p(selector)
	}
	for _, p := range sbaq.order {
		p(selector)
	}
	if offset := sbaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sbaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StoreBalanceAccountGroupBy is the group-by builder for StoreBalanceAccount entities.
type StoreBalanceAccountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sbagb *StoreBalanceAccountGroupBy) Aggregate(fns ...AggregateFunc) *StoreBalanceAccountGroupBy {
	sbagb.fns = append(sbagb.fns, fns...)
	return sbagb
}

// Scan applies the group-by query and scans the result into the given value.
func (sbagb *StoreBalanceAccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sbagb.path(ctx)
	if err != nil {
		return err
	}
	sbagb.sql = query
	return sbagb.sqlScan(ctx, v)
}

func (sbagb *StoreBalanceAccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sbagb.fields {
		if !storebalanceaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sbagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sbagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sbagb *StoreBalanceAccountGroupBy) sqlQuery() *sql.Selector {
	selector := sbagb.sql.Select()
	aggregation := make([]string, 0, len(sbagb.fns))
	for _, fn := range sbagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sbagb.fields)+len(sbagb.fns))
		for _, f := range sbagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sbagb.fields...)...)
}

// StoreBalanceAccountSelect is the builder for selecting fields of StoreBalanceAccount entities.
type StoreBalanceAccountSelect struct {
	*StoreBalanceAccountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sbas *StoreBalanceAccountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sbas.prepareQuery(ctx); err != nil {
		return err
	}
	sbas.sql = sbas.StoreBalanceAccountQuery.sqlQuery(ctx)
	return sbas.sqlScan(ctx, v)
}

func (sbas *StoreBalanceAccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sbas.sql.Query()
	if err := sbas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
