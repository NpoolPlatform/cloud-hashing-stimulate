// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/newuserrewardsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
)

// NewUserRewardSettingQuery is the builder for querying NewUserRewardSetting entities.
type NewUserRewardSettingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NewUserRewardSetting
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NewUserRewardSettingQuery builder.
func (nursq *NewUserRewardSettingQuery) Where(ps ...predicate.NewUserRewardSetting) *NewUserRewardSettingQuery {
	nursq.predicates = append(nursq.predicates, ps...)
	return nursq
}

// Limit adds a limit step to the query.
func (nursq *NewUserRewardSettingQuery) Limit(limit int) *NewUserRewardSettingQuery {
	nursq.limit = &limit
	return nursq
}

// Offset adds an offset step to the query.
func (nursq *NewUserRewardSettingQuery) Offset(offset int) *NewUserRewardSettingQuery {
	nursq.offset = &offset
	return nursq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nursq *NewUserRewardSettingQuery) Unique(unique bool) *NewUserRewardSettingQuery {
	nursq.unique = &unique
	return nursq
}

// Order adds an order step to the query.
func (nursq *NewUserRewardSettingQuery) Order(o ...OrderFunc) *NewUserRewardSettingQuery {
	nursq.order = append(nursq.order, o...)
	return nursq
}

// First returns the first NewUserRewardSetting entity from the query.
// Returns a *NotFoundError when no NewUserRewardSetting was found.
func (nursq *NewUserRewardSettingQuery) First(ctx context.Context) (*NewUserRewardSetting, error) {
	nodes, err := nursq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{newuserrewardsetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nursq *NewUserRewardSettingQuery) FirstX(ctx context.Context) *NewUserRewardSetting {
	node, err := nursq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NewUserRewardSetting ID from the query.
// Returns a *NotFoundError when no NewUserRewardSetting ID was found.
func (nursq *NewUserRewardSettingQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nursq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{newuserrewardsetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nursq *NewUserRewardSettingQuery) FirstIDX(ctx context.Context) int {
	id, err := nursq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NewUserRewardSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one NewUserRewardSetting entity is not found.
// Returns a *NotFoundError when no NewUserRewardSetting entities are found.
func (nursq *NewUserRewardSettingQuery) Only(ctx context.Context) (*NewUserRewardSetting, error) {
	nodes, err := nursq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{newuserrewardsetting.Label}
	default:
		return nil, &NotSingularError{newuserrewardsetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nursq *NewUserRewardSettingQuery) OnlyX(ctx context.Context) *NewUserRewardSetting {
	node, err := nursq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NewUserRewardSetting ID in the query.
// Returns a *NotSingularError when exactly one NewUserRewardSetting ID is not found.
// Returns a *NotFoundError when no entities are found.
func (nursq *NewUserRewardSettingQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nursq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{newuserrewardsetting.Label}
	default:
		err = &NotSingularError{newuserrewardsetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nursq *NewUserRewardSettingQuery) OnlyIDX(ctx context.Context) int {
	id, err := nursq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NewUserRewardSettings.
func (nursq *NewUserRewardSettingQuery) All(ctx context.Context) ([]*NewUserRewardSetting, error) {
	if err := nursq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nursq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nursq *NewUserRewardSettingQuery) AllX(ctx context.Context) []*NewUserRewardSetting {
	nodes, err := nursq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NewUserRewardSetting IDs.
func (nursq *NewUserRewardSettingQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := nursq.Select(newuserrewardsetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nursq *NewUserRewardSettingQuery) IDsX(ctx context.Context) []int {
	ids, err := nursq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nursq *NewUserRewardSettingQuery) Count(ctx context.Context) (int, error) {
	if err := nursq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nursq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nursq *NewUserRewardSettingQuery) CountX(ctx context.Context) int {
	count, err := nursq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nursq *NewUserRewardSettingQuery) Exist(ctx context.Context) (bool, error) {
	if err := nursq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nursq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nursq *NewUserRewardSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := nursq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NewUserRewardSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nursq *NewUserRewardSettingQuery) Clone() *NewUserRewardSettingQuery {
	if nursq == nil {
		return nil
	}
	return &NewUserRewardSettingQuery{
		config:     nursq.config,
		limit:      nursq.limit,
		offset:     nursq.offset,
		order:      append([]OrderFunc{}, nursq.order...),
		predicates: append([]predicate.NewUserRewardSetting{}, nursq.predicates...),
		// clone intermediate query.
		sql:  nursq.sql.Clone(),
		path: nursq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (nursq *NewUserRewardSettingQuery) GroupBy(field string, fields ...string) *NewUserRewardSettingGroupBy {
	group := &NewUserRewardSettingGroupBy{config: nursq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nursq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nursq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (nursq *NewUserRewardSettingQuery) Select(fields ...string) *NewUserRewardSettingSelect {
	nursq.fields = append(nursq.fields, fields...)
	return &NewUserRewardSettingSelect{NewUserRewardSettingQuery: nursq}
}

func (nursq *NewUserRewardSettingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nursq.fields {
		if !newuserrewardsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nursq.path != nil {
		prev, err := nursq.path(ctx)
		if err != nil {
			return err
		}
		nursq.sql = prev
	}
	return nil
}

func (nursq *NewUserRewardSettingQuery) sqlAll(ctx context.Context) ([]*NewUserRewardSetting, error) {
	var (
		nodes = []*NewUserRewardSetting{}
		_spec = nursq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &NewUserRewardSetting{config: nursq.config}
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
	if err := sqlgraph.QueryNodes(ctx, nursq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (nursq *NewUserRewardSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nursq.querySpec()
	return sqlgraph.CountNodes(ctx, nursq.driver, _spec)
}

func (nursq *NewUserRewardSettingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nursq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (nursq *NewUserRewardSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   newuserrewardsetting.Table,
			Columns: newuserrewardsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: newuserrewardsetting.FieldID,
			},
		},
		From:   nursq.sql,
		Unique: true,
	}
	if unique := nursq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nursq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, newuserrewardsetting.FieldID)
		for i := range fields {
			if fields[i] != newuserrewardsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nursq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nursq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nursq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nursq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nursq *NewUserRewardSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nursq.driver.Dialect())
	t1 := builder.Table(newuserrewardsetting.Table)
	columns := nursq.fields
	if len(columns) == 0 {
		columns = newuserrewardsetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nursq.sql != nil {
		selector = nursq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range nursq.predicates {
		p(selector)
	}
	for _, p := range nursq.order {
		p(selector)
	}
	if offset := nursq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nursq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NewUserRewardSettingGroupBy is the group-by builder for NewUserRewardSetting entities.
type NewUserRewardSettingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nursgb *NewUserRewardSettingGroupBy) Aggregate(fns ...AggregateFunc) *NewUserRewardSettingGroupBy {
	nursgb.fns = append(nursgb.fns, fns...)
	return nursgb
}

// Scan applies the group-by query and scans the result into the given value.
func (nursgb *NewUserRewardSettingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := nursgb.path(ctx)
	if err != nil {
		return err
	}
	nursgb.sql = query
	return nursgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nursgb *NewUserRewardSettingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := nursgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (nursgb *NewUserRewardSettingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(nursgb.fields) > 1 {
		return nil, errors.New("ent: NewUserRewardSettingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := nursgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nursgb *NewUserRewardSettingGroupBy) StringsX(ctx context.Context) []string {
	v, err := nursgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nursgb *NewUserRewardSettingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nursgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newuserrewardsetting.Label}
	default:
		err = fmt.Errorf("ent: NewUserRewardSettingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nursgb *NewUserRewardSettingGroupBy) StringX(ctx context.Context) string {
	v, err := nursgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (nursgb *NewUserRewardSettingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(nursgb.fields) > 1 {
		return nil, errors.New("ent: NewUserRewardSettingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := nursgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nursgb *NewUserRewardSettingGroupBy) IntsX(ctx context.Context) []int {
	v, err := nursgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nursgb *NewUserRewardSettingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nursgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newuserrewardsetting.Label}
	default:
		err = fmt.Errorf("ent: NewUserRewardSettingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nursgb *NewUserRewardSettingGroupBy) IntX(ctx context.Context) int {
	v, err := nursgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (nursgb *NewUserRewardSettingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(nursgb.fields) > 1 {
		return nil, errors.New("ent: NewUserRewardSettingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := nursgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nursgb *NewUserRewardSettingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := nursgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nursgb *NewUserRewardSettingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nursgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newuserrewardsetting.Label}
	default:
		err = fmt.Errorf("ent: NewUserRewardSettingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nursgb *NewUserRewardSettingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := nursgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (nursgb *NewUserRewardSettingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(nursgb.fields) > 1 {
		return nil, errors.New("ent: NewUserRewardSettingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := nursgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nursgb *NewUserRewardSettingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := nursgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nursgb *NewUserRewardSettingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nursgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newuserrewardsetting.Label}
	default:
		err = fmt.Errorf("ent: NewUserRewardSettingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nursgb *NewUserRewardSettingGroupBy) BoolX(ctx context.Context) bool {
	v, err := nursgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nursgb *NewUserRewardSettingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range nursgb.fields {
		if !newuserrewardsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := nursgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nursgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (nursgb *NewUserRewardSettingGroupBy) sqlQuery() *sql.Selector {
	selector := nursgb.sql.Select()
	aggregation := make([]string, 0, len(nursgb.fns))
	for _, fn := range nursgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(nursgb.fields)+len(nursgb.fns))
		for _, f := range nursgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(nursgb.fields...)...)
}

// NewUserRewardSettingSelect is the builder for selecting fields of NewUserRewardSetting entities.
type NewUserRewardSettingSelect struct {
	*NewUserRewardSettingQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (nurss *NewUserRewardSettingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := nurss.prepareQuery(ctx); err != nil {
		return err
	}
	nurss.sql = nurss.NewUserRewardSettingQuery.sqlQuery(ctx)
	return nurss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nurss *NewUserRewardSettingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := nurss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (nurss *NewUserRewardSettingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(nurss.fields) > 1 {
		return nil, errors.New("ent: NewUserRewardSettingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := nurss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nurss *NewUserRewardSettingSelect) StringsX(ctx context.Context) []string {
	v, err := nurss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (nurss *NewUserRewardSettingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nurss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newuserrewardsetting.Label}
	default:
		err = fmt.Errorf("ent: NewUserRewardSettingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nurss *NewUserRewardSettingSelect) StringX(ctx context.Context) string {
	v, err := nurss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (nurss *NewUserRewardSettingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(nurss.fields) > 1 {
		return nil, errors.New("ent: NewUserRewardSettingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := nurss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nurss *NewUserRewardSettingSelect) IntsX(ctx context.Context) []int {
	v, err := nurss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (nurss *NewUserRewardSettingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nurss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newuserrewardsetting.Label}
	default:
		err = fmt.Errorf("ent: NewUserRewardSettingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nurss *NewUserRewardSettingSelect) IntX(ctx context.Context) int {
	v, err := nurss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (nurss *NewUserRewardSettingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(nurss.fields) > 1 {
		return nil, errors.New("ent: NewUserRewardSettingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := nurss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nurss *NewUserRewardSettingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := nurss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (nurss *NewUserRewardSettingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nurss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newuserrewardsetting.Label}
	default:
		err = fmt.Errorf("ent: NewUserRewardSettingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nurss *NewUserRewardSettingSelect) Float64X(ctx context.Context) float64 {
	v, err := nurss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (nurss *NewUserRewardSettingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(nurss.fields) > 1 {
		return nil, errors.New("ent: NewUserRewardSettingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := nurss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nurss *NewUserRewardSettingSelect) BoolsX(ctx context.Context) []bool {
	v, err := nurss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (nurss *NewUserRewardSettingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nurss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newuserrewardsetting.Label}
	default:
		err = fmt.Errorf("ent: NewUserRewardSettingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nurss *NewUserRewardSettingSelect) BoolX(ctx context.Context) bool {
	v, err := nurss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nurss *NewUserRewardSettingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := nurss.sql.Query()
	if err := nurss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
