// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/defaultkpisetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
)

// DefaultKpiSettingDelete is the builder for deleting a DefaultKpiSetting entity.
type DefaultKpiSettingDelete struct {
	config
	hooks    []Hook
	mutation *DefaultKpiSettingMutation
}

// Where appends a list predicates to the DefaultKpiSettingDelete builder.
func (dksd *DefaultKpiSettingDelete) Where(ps ...predicate.DefaultKpiSetting) *DefaultKpiSettingDelete {
	dksd.mutation.Where(ps...)
	return dksd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dksd *DefaultKpiSettingDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dksd.hooks) == 0 {
		affected, err = dksd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DefaultKpiSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dksd.mutation = mutation
			affected, err = dksd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dksd.hooks) - 1; i >= 0; i-- {
			if dksd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dksd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dksd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dksd *DefaultKpiSettingDelete) ExecX(ctx context.Context) int {
	n, err := dksd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dksd *DefaultKpiSettingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: defaultkpisetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: defaultkpisetting.FieldID,
			},
		},
	}
	if ps := dksd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dksd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DefaultKpiSettingDeleteOne is the builder for deleting a single DefaultKpiSetting entity.
type DefaultKpiSettingDeleteOne struct {
	dksd *DefaultKpiSettingDelete
}

// Exec executes the deletion query.
func (dksdo *DefaultKpiSettingDeleteOne) Exec(ctx context.Context) error {
	n, err := dksdo.dksd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{defaultkpisetting.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dksdo *DefaultKpiSettingDeleteOne) ExecX(ctx context.Context) {
	dksdo.dksd.ExecX(ctx)
}
