// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userinvitationcode"
)

// UserInvitationCodeDelete is the builder for deleting a UserInvitationCode entity.
type UserInvitationCodeDelete struct {
	config
	hooks    []Hook
	mutation *UserInvitationCodeMutation
}

// Where appends a list predicates to the UserInvitationCodeDelete builder.
func (uicd *UserInvitationCodeDelete) Where(ps ...predicate.UserInvitationCode) *UserInvitationCodeDelete {
	uicd.mutation.Where(ps...)
	return uicd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uicd *UserInvitationCodeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uicd.hooks) == 0 {
		affected, err = uicd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserInvitationCodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uicd.mutation = mutation
			affected, err = uicd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uicd.hooks) - 1; i >= 0; i-- {
			if uicd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uicd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uicd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uicd *UserInvitationCodeDelete) ExecX(ctx context.Context) int {
	n, err := uicd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uicd *UserInvitationCodeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userinvitationcode.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userinvitationcode.FieldID,
			},
		},
	}
	if ps := uicd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, uicd.driver, _spec)
}

// UserInvitationCodeDeleteOne is the builder for deleting a single UserInvitationCode entity.
type UserInvitationCodeDeleteOne struct {
	uicd *UserInvitationCodeDelete
}

// Exec executes the deletion query.
func (uicdo *UserInvitationCodeDeleteOne) Exec(ctx context.Context) error {
	n, err := uicdo.uicd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userinvitationcode.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uicdo *UserInvitationCodeDeleteOne) ExecX(ctx context.Context) {
	uicdo.uicd.ExecX(ctx)
}
