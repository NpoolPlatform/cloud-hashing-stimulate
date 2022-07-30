// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponpool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
)

// CouponPoolDelete is the builder for deleting a CouponPool entity.
type CouponPoolDelete struct {
	config
	hooks    []Hook
	mutation *CouponPoolMutation
}

// Where appends a list predicates to the CouponPoolDelete builder.
func (cpd *CouponPoolDelete) Where(ps ...predicate.CouponPool) *CouponPoolDelete {
	cpd.mutation.Where(ps...)
	return cpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cpd *CouponPoolDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cpd.hooks) == 0 {
		affected, err = cpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponPoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cpd.mutation = mutation
			affected, err = cpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cpd.hooks) - 1; i >= 0; i-- {
			if cpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpd *CouponPoolDelete) ExecX(ctx context.Context) int {
	n, err := cpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cpd *CouponPoolDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: couponpool.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponpool.FieldID,
			},
		},
	}
	if ps := cpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CouponPoolDeleteOne is the builder for deleting a single CouponPool entity.
type CouponPoolDeleteOne struct {
	cpd *CouponPoolDelete
}

// Exec executes the deletion query.
func (cpdo *CouponPoolDeleteOne) Exec(ctx context.Context) error {
	n, err := cpdo.cpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{couponpool.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cpdo *CouponPoolDeleteOne) ExecX(ctx context.Context) {
	cpdo.cpd.ExecX(ctx)
}
