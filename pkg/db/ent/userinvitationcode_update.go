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

// UserInvitationCodeUpdate is the builder for updating UserInvitationCode entities.
type UserInvitationCodeUpdate struct {
	config
	hooks    []Hook
	mutation *UserInvitationCodeMutation
}

// Where appends a list predicates to the UserInvitationCodeUpdate builder.
func (uicu *UserInvitationCodeUpdate) Where(ps ...predicate.UserInvitationCode) *UserInvitationCodeUpdate {
	uicu.mutation.Where(ps...)
	return uicu
}

// Mutation returns the UserInvitationCodeMutation object of the builder.
func (uicu *UserInvitationCodeUpdate) Mutation() *UserInvitationCodeMutation {
	return uicu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uicu *UserInvitationCodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uicu.hooks) == 0 {
		affected, err = uicu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserInvitationCodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uicu.mutation = mutation
			affected, err = uicu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uicu.hooks) - 1; i >= 0; i-- {
			if uicu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uicu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uicu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uicu *UserInvitationCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := uicu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uicu *UserInvitationCodeUpdate) Exec(ctx context.Context) error {
	_, err := uicu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uicu *UserInvitationCodeUpdate) ExecX(ctx context.Context) {
	if err := uicu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uicu *UserInvitationCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userinvitationcode.Table,
			Columns: userinvitationcode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userinvitationcode.FieldID,
			},
		},
	}
	if ps := uicu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uicu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userinvitationcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserInvitationCodeUpdateOne is the builder for updating a single UserInvitationCode entity.
type UserInvitationCodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserInvitationCodeMutation
}

// Mutation returns the UserInvitationCodeMutation object of the builder.
func (uicuo *UserInvitationCodeUpdateOne) Mutation() *UserInvitationCodeMutation {
	return uicuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uicuo *UserInvitationCodeUpdateOne) Select(field string, fields ...string) *UserInvitationCodeUpdateOne {
	uicuo.fields = append([]string{field}, fields...)
	return uicuo
}

// Save executes the query and returns the updated UserInvitationCode entity.
func (uicuo *UserInvitationCodeUpdateOne) Save(ctx context.Context) (*UserInvitationCode, error) {
	var (
		err  error
		node *UserInvitationCode
	)
	if len(uicuo.hooks) == 0 {
		node, err = uicuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserInvitationCodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uicuo.mutation = mutation
			node, err = uicuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uicuo.hooks) - 1; i >= 0; i-- {
			if uicuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uicuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uicuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uicuo *UserInvitationCodeUpdateOne) SaveX(ctx context.Context) *UserInvitationCode {
	node, err := uicuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uicuo *UserInvitationCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := uicuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uicuo *UserInvitationCodeUpdateOne) ExecX(ctx context.Context) {
	if err := uicuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uicuo *UserInvitationCodeUpdateOne) sqlSave(ctx context.Context) (_node *UserInvitationCode, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userinvitationcode.Table,
			Columns: userinvitationcode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userinvitationcode.FieldID,
			},
		},
	}
	id, ok := uicuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserInvitationCode.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uicuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userinvitationcode.FieldID)
		for _, f := range fields {
			if !userinvitationcode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userinvitationcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uicuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &UserInvitationCode{config: uicuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uicuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userinvitationcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
