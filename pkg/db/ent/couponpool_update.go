// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponpool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CouponPoolUpdate is the builder for updating CouponPool entities.
type CouponPoolUpdate struct {
	config
	hooks    []Hook
	mutation *CouponPoolMutation
}

// Where appends a list predicates to the CouponPoolUpdate builder.
func (cpu *CouponPoolUpdate) Where(ps ...predicate.CouponPool) *CouponPoolUpdate {
	cpu.mutation.Where(ps...)
	return cpu
}

// SetAppID sets the "app_id" field.
func (cpu *CouponPoolUpdate) SetAppID(u uuid.UUID) *CouponPoolUpdate {
	cpu.mutation.SetAppID(u)
	return cpu
}

// SetDenomination sets the "denomination" field.
func (cpu *CouponPoolUpdate) SetDenomination(u uint64) *CouponPoolUpdate {
	cpu.mutation.ResetDenomination()
	cpu.mutation.SetDenomination(u)
	return cpu
}

// AddDenomination adds u to the "denomination" field.
func (cpu *CouponPoolUpdate) AddDenomination(u int64) *CouponPoolUpdate {
	cpu.mutation.AddDenomination(u)
	return cpu
}

// SetCirculation sets the "circulation" field.
func (cpu *CouponPoolUpdate) SetCirculation(i int32) *CouponPoolUpdate {
	cpu.mutation.ResetCirculation()
	cpu.mutation.SetCirculation(i)
	return cpu
}

// AddCirculation adds i to the "circulation" field.
func (cpu *CouponPoolUpdate) AddCirculation(i int32) *CouponPoolUpdate {
	cpu.mutation.AddCirculation(i)
	return cpu
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (cpu *CouponPoolUpdate) SetReleaseByUserID(u uuid.UUID) *CouponPoolUpdate {
	cpu.mutation.SetReleaseByUserID(u)
	return cpu
}

// SetStart sets the "start" field.
func (cpu *CouponPoolUpdate) SetStart(u uint32) *CouponPoolUpdate {
	cpu.mutation.ResetStart()
	cpu.mutation.SetStart(u)
	return cpu
}

// AddStart adds u to the "start" field.
func (cpu *CouponPoolUpdate) AddStart(u int32) *CouponPoolUpdate {
	cpu.mutation.AddStart(u)
	return cpu
}

// SetDurationDays sets the "duration_days" field.
func (cpu *CouponPoolUpdate) SetDurationDays(i int32) *CouponPoolUpdate {
	cpu.mutation.ResetDurationDays()
	cpu.mutation.SetDurationDays(i)
	return cpu
}

// AddDurationDays adds i to the "duration_days" field.
func (cpu *CouponPoolUpdate) AddDurationDays(i int32) *CouponPoolUpdate {
	cpu.mutation.AddDurationDays(i)
	return cpu
}

// SetMessage sets the "message" field.
func (cpu *CouponPoolUpdate) SetMessage(s string) *CouponPoolUpdate {
	cpu.mutation.SetMessage(s)
	return cpu
}

// SetName sets the "name" field.
func (cpu *CouponPoolUpdate) SetName(s string) *CouponPoolUpdate {
	cpu.mutation.SetName(s)
	return cpu
}

// SetCreateAt sets the "create_at" field.
func (cpu *CouponPoolUpdate) SetCreateAt(u uint32) *CouponPoolUpdate {
	cpu.mutation.ResetCreateAt()
	cpu.mutation.SetCreateAt(u)
	return cpu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cpu *CouponPoolUpdate) SetNillableCreateAt(u *uint32) *CouponPoolUpdate {
	if u != nil {
		cpu.SetCreateAt(*u)
	}
	return cpu
}

// AddCreateAt adds u to the "create_at" field.
func (cpu *CouponPoolUpdate) AddCreateAt(u int32) *CouponPoolUpdate {
	cpu.mutation.AddCreateAt(u)
	return cpu
}

// SetUpdateAt sets the "update_at" field.
func (cpu *CouponPoolUpdate) SetUpdateAt(u uint32) *CouponPoolUpdate {
	cpu.mutation.ResetUpdateAt()
	cpu.mutation.SetUpdateAt(u)
	return cpu
}

// AddUpdateAt adds u to the "update_at" field.
func (cpu *CouponPoolUpdate) AddUpdateAt(u int32) *CouponPoolUpdate {
	cpu.mutation.AddUpdateAt(u)
	return cpu
}

// SetDeleteAt sets the "delete_at" field.
func (cpu *CouponPoolUpdate) SetDeleteAt(u uint32) *CouponPoolUpdate {
	cpu.mutation.ResetDeleteAt()
	cpu.mutation.SetDeleteAt(u)
	return cpu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cpu *CouponPoolUpdate) SetNillableDeleteAt(u *uint32) *CouponPoolUpdate {
	if u != nil {
		cpu.SetDeleteAt(*u)
	}
	return cpu
}

// AddDeleteAt adds u to the "delete_at" field.
func (cpu *CouponPoolUpdate) AddDeleteAt(u int32) *CouponPoolUpdate {
	cpu.mutation.AddDeleteAt(u)
	return cpu
}

// Mutation returns the CouponPoolMutation object of the builder.
func (cpu *CouponPoolUpdate) Mutation() *CouponPoolMutation {
	return cpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cpu *CouponPoolUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cpu.defaults()
	if len(cpu.hooks) == 0 {
		if err = cpu.check(); err != nil {
			return 0, err
		}
		affected, err = cpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponPoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpu.check(); err != nil {
				return 0, err
			}
			cpu.mutation = mutation
			affected, err = cpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cpu.hooks) - 1; i >= 0; i-- {
			if cpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cpu *CouponPoolUpdate) SaveX(ctx context.Context) int {
	affected, err := cpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cpu *CouponPoolUpdate) Exec(ctx context.Context) error {
	_, err := cpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpu *CouponPoolUpdate) ExecX(ctx context.Context) {
	if err := cpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpu *CouponPoolUpdate) defaults() {
	if _, ok := cpu.mutation.UpdateAt(); !ok {
		v := couponpool.UpdateDefaultUpdateAt()
		cpu.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpu *CouponPoolUpdate) check() error {
	if v, ok := cpu.mutation.Message(); ok {
		if err := couponpool.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "CouponPool.message": %w`, err)}
		}
	}
	if v, ok := cpu.mutation.Name(); ok {
		if err := couponpool.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CouponPool.name": %w`, err)}
		}
	}
	return nil
}

func (cpu *CouponPoolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponpool.Table,
			Columns: couponpool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponpool.FieldID,
			},
		},
	}
	if ps := cpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponpool.FieldAppID,
		})
	}
	if value, ok := cpu.mutation.Denomination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: couponpool.FieldDenomination,
		})
	}
	if value, ok := cpu.mutation.AddedDenomination(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: couponpool.FieldDenomination,
		})
	}
	if value, ok := cpu.mutation.Circulation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: couponpool.FieldCirculation,
		})
	}
	if value, ok := cpu.mutation.AddedCirculation(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: couponpool.FieldCirculation,
		})
	}
	if value, ok := cpu.mutation.ReleaseByUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponpool.FieldReleaseByUserID,
		})
	}
	if value, ok := cpu.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldStart,
		})
	}
	if value, ok := cpu.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldStart,
		})
	}
	if value, ok := cpu.mutation.DurationDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: couponpool.FieldDurationDays,
		})
	}
	if value, ok := cpu.mutation.AddedDurationDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: couponpool.FieldDurationDays,
		})
	}
	if value, ok := cpu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponpool.FieldMessage,
		})
	}
	if value, ok := cpu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponpool.FieldName,
		})
	}
	if value, ok := cpu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldCreateAt,
		})
	}
	if value, ok := cpu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldCreateAt,
		})
	}
	if value, ok := cpu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldUpdateAt,
		})
	}
	if value, ok := cpu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldUpdateAt,
		})
	}
	if value, ok := cpu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldDeleteAt,
		})
	}
	if value, ok := cpu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{couponpool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CouponPoolUpdateOne is the builder for updating a single CouponPool entity.
type CouponPoolUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CouponPoolMutation
}

// SetAppID sets the "app_id" field.
func (cpuo *CouponPoolUpdateOne) SetAppID(u uuid.UUID) *CouponPoolUpdateOne {
	cpuo.mutation.SetAppID(u)
	return cpuo
}

// SetDenomination sets the "denomination" field.
func (cpuo *CouponPoolUpdateOne) SetDenomination(u uint64) *CouponPoolUpdateOne {
	cpuo.mutation.ResetDenomination()
	cpuo.mutation.SetDenomination(u)
	return cpuo
}

// AddDenomination adds u to the "denomination" field.
func (cpuo *CouponPoolUpdateOne) AddDenomination(u int64) *CouponPoolUpdateOne {
	cpuo.mutation.AddDenomination(u)
	return cpuo
}

// SetCirculation sets the "circulation" field.
func (cpuo *CouponPoolUpdateOne) SetCirculation(i int32) *CouponPoolUpdateOne {
	cpuo.mutation.ResetCirculation()
	cpuo.mutation.SetCirculation(i)
	return cpuo
}

// AddCirculation adds i to the "circulation" field.
func (cpuo *CouponPoolUpdateOne) AddCirculation(i int32) *CouponPoolUpdateOne {
	cpuo.mutation.AddCirculation(i)
	return cpuo
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (cpuo *CouponPoolUpdateOne) SetReleaseByUserID(u uuid.UUID) *CouponPoolUpdateOne {
	cpuo.mutation.SetReleaseByUserID(u)
	return cpuo
}

// SetStart sets the "start" field.
func (cpuo *CouponPoolUpdateOne) SetStart(u uint32) *CouponPoolUpdateOne {
	cpuo.mutation.ResetStart()
	cpuo.mutation.SetStart(u)
	return cpuo
}

// AddStart adds u to the "start" field.
func (cpuo *CouponPoolUpdateOne) AddStart(u int32) *CouponPoolUpdateOne {
	cpuo.mutation.AddStart(u)
	return cpuo
}

// SetDurationDays sets the "duration_days" field.
func (cpuo *CouponPoolUpdateOne) SetDurationDays(i int32) *CouponPoolUpdateOne {
	cpuo.mutation.ResetDurationDays()
	cpuo.mutation.SetDurationDays(i)
	return cpuo
}

// AddDurationDays adds i to the "duration_days" field.
func (cpuo *CouponPoolUpdateOne) AddDurationDays(i int32) *CouponPoolUpdateOne {
	cpuo.mutation.AddDurationDays(i)
	return cpuo
}

// SetMessage sets the "message" field.
func (cpuo *CouponPoolUpdateOne) SetMessage(s string) *CouponPoolUpdateOne {
	cpuo.mutation.SetMessage(s)
	return cpuo
}

// SetName sets the "name" field.
func (cpuo *CouponPoolUpdateOne) SetName(s string) *CouponPoolUpdateOne {
	cpuo.mutation.SetName(s)
	return cpuo
}

// SetCreateAt sets the "create_at" field.
func (cpuo *CouponPoolUpdateOne) SetCreateAt(u uint32) *CouponPoolUpdateOne {
	cpuo.mutation.ResetCreateAt()
	cpuo.mutation.SetCreateAt(u)
	return cpuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cpuo *CouponPoolUpdateOne) SetNillableCreateAt(u *uint32) *CouponPoolUpdateOne {
	if u != nil {
		cpuo.SetCreateAt(*u)
	}
	return cpuo
}

// AddCreateAt adds u to the "create_at" field.
func (cpuo *CouponPoolUpdateOne) AddCreateAt(u int32) *CouponPoolUpdateOne {
	cpuo.mutation.AddCreateAt(u)
	return cpuo
}

// SetUpdateAt sets the "update_at" field.
func (cpuo *CouponPoolUpdateOne) SetUpdateAt(u uint32) *CouponPoolUpdateOne {
	cpuo.mutation.ResetUpdateAt()
	cpuo.mutation.SetUpdateAt(u)
	return cpuo
}

// AddUpdateAt adds u to the "update_at" field.
func (cpuo *CouponPoolUpdateOne) AddUpdateAt(u int32) *CouponPoolUpdateOne {
	cpuo.mutation.AddUpdateAt(u)
	return cpuo
}

// SetDeleteAt sets the "delete_at" field.
func (cpuo *CouponPoolUpdateOne) SetDeleteAt(u uint32) *CouponPoolUpdateOne {
	cpuo.mutation.ResetDeleteAt()
	cpuo.mutation.SetDeleteAt(u)
	return cpuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cpuo *CouponPoolUpdateOne) SetNillableDeleteAt(u *uint32) *CouponPoolUpdateOne {
	if u != nil {
		cpuo.SetDeleteAt(*u)
	}
	return cpuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (cpuo *CouponPoolUpdateOne) AddDeleteAt(u int32) *CouponPoolUpdateOne {
	cpuo.mutation.AddDeleteAt(u)
	return cpuo
}

// Mutation returns the CouponPoolMutation object of the builder.
func (cpuo *CouponPoolUpdateOne) Mutation() *CouponPoolMutation {
	return cpuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cpuo *CouponPoolUpdateOne) Select(field string, fields ...string) *CouponPoolUpdateOne {
	cpuo.fields = append([]string{field}, fields...)
	return cpuo
}

// Save executes the query and returns the updated CouponPool entity.
func (cpuo *CouponPoolUpdateOne) Save(ctx context.Context) (*CouponPool, error) {
	var (
		err  error
		node *CouponPool
	)
	cpuo.defaults()
	if len(cpuo.hooks) == 0 {
		if err = cpuo.check(); err != nil {
			return nil, err
		}
		node, err = cpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponPoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpuo.check(); err != nil {
				return nil, err
			}
			cpuo.mutation = mutation
			node, err = cpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cpuo.hooks) - 1; i >= 0; i-- {
			if cpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cpuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CouponPool)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CouponPoolMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cpuo *CouponPoolUpdateOne) SaveX(ctx context.Context) *CouponPool {
	node, err := cpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cpuo *CouponPoolUpdateOne) Exec(ctx context.Context) error {
	_, err := cpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpuo *CouponPoolUpdateOne) ExecX(ctx context.Context) {
	if err := cpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpuo *CouponPoolUpdateOne) defaults() {
	if _, ok := cpuo.mutation.UpdateAt(); !ok {
		v := couponpool.UpdateDefaultUpdateAt()
		cpuo.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpuo *CouponPoolUpdateOne) check() error {
	if v, ok := cpuo.mutation.Message(); ok {
		if err := couponpool.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "CouponPool.message": %w`, err)}
		}
	}
	if v, ok := cpuo.mutation.Name(); ok {
		if err := couponpool.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CouponPool.name": %w`, err)}
		}
	}
	return nil
}

func (cpuo *CouponPoolUpdateOne) sqlSave(ctx context.Context) (_node *CouponPool, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponpool.Table,
			Columns: couponpool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponpool.FieldID,
			},
		},
	}
	id, ok := cpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CouponPool.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, couponpool.FieldID)
		for _, f := range fields {
			if !couponpool.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != couponpool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponpool.FieldAppID,
		})
	}
	if value, ok := cpuo.mutation.Denomination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: couponpool.FieldDenomination,
		})
	}
	if value, ok := cpuo.mutation.AddedDenomination(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: couponpool.FieldDenomination,
		})
	}
	if value, ok := cpuo.mutation.Circulation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: couponpool.FieldCirculation,
		})
	}
	if value, ok := cpuo.mutation.AddedCirculation(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: couponpool.FieldCirculation,
		})
	}
	if value, ok := cpuo.mutation.ReleaseByUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponpool.FieldReleaseByUserID,
		})
	}
	if value, ok := cpuo.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldStart,
		})
	}
	if value, ok := cpuo.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldStart,
		})
	}
	if value, ok := cpuo.mutation.DurationDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: couponpool.FieldDurationDays,
		})
	}
	if value, ok := cpuo.mutation.AddedDurationDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: couponpool.FieldDurationDays,
		})
	}
	if value, ok := cpuo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponpool.FieldMessage,
		})
	}
	if value, ok := cpuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponpool.FieldName,
		})
	}
	if value, ok := cpuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldCreateAt,
		})
	}
	if value, ok := cpuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldCreateAt,
		})
	}
	if value, ok := cpuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldUpdateAt,
		})
	}
	if value, ok := cpuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldUpdateAt,
		})
	}
	if value, ok := cpuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldDeleteAt,
		})
	}
	if value, ok := cpuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldDeleteAt,
		})
	}
	_node = &CouponPool{config: cpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{couponpool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
