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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcommissionsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppCommissionSettingQuery is the builder for querying AppCommissionSetting entities.
type AppCommissionSettingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppCommissionSetting
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppCommissionSettingQuery builder.
func (acsq *AppCommissionSettingQuery) Where(ps ...predicate.AppCommissionSetting) *AppCommissionSettingQuery {
	acsq.predicates = append(acsq.predicates, ps...)
	return acsq
}

// Limit adds a limit step to the query.
func (acsq *AppCommissionSettingQuery) Limit(limit int) *AppCommissionSettingQuery {
	acsq.limit = &limit
	return acsq
}

// Offset adds an offset step to the query.
func (acsq *AppCommissionSettingQuery) Offset(offset int) *AppCommissionSettingQuery {
	acsq.offset = &offset
	return acsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acsq *AppCommissionSettingQuery) Unique(unique bool) *AppCommissionSettingQuery {
	acsq.unique = &unique
	return acsq
}

// Order adds an order step to the query.
func (acsq *AppCommissionSettingQuery) Order(o ...OrderFunc) *AppCommissionSettingQuery {
	acsq.order = append(acsq.order, o...)
	return acsq
}

// First returns the first AppCommissionSetting entity from the query.
// Returns a *NotFoundError when no AppCommissionSetting was found.
func (acsq *AppCommissionSettingQuery) First(ctx context.Context) (*AppCommissionSetting, error) {
	nodes, err := acsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appcommissionsetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acsq *AppCommissionSettingQuery) FirstX(ctx context.Context) *AppCommissionSetting {
	node, err := acsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppCommissionSetting ID from the query.
// Returns a *NotFoundError when no AppCommissionSetting ID was found.
func (acsq *AppCommissionSettingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appcommissionsetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acsq *AppCommissionSettingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := acsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppCommissionSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AppCommissionSetting entity is not found.
// Returns a *NotFoundError when no AppCommissionSetting entities are found.
func (acsq *AppCommissionSettingQuery) Only(ctx context.Context) (*AppCommissionSetting, error) {
	nodes, err := acsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appcommissionsetting.Label}
	default:
		return nil, &NotSingularError{appcommissionsetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acsq *AppCommissionSettingQuery) OnlyX(ctx context.Context) *AppCommissionSetting {
	node, err := acsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppCommissionSetting ID in the query.
// Returns a *NotSingularError when exactly one AppCommissionSetting ID is not found.
// Returns a *NotFoundError when no entities are found.
func (acsq *AppCommissionSettingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appcommissionsetting.Label}
	default:
		err = &NotSingularError{appcommissionsetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acsq *AppCommissionSettingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := acsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppCommissionSettings.
func (acsq *AppCommissionSettingQuery) All(ctx context.Context) ([]*AppCommissionSetting, error) {
	if err := acsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return acsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (acsq *AppCommissionSettingQuery) AllX(ctx context.Context) []*AppCommissionSetting {
	nodes, err := acsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppCommissionSetting IDs.
func (acsq *AppCommissionSettingQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := acsq.Select(appcommissionsetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acsq *AppCommissionSettingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := acsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acsq *AppCommissionSettingQuery) Count(ctx context.Context) (int, error) {
	if err := acsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return acsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (acsq *AppCommissionSettingQuery) CountX(ctx context.Context) int {
	count, err := acsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acsq *AppCommissionSettingQuery) Exist(ctx context.Context) (bool, error) {
	if err := acsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return acsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (acsq *AppCommissionSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := acsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppCommissionSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acsq *AppCommissionSettingQuery) Clone() *AppCommissionSettingQuery {
	if acsq == nil {
		return nil
	}
	return &AppCommissionSettingQuery{
		config:     acsq.config,
		limit:      acsq.limit,
		offset:     acsq.offset,
		order:      append([]OrderFunc{}, acsq.order...),
		predicates: append([]predicate.AppCommissionSetting{}, acsq.predicates...),
		// clone intermediate query.
		sql:  acsq.sql.Clone(),
		path: acsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppCommissionSetting.Query().
//		GroupBy(appcommissionsetting.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (acsq *AppCommissionSettingQuery) GroupBy(field string, fields ...string) *AppCommissionSettingGroupBy {
	group := &AppCommissionSettingGroupBy{config: acsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := acsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return acsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//	}
//
//	client.AppCommissionSetting.Query().
//		Select(appcommissionsetting.FieldAppID).
//		Scan(ctx, &v)
//
func (acsq *AppCommissionSettingQuery) Select(fields ...string) *AppCommissionSettingSelect {
	acsq.fields = append(acsq.fields, fields...)
	return &AppCommissionSettingSelect{AppCommissionSettingQuery: acsq}
}

func (acsq *AppCommissionSettingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range acsq.fields {
		if !appcommissionsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acsq.path != nil {
		prev, err := acsq.path(ctx)
		if err != nil {
			return err
		}
		acsq.sql = prev
	}
	return nil
}

func (acsq *AppCommissionSettingQuery) sqlAll(ctx context.Context) ([]*AppCommissionSetting, error) {
	var (
		nodes = []*AppCommissionSetting{}
		_spec = acsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppCommissionSetting{config: acsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, acsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (acsq *AppCommissionSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acsq.querySpec()
	_spec.Node.Columns = acsq.fields
	if len(acsq.fields) > 0 {
		_spec.Unique = acsq.unique != nil && *acsq.unique
	}
	return sqlgraph.CountNodes(ctx, acsq.driver, _spec)
}

func (acsq *AppCommissionSettingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := acsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (acsq *AppCommissionSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appcommissionsetting.Table,
			Columns: appcommissionsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcommissionsetting.FieldID,
			},
		},
		From:   acsq.sql,
		Unique: true,
	}
	if unique := acsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := acsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appcommissionsetting.FieldID)
		for i := range fields {
			if fields[i] != appcommissionsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acsq *AppCommissionSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acsq.driver.Dialect())
	t1 := builder.Table(appcommissionsetting.Table)
	columns := acsq.fields
	if len(columns) == 0 {
		columns = appcommissionsetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acsq.sql != nil {
		selector = acsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acsq.unique != nil && *acsq.unique {
		selector.Distinct()
	}
	for _, p := range acsq.predicates {
		p(selector)
	}
	for _, p := range acsq.order {
		p(selector)
	}
	if offset := acsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppCommissionSettingGroupBy is the group-by builder for AppCommissionSetting entities.
type AppCommissionSettingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acsgb *AppCommissionSettingGroupBy) Aggregate(fns ...AggregateFunc) *AppCommissionSettingGroupBy {
	acsgb.fns = append(acsgb.fns, fns...)
	return acsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (acsgb *AppCommissionSettingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := acsgb.path(ctx)
	if err != nil {
		return err
	}
	acsgb.sql = query
	return acsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (acsgb *AppCommissionSettingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := acsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCommissionSettingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(acsgb.fields) > 1 {
		return nil, errors.New("ent: AppCommissionSettingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := acsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (acsgb *AppCommissionSettingGroupBy) StringsX(ctx context.Context) []string {
	v, err := acsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCommissionSettingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = acsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcommissionsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCommissionSettingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (acsgb *AppCommissionSettingGroupBy) StringX(ctx context.Context) string {
	v, err := acsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCommissionSettingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(acsgb.fields) > 1 {
		return nil, errors.New("ent: AppCommissionSettingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := acsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (acsgb *AppCommissionSettingGroupBy) IntsX(ctx context.Context) []int {
	v, err := acsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCommissionSettingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = acsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcommissionsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCommissionSettingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (acsgb *AppCommissionSettingGroupBy) IntX(ctx context.Context) int {
	v, err := acsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCommissionSettingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(acsgb.fields) > 1 {
		return nil, errors.New("ent: AppCommissionSettingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := acsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (acsgb *AppCommissionSettingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := acsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCommissionSettingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = acsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcommissionsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCommissionSettingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (acsgb *AppCommissionSettingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := acsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCommissionSettingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(acsgb.fields) > 1 {
		return nil, errors.New("ent: AppCommissionSettingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := acsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (acsgb *AppCommissionSettingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := acsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCommissionSettingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = acsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcommissionsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCommissionSettingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (acsgb *AppCommissionSettingGroupBy) BoolX(ctx context.Context) bool {
	v, err := acsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (acsgb *AppCommissionSettingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range acsgb.fields {
		if !appcommissionsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := acsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (acsgb *AppCommissionSettingGroupBy) sqlQuery() *sql.Selector {
	selector := acsgb.sql.Select()
	aggregation := make([]string, 0, len(acsgb.fns))
	for _, fn := range acsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(acsgb.fields)+len(acsgb.fns))
		for _, f := range acsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(acsgb.fields...)...)
}

// AppCommissionSettingSelect is the builder for selecting fields of AppCommissionSetting entities.
type AppCommissionSettingSelect struct {
	*AppCommissionSettingQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (acss *AppCommissionSettingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := acss.prepareQuery(ctx); err != nil {
		return err
	}
	acss.sql = acss.AppCommissionSettingQuery.sqlQuery(ctx)
	return acss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (acss *AppCommissionSettingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := acss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (acss *AppCommissionSettingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(acss.fields) > 1 {
		return nil, errors.New("ent: AppCommissionSettingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := acss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (acss *AppCommissionSettingSelect) StringsX(ctx context.Context) []string {
	v, err := acss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (acss *AppCommissionSettingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = acss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcommissionsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCommissionSettingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (acss *AppCommissionSettingSelect) StringX(ctx context.Context) string {
	v, err := acss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (acss *AppCommissionSettingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(acss.fields) > 1 {
		return nil, errors.New("ent: AppCommissionSettingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := acss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (acss *AppCommissionSettingSelect) IntsX(ctx context.Context) []int {
	v, err := acss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (acss *AppCommissionSettingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = acss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcommissionsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCommissionSettingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (acss *AppCommissionSettingSelect) IntX(ctx context.Context) int {
	v, err := acss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (acss *AppCommissionSettingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(acss.fields) > 1 {
		return nil, errors.New("ent: AppCommissionSettingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := acss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (acss *AppCommissionSettingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := acss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (acss *AppCommissionSettingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = acss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcommissionsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCommissionSettingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (acss *AppCommissionSettingSelect) Float64X(ctx context.Context) float64 {
	v, err := acss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (acss *AppCommissionSettingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(acss.fields) > 1 {
		return nil, errors.New("ent: AppCommissionSettingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := acss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (acss *AppCommissionSettingSelect) BoolsX(ctx context.Context) []bool {
	v, err := acss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (acss *AppCommissionSettingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = acss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcommissionsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCommissionSettingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (acss *AppCommissionSettingSelect) BoolX(ctx context.Context) bool {
	v, err := acss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (acss *AppCommissionSettingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := acss.sql.Query()
	if err := acss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}