// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponallocated"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CouponAllocatedUpdate is the builder for updating CouponAllocated entities.
type CouponAllocatedUpdate struct {
	config
	hooks    []Hook
	mutation *CouponAllocatedMutation
}

// Where appends a list predicates to the CouponAllocatedUpdate builder.
func (cau *CouponAllocatedUpdate) Where(ps ...predicate.CouponAllocated) *CouponAllocatedUpdate {
	cau.mutation.Where(ps...)
	return cau
}

// SetUserID sets the "user_id" field.
func (cau *CouponAllocatedUpdate) SetUserID(u uuid.UUID) *CouponAllocatedUpdate {
	cau.mutation.SetUserID(u)
	return cau
}

// SetAppID sets the "app_id" field.
func (cau *CouponAllocatedUpdate) SetAppID(u uuid.UUID) *CouponAllocatedUpdate {
	cau.mutation.SetAppID(u)
	return cau
}

// SetUsed sets the "used" field.
func (cau *CouponAllocatedUpdate) SetUsed(b bool) *CouponAllocatedUpdate {
	cau.mutation.SetUsed(b)
	return cau
}

// SetNillableUsed sets the "used" field if the given value is not nil.
func (cau *CouponAllocatedUpdate) SetNillableUsed(b *bool) *CouponAllocatedUpdate {
	if b != nil {
		cau.SetUsed(*b)
	}
	return cau
}

// SetCouponID sets the "coupon_id" field.
func (cau *CouponAllocatedUpdate) SetCouponID(u uuid.UUID) *CouponAllocatedUpdate {
	cau.mutation.SetCouponID(u)
	return cau
}

// SetCreateAt sets the "create_at" field.
func (cau *CouponAllocatedUpdate) SetCreateAt(u uint32) *CouponAllocatedUpdate {
	cau.mutation.ResetCreateAt()
	cau.mutation.SetCreateAt(u)
	return cau
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cau *CouponAllocatedUpdate) SetNillableCreateAt(u *uint32) *CouponAllocatedUpdate {
	if u != nil {
		cau.SetCreateAt(*u)
	}
	return cau
}

// AddCreateAt adds u to the "create_at" field.
func (cau *CouponAllocatedUpdate) AddCreateAt(u uint32) *CouponAllocatedUpdate {
	cau.mutation.AddCreateAt(u)
	return cau
}

// SetUpdateAt sets the "update_at" field.
func (cau *CouponAllocatedUpdate) SetUpdateAt(u uint32) *CouponAllocatedUpdate {
	cau.mutation.ResetUpdateAt()
	cau.mutation.SetUpdateAt(u)
	return cau
}

// AddUpdateAt adds u to the "update_at" field.
func (cau *CouponAllocatedUpdate) AddUpdateAt(u uint32) *CouponAllocatedUpdate {
	cau.mutation.AddUpdateAt(u)
	return cau
}

// SetDeleteAt sets the "delete_at" field.
func (cau *CouponAllocatedUpdate) SetDeleteAt(u uint32) *CouponAllocatedUpdate {
	cau.mutation.ResetDeleteAt()
	cau.mutation.SetDeleteAt(u)
	return cau
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cau *CouponAllocatedUpdate) SetNillableDeleteAt(u *uint32) *CouponAllocatedUpdate {
	if u != nil {
		cau.SetDeleteAt(*u)
	}
	return cau
}

// AddDeleteAt adds u to the "delete_at" field.
func (cau *CouponAllocatedUpdate) AddDeleteAt(u uint32) *CouponAllocatedUpdate {
	cau.mutation.AddDeleteAt(u)
	return cau
}

// Mutation returns the CouponAllocatedMutation object of the builder.
func (cau *CouponAllocatedUpdate) Mutation() *CouponAllocatedMutation {
	return cau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cau *CouponAllocatedUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cau.defaults()
	if len(cau.hooks) == 0 {
		affected, err = cau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponAllocatedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cau.mutation = mutation
			affected, err = cau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cau.hooks) - 1; i >= 0; i-- {
			if cau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cau *CouponAllocatedUpdate) SaveX(ctx context.Context) int {
	affected, err := cau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cau *CouponAllocatedUpdate) Exec(ctx context.Context) error {
	_, err := cau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cau *CouponAllocatedUpdate) ExecX(ctx context.Context) {
	if err := cau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cau *CouponAllocatedUpdate) defaults() {
	if _, ok := cau.mutation.UpdateAt(); !ok {
		v := couponallocated.UpdateDefaultUpdateAt()
		cau.mutation.SetUpdateAt(v)
	}
}

func (cau *CouponAllocatedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponallocated.Table,
			Columns: couponallocated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponallocated.FieldID,
			},
		},
	}
	if ps := cau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cau.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldUserID,
		})
	}
	if value, ok := cau.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldAppID,
		})
	}
	if value, ok := cau.mutation.Used(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: couponallocated.FieldUsed,
		})
	}
	if value, ok := cau.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldCouponID,
		})
	}
	if value, ok := cau.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldCreateAt,
		})
	}
	if value, ok := cau.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldCreateAt,
		})
	}
	if value, ok := cau.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUpdateAt,
		})
	}
	if value, ok := cau.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUpdateAt,
		})
	}
	if value, ok := cau.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldDeleteAt,
		})
	}
	if value, ok := cau.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{couponallocated.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CouponAllocatedUpdateOne is the builder for updating a single CouponAllocated entity.
type CouponAllocatedUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CouponAllocatedMutation
}

// SetUserID sets the "user_id" field.
func (cauo *CouponAllocatedUpdateOne) SetUserID(u uuid.UUID) *CouponAllocatedUpdateOne {
	cauo.mutation.SetUserID(u)
	return cauo
}

// SetAppID sets the "app_id" field.
func (cauo *CouponAllocatedUpdateOne) SetAppID(u uuid.UUID) *CouponAllocatedUpdateOne {
	cauo.mutation.SetAppID(u)
	return cauo
}

// SetUsed sets the "used" field.
func (cauo *CouponAllocatedUpdateOne) SetUsed(b bool) *CouponAllocatedUpdateOne {
	cauo.mutation.SetUsed(b)
	return cauo
}

// SetNillableUsed sets the "used" field if the given value is not nil.
func (cauo *CouponAllocatedUpdateOne) SetNillableUsed(b *bool) *CouponAllocatedUpdateOne {
	if b != nil {
		cauo.SetUsed(*b)
	}
	return cauo
}

// SetCouponID sets the "coupon_id" field.
func (cauo *CouponAllocatedUpdateOne) SetCouponID(u uuid.UUID) *CouponAllocatedUpdateOne {
	cauo.mutation.SetCouponID(u)
	return cauo
}

// SetCreateAt sets the "create_at" field.
func (cauo *CouponAllocatedUpdateOne) SetCreateAt(u uint32) *CouponAllocatedUpdateOne {
	cauo.mutation.ResetCreateAt()
	cauo.mutation.SetCreateAt(u)
	return cauo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cauo *CouponAllocatedUpdateOne) SetNillableCreateAt(u *uint32) *CouponAllocatedUpdateOne {
	if u != nil {
		cauo.SetCreateAt(*u)
	}
	return cauo
}

// AddCreateAt adds u to the "create_at" field.
func (cauo *CouponAllocatedUpdateOne) AddCreateAt(u uint32) *CouponAllocatedUpdateOne {
	cauo.mutation.AddCreateAt(u)
	return cauo
}

// SetUpdateAt sets the "update_at" field.
func (cauo *CouponAllocatedUpdateOne) SetUpdateAt(u uint32) *CouponAllocatedUpdateOne {
	cauo.mutation.ResetUpdateAt()
	cauo.mutation.SetUpdateAt(u)
	return cauo
}

// AddUpdateAt adds u to the "update_at" field.
func (cauo *CouponAllocatedUpdateOne) AddUpdateAt(u uint32) *CouponAllocatedUpdateOne {
	cauo.mutation.AddUpdateAt(u)
	return cauo
}

// SetDeleteAt sets the "delete_at" field.
func (cauo *CouponAllocatedUpdateOne) SetDeleteAt(u uint32) *CouponAllocatedUpdateOne {
	cauo.mutation.ResetDeleteAt()
	cauo.mutation.SetDeleteAt(u)
	return cauo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cauo *CouponAllocatedUpdateOne) SetNillableDeleteAt(u *uint32) *CouponAllocatedUpdateOne {
	if u != nil {
		cauo.SetDeleteAt(*u)
	}
	return cauo
}

// AddDeleteAt adds u to the "delete_at" field.
func (cauo *CouponAllocatedUpdateOne) AddDeleteAt(u uint32) *CouponAllocatedUpdateOne {
	cauo.mutation.AddDeleteAt(u)
	return cauo
}

// Mutation returns the CouponAllocatedMutation object of the builder.
func (cauo *CouponAllocatedUpdateOne) Mutation() *CouponAllocatedMutation {
	return cauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cauo *CouponAllocatedUpdateOne) Select(field string, fields ...string) *CouponAllocatedUpdateOne {
	cauo.fields = append([]string{field}, fields...)
	return cauo
}

// Save executes the query and returns the updated CouponAllocated entity.
func (cauo *CouponAllocatedUpdateOne) Save(ctx context.Context) (*CouponAllocated, error) {
	var (
		err  error
		node *CouponAllocated
	)
	cauo.defaults()
	if len(cauo.hooks) == 0 {
		node, err = cauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponAllocatedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cauo.mutation = mutation
			node, err = cauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cauo.hooks) - 1; i >= 0; i-- {
			if cauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cauo *CouponAllocatedUpdateOne) SaveX(ctx context.Context) *CouponAllocated {
	node, err := cauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cauo *CouponAllocatedUpdateOne) Exec(ctx context.Context) error {
	_, err := cauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cauo *CouponAllocatedUpdateOne) ExecX(ctx context.Context) {
	if err := cauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cauo *CouponAllocatedUpdateOne) defaults() {
	if _, ok := cauo.mutation.UpdateAt(); !ok {
		v := couponallocated.UpdateDefaultUpdateAt()
		cauo.mutation.SetUpdateAt(v)
	}
}

func (cauo *CouponAllocatedUpdateOne) sqlSave(ctx context.Context) (_node *CouponAllocated, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponallocated.Table,
			Columns: couponallocated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponallocated.FieldID,
			},
		},
	}
	id, ok := cauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CouponAllocated.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, couponallocated.FieldID)
		for _, f := range fields {
			if !couponallocated.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != couponallocated.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cauo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldUserID,
		})
	}
	if value, ok := cauo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldAppID,
		})
	}
	if value, ok := cauo.mutation.Used(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: couponallocated.FieldUsed,
		})
	}
	if value, ok := cauo.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldCouponID,
		})
	}
	if value, ok := cauo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldCreateAt,
		})
	}
	if value, ok := cauo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldCreateAt,
		})
	}
	if value, ok := cauo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUpdateAt,
		})
	}
	if value, ok := cauo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUpdateAt,
		})
	}
	if value, ok := cauo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldDeleteAt,
		})
	}
	if value, ok := cauo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldDeleteAt,
		})
	}
	_node = &CouponAllocated{config: cauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{couponallocated.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
