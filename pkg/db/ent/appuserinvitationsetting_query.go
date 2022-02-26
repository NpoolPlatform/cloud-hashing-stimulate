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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appuserinvitationsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserInvitationSettingQuery is the builder for querying AppUserInvitationSetting entities.
type AppUserInvitationSettingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppUserInvitationSetting
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppUserInvitationSettingQuery builder.
func (auisq *AppUserInvitationSettingQuery) Where(ps ...predicate.AppUserInvitationSetting) *AppUserInvitationSettingQuery {
	auisq.predicates = append(auisq.predicates, ps...)
	return auisq
}

// Limit adds a limit step to the query.
func (auisq *AppUserInvitationSettingQuery) Limit(limit int) *AppUserInvitationSettingQuery {
	auisq.limit = &limit
	return auisq
}

// Offset adds an offset step to the query.
func (auisq *AppUserInvitationSettingQuery) Offset(offset int) *AppUserInvitationSettingQuery {
	auisq.offset = &offset
	return auisq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (auisq *AppUserInvitationSettingQuery) Unique(unique bool) *AppUserInvitationSettingQuery {
	auisq.unique = &unique
	return auisq
}

// Order adds an order step to the query.
func (auisq *AppUserInvitationSettingQuery) Order(o ...OrderFunc) *AppUserInvitationSettingQuery {
	auisq.order = append(auisq.order, o...)
	return auisq
}

// First returns the first AppUserInvitationSetting entity from the query.
// Returns a *NotFoundError when no AppUserInvitationSetting was found.
func (auisq *AppUserInvitationSettingQuery) First(ctx context.Context) (*AppUserInvitationSetting, error) {
	nodes, err := auisq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appuserinvitationsetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (auisq *AppUserInvitationSettingQuery) FirstX(ctx context.Context) *AppUserInvitationSetting {
	node, err := auisq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppUserInvitationSetting ID from the query.
// Returns a *NotFoundError when no AppUserInvitationSetting ID was found.
func (auisq *AppUserInvitationSettingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = auisq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appuserinvitationsetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (auisq *AppUserInvitationSettingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := auisq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppUserInvitationSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AppUserInvitationSetting entity is not found.
// Returns a *NotFoundError when no AppUserInvitationSetting entities are found.
func (auisq *AppUserInvitationSettingQuery) Only(ctx context.Context) (*AppUserInvitationSetting, error) {
	nodes, err := auisq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appuserinvitationsetting.Label}
	default:
		return nil, &NotSingularError{appuserinvitationsetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (auisq *AppUserInvitationSettingQuery) OnlyX(ctx context.Context) *AppUserInvitationSetting {
	node, err := auisq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppUserInvitationSetting ID in the query.
// Returns a *NotSingularError when exactly one AppUserInvitationSetting ID is not found.
// Returns a *NotFoundError when no entities are found.
func (auisq *AppUserInvitationSettingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = auisq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appuserinvitationsetting.Label}
	default:
		err = &NotSingularError{appuserinvitationsetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (auisq *AppUserInvitationSettingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := auisq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppUserInvitationSettings.
func (auisq *AppUserInvitationSettingQuery) All(ctx context.Context) ([]*AppUserInvitationSetting, error) {
	if err := auisq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return auisq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (auisq *AppUserInvitationSettingQuery) AllX(ctx context.Context) []*AppUserInvitationSetting {
	nodes, err := auisq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppUserInvitationSetting IDs.
func (auisq *AppUserInvitationSettingQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := auisq.Select(appuserinvitationsetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (auisq *AppUserInvitationSettingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := auisq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (auisq *AppUserInvitationSettingQuery) Count(ctx context.Context) (int, error) {
	if err := auisq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return auisq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (auisq *AppUserInvitationSettingQuery) CountX(ctx context.Context) int {
	count, err := auisq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (auisq *AppUserInvitationSettingQuery) Exist(ctx context.Context) (bool, error) {
	if err := auisq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return auisq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (auisq *AppUserInvitationSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := auisq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppUserInvitationSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (auisq *AppUserInvitationSettingQuery) Clone() *AppUserInvitationSettingQuery {
	if auisq == nil {
		return nil
	}
	return &AppUserInvitationSettingQuery{
		config:     auisq.config,
		limit:      auisq.limit,
		offset:     auisq.offset,
		order:      append([]OrderFunc{}, auisq.order...),
		predicates: append([]predicate.AppUserInvitationSetting{}, auisq.predicates...),
		// clone intermediate query.
		sql:  auisq.sql.Clone(),
		path: auisq.path,
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
//	client.AppUserInvitationSetting.Query().
//		GroupBy(appuserinvitationsetting.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (auisq *AppUserInvitationSettingQuery) GroupBy(field string, fields ...string) *AppUserInvitationSettingGroupBy {
	group := &AppUserInvitationSettingGroupBy{config: auisq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := auisq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return auisq.sqlQuery(ctx), nil
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
//	client.AppUserInvitationSetting.Query().
//		Select(appuserinvitationsetting.FieldAppID).
//		Scan(ctx, &v)
//
func (auisq *AppUserInvitationSettingQuery) Select(fields ...string) *AppUserInvitationSettingSelect {
	auisq.fields = append(auisq.fields, fields...)
	return &AppUserInvitationSettingSelect{AppUserInvitationSettingQuery: auisq}
}

func (auisq *AppUserInvitationSettingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range auisq.fields {
		if !appuserinvitationsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if auisq.path != nil {
		prev, err := auisq.path(ctx)
		if err != nil {
			return err
		}
		auisq.sql = prev
	}
	return nil
}

func (auisq *AppUserInvitationSettingQuery) sqlAll(ctx context.Context) ([]*AppUserInvitationSetting, error) {
	var (
		nodes = []*AppUserInvitationSetting{}
		_spec = auisq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppUserInvitationSetting{config: auisq.config}
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
	if err := sqlgraph.QueryNodes(ctx, auisq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (auisq *AppUserInvitationSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := auisq.querySpec()
	_spec.Node.Columns = auisq.fields
	if len(auisq.fields) > 0 {
		_spec.Unique = auisq.unique != nil && *auisq.unique
	}
	return sqlgraph.CountNodes(ctx, auisq.driver, _spec)
}

func (auisq *AppUserInvitationSettingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := auisq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (auisq *AppUserInvitationSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuserinvitationsetting.Table,
			Columns: appuserinvitationsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserinvitationsetting.FieldID,
			},
		},
		From:   auisq.sql,
		Unique: true,
	}
	if unique := auisq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := auisq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appuserinvitationsetting.FieldID)
		for i := range fields {
			if fields[i] != appuserinvitationsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := auisq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := auisq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := auisq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := auisq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (auisq *AppUserInvitationSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(auisq.driver.Dialect())
	t1 := builder.Table(appuserinvitationsetting.Table)
	columns := auisq.fields
	if len(columns) == 0 {
		columns = appuserinvitationsetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if auisq.sql != nil {
		selector = auisq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if auisq.unique != nil && *auisq.unique {
		selector.Distinct()
	}
	for _, p := range auisq.predicates {
		p(selector)
	}
	for _, p := range auisq.order {
		p(selector)
	}
	if offset := auisq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := auisq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppUserInvitationSettingGroupBy is the group-by builder for AppUserInvitationSetting entities.
type AppUserInvitationSettingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (auisgb *AppUserInvitationSettingGroupBy) Aggregate(fns ...AggregateFunc) *AppUserInvitationSettingGroupBy {
	auisgb.fns = append(auisgb.fns, fns...)
	return auisgb
}

// Scan applies the group-by query and scans the result into the given value.
func (auisgb *AppUserInvitationSettingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := auisgb.path(ctx)
	if err != nil {
		return err
	}
	auisgb.sql = query
	return auisgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (auisgb *AppUserInvitationSettingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := auisgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (auisgb *AppUserInvitationSettingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(auisgb.fields) > 1 {
		return nil, errors.New("ent: AppUserInvitationSettingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := auisgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (auisgb *AppUserInvitationSettingGroupBy) StringsX(ctx context.Context) []string {
	v, err := auisgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auisgb *AppUserInvitationSettingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = auisgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppUserInvitationSettingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (auisgb *AppUserInvitationSettingGroupBy) StringX(ctx context.Context) string {
	v, err := auisgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (auisgb *AppUserInvitationSettingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(auisgb.fields) > 1 {
		return nil, errors.New("ent: AppUserInvitationSettingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := auisgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (auisgb *AppUserInvitationSettingGroupBy) IntsX(ctx context.Context) []int {
	v, err := auisgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auisgb *AppUserInvitationSettingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = auisgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppUserInvitationSettingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (auisgb *AppUserInvitationSettingGroupBy) IntX(ctx context.Context) int {
	v, err := auisgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (auisgb *AppUserInvitationSettingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(auisgb.fields) > 1 {
		return nil, errors.New("ent: AppUserInvitationSettingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := auisgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (auisgb *AppUserInvitationSettingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := auisgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auisgb *AppUserInvitationSettingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = auisgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppUserInvitationSettingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (auisgb *AppUserInvitationSettingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := auisgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (auisgb *AppUserInvitationSettingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(auisgb.fields) > 1 {
		return nil, errors.New("ent: AppUserInvitationSettingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := auisgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (auisgb *AppUserInvitationSettingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := auisgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auisgb *AppUserInvitationSettingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = auisgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppUserInvitationSettingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (auisgb *AppUserInvitationSettingGroupBy) BoolX(ctx context.Context) bool {
	v, err := auisgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (auisgb *AppUserInvitationSettingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range auisgb.fields {
		if !appuserinvitationsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := auisgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := auisgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (auisgb *AppUserInvitationSettingGroupBy) sqlQuery() *sql.Selector {
	selector := auisgb.sql.Select()
	aggregation := make([]string, 0, len(auisgb.fns))
	for _, fn := range auisgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(auisgb.fields)+len(auisgb.fns))
		for _, f := range auisgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(auisgb.fields...)...)
}

// AppUserInvitationSettingSelect is the builder for selecting fields of AppUserInvitationSetting entities.
type AppUserInvitationSettingSelect struct {
	*AppUserInvitationSettingQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (auiss *AppUserInvitationSettingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := auiss.prepareQuery(ctx); err != nil {
		return err
	}
	auiss.sql = auiss.AppUserInvitationSettingQuery.sqlQuery(ctx)
	return auiss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (auiss *AppUserInvitationSettingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := auiss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (auiss *AppUserInvitationSettingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(auiss.fields) > 1 {
		return nil, errors.New("ent: AppUserInvitationSettingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := auiss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (auiss *AppUserInvitationSettingSelect) StringsX(ctx context.Context) []string {
	v, err := auiss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (auiss *AppUserInvitationSettingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = auiss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppUserInvitationSettingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (auiss *AppUserInvitationSettingSelect) StringX(ctx context.Context) string {
	v, err := auiss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (auiss *AppUserInvitationSettingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(auiss.fields) > 1 {
		return nil, errors.New("ent: AppUserInvitationSettingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := auiss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (auiss *AppUserInvitationSettingSelect) IntsX(ctx context.Context) []int {
	v, err := auiss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (auiss *AppUserInvitationSettingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = auiss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppUserInvitationSettingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (auiss *AppUserInvitationSettingSelect) IntX(ctx context.Context) int {
	v, err := auiss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (auiss *AppUserInvitationSettingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(auiss.fields) > 1 {
		return nil, errors.New("ent: AppUserInvitationSettingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := auiss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (auiss *AppUserInvitationSettingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := auiss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (auiss *AppUserInvitationSettingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = auiss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppUserInvitationSettingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (auiss *AppUserInvitationSettingSelect) Float64X(ctx context.Context) float64 {
	v, err := auiss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (auiss *AppUserInvitationSettingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(auiss.fields) > 1 {
		return nil, errors.New("ent: AppUserInvitationSettingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := auiss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (auiss *AppUserInvitationSettingSelect) BoolsX(ctx context.Context) []bool {
	v, err := auiss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (auiss *AppUserInvitationSettingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = auiss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppUserInvitationSettingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (auiss *AppUserInvitationSettingSelect) BoolX(ctx context.Context) bool {
	v, err := auiss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (auiss *AppUserInvitationSettingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := auiss.sql.Query()
	if err := auiss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}