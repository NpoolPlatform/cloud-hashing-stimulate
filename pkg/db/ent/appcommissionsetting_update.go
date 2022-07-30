// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcommissionsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppCommissionSettingUpdate is the builder for updating AppCommissionSetting entities.
type AppCommissionSettingUpdate struct {
	config
	hooks    []Hook
	mutation *AppCommissionSettingMutation
}

// Where appends a list predicates to the AppCommissionSettingUpdate builder.
func (acsu *AppCommissionSettingUpdate) Where(ps ...predicate.AppCommissionSetting) *AppCommissionSettingUpdate {
	acsu.mutation.Where(ps...)
	return acsu
}

// SetAppID sets the "app_id" field.
func (acsu *AppCommissionSettingUpdate) SetAppID(u uuid.UUID) *AppCommissionSettingUpdate {
	acsu.mutation.SetAppID(u)
	return acsu
}

// SetType sets the "type" field.
func (acsu *AppCommissionSettingUpdate) SetType(s string) *AppCommissionSettingUpdate {
	acsu.mutation.SetType(s)
	return acsu
}

// SetLevel sets the "level" field.
func (acsu *AppCommissionSettingUpdate) SetLevel(u uint32) *AppCommissionSettingUpdate {
	acsu.mutation.ResetLevel()
	acsu.mutation.SetLevel(u)
	return acsu
}

// AddLevel adds u to the "level" field.
func (acsu *AppCommissionSettingUpdate) AddLevel(u int32) *AppCommissionSettingUpdate {
	acsu.mutation.AddLevel(u)
	return acsu
}

// SetInvitationDiscount sets the "invitation_discount" field.
func (acsu *AppCommissionSettingUpdate) SetInvitationDiscount(b bool) *AppCommissionSettingUpdate {
	acsu.mutation.SetInvitationDiscount(b)
	return acsu
}

// SetUniqueSetting sets the "unique_setting" field.
func (acsu *AppCommissionSettingUpdate) SetUniqueSetting(b bool) *AppCommissionSettingUpdate {
	acsu.mutation.SetUniqueSetting(b)
	return acsu
}

// SetKpiSetting sets the "kpi_setting" field.
func (acsu *AppCommissionSettingUpdate) SetKpiSetting(b bool) *AppCommissionSettingUpdate {
	acsu.mutation.SetKpiSetting(b)
	return acsu
}

// SetCreateAt sets the "create_at" field.
func (acsu *AppCommissionSettingUpdate) SetCreateAt(u uint32) *AppCommissionSettingUpdate {
	acsu.mutation.ResetCreateAt()
	acsu.mutation.SetCreateAt(u)
	return acsu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (acsu *AppCommissionSettingUpdate) SetNillableCreateAt(u *uint32) *AppCommissionSettingUpdate {
	if u != nil {
		acsu.SetCreateAt(*u)
	}
	return acsu
}

// AddCreateAt adds u to the "create_at" field.
func (acsu *AppCommissionSettingUpdate) AddCreateAt(u int32) *AppCommissionSettingUpdate {
	acsu.mutation.AddCreateAt(u)
	return acsu
}

// SetUpdateAt sets the "update_at" field.
func (acsu *AppCommissionSettingUpdate) SetUpdateAt(u uint32) *AppCommissionSettingUpdate {
	acsu.mutation.ResetUpdateAt()
	acsu.mutation.SetUpdateAt(u)
	return acsu
}

// AddUpdateAt adds u to the "update_at" field.
func (acsu *AppCommissionSettingUpdate) AddUpdateAt(u int32) *AppCommissionSettingUpdate {
	acsu.mutation.AddUpdateAt(u)
	return acsu
}

// SetDeleteAt sets the "delete_at" field.
func (acsu *AppCommissionSettingUpdate) SetDeleteAt(u uint32) *AppCommissionSettingUpdate {
	acsu.mutation.ResetDeleteAt()
	acsu.mutation.SetDeleteAt(u)
	return acsu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (acsu *AppCommissionSettingUpdate) SetNillableDeleteAt(u *uint32) *AppCommissionSettingUpdate {
	if u != nil {
		acsu.SetDeleteAt(*u)
	}
	return acsu
}

// AddDeleteAt adds u to the "delete_at" field.
func (acsu *AppCommissionSettingUpdate) AddDeleteAt(u int32) *AppCommissionSettingUpdate {
	acsu.mutation.AddDeleteAt(u)
	return acsu
}

// Mutation returns the AppCommissionSettingMutation object of the builder.
func (acsu *AppCommissionSettingUpdate) Mutation() *AppCommissionSettingMutation {
	return acsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acsu *AppCommissionSettingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	acsu.defaults()
	if len(acsu.hooks) == 0 {
		affected, err = acsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppCommissionSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			acsu.mutation = mutation
			affected, err = acsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(acsu.hooks) - 1; i >= 0; i-- {
			if acsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (acsu *AppCommissionSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := acsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acsu *AppCommissionSettingUpdate) Exec(ctx context.Context) error {
	_, err := acsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acsu *AppCommissionSettingUpdate) ExecX(ctx context.Context) {
	if err := acsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acsu *AppCommissionSettingUpdate) defaults() {
	if _, ok := acsu.mutation.UpdateAt(); !ok {
		v := appcommissionsetting.UpdateDefaultUpdateAt()
		acsu.mutation.SetUpdateAt(v)
	}
}

func (acsu *AppCommissionSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appcommissionsetting.Table,
			Columns: appcommissionsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcommissionsetting.FieldID,
			},
		},
	}
	if ps := acsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acsu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcommissionsetting.FieldAppID,
		})
	}
	if value, ok := acsu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appcommissionsetting.FieldType,
		})
	}
	if value, ok := acsu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldLevel,
		})
	}
	if value, ok := acsu.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldLevel,
		})
	}
	if value, ok := acsu.mutation.InvitationDiscount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appcommissionsetting.FieldInvitationDiscount,
		})
	}
	if value, ok := acsu.mutation.UniqueSetting(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appcommissionsetting.FieldUniqueSetting,
		})
	}
	if value, ok := acsu.mutation.KpiSetting(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appcommissionsetting.FieldKpiSetting,
		})
	}
	if value, ok := acsu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldCreateAt,
		})
	}
	if value, ok := acsu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldCreateAt,
		})
	}
	if value, ok := acsu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldUpdateAt,
		})
	}
	if value, ok := acsu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldUpdateAt,
		})
	}
	if value, ok := acsu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldDeleteAt,
		})
	}
	if value, ok := acsu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appcommissionsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppCommissionSettingUpdateOne is the builder for updating a single AppCommissionSetting entity.
type AppCommissionSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppCommissionSettingMutation
}

// SetAppID sets the "app_id" field.
func (acsuo *AppCommissionSettingUpdateOne) SetAppID(u uuid.UUID) *AppCommissionSettingUpdateOne {
	acsuo.mutation.SetAppID(u)
	return acsuo
}

// SetType sets the "type" field.
func (acsuo *AppCommissionSettingUpdateOne) SetType(s string) *AppCommissionSettingUpdateOne {
	acsuo.mutation.SetType(s)
	return acsuo
}

// SetLevel sets the "level" field.
func (acsuo *AppCommissionSettingUpdateOne) SetLevel(u uint32) *AppCommissionSettingUpdateOne {
	acsuo.mutation.ResetLevel()
	acsuo.mutation.SetLevel(u)
	return acsuo
}

// AddLevel adds u to the "level" field.
func (acsuo *AppCommissionSettingUpdateOne) AddLevel(u int32) *AppCommissionSettingUpdateOne {
	acsuo.mutation.AddLevel(u)
	return acsuo
}

// SetInvitationDiscount sets the "invitation_discount" field.
func (acsuo *AppCommissionSettingUpdateOne) SetInvitationDiscount(b bool) *AppCommissionSettingUpdateOne {
	acsuo.mutation.SetInvitationDiscount(b)
	return acsuo
}

// SetUniqueSetting sets the "unique_setting" field.
func (acsuo *AppCommissionSettingUpdateOne) SetUniqueSetting(b bool) *AppCommissionSettingUpdateOne {
	acsuo.mutation.SetUniqueSetting(b)
	return acsuo
}

// SetKpiSetting sets the "kpi_setting" field.
func (acsuo *AppCommissionSettingUpdateOne) SetKpiSetting(b bool) *AppCommissionSettingUpdateOne {
	acsuo.mutation.SetKpiSetting(b)
	return acsuo
}

// SetCreateAt sets the "create_at" field.
func (acsuo *AppCommissionSettingUpdateOne) SetCreateAt(u uint32) *AppCommissionSettingUpdateOne {
	acsuo.mutation.ResetCreateAt()
	acsuo.mutation.SetCreateAt(u)
	return acsuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (acsuo *AppCommissionSettingUpdateOne) SetNillableCreateAt(u *uint32) *AppCommissionSettingUpdateOne {
	if u != nil {
		acsuo.SetCreateAt(*u)
	}
	return acsuo
}

// AddCreateAt adds u to the "create_at" field.
func (acsuo *AppCommissionSettingUpdateOne) AddCreateAt(u int32) *AppCommissionSettingUpdateOne {
	acsuo.mutation.AddCreateAt(u)
	return acsuo
}

// SetUpdateAt sets the "update_at" field.
func (acsuo *AppCommissionSettingUpdateOne) SetUpdateAt(u uint32) *AppCommissionSettingUpdateOne {
	acsuo.mutation.ResetUpdateAt()
	acsuo.mutation.SetUpdateAt(u)
	return acsuo
}

// AddUpdateAt adds u to the "update_at" field.
func (acsuo *AppCommissionSettingUpdateOne) AddUpdateAt(u int32) *AppCommissionSettingUpdateOne {
	acsuo.mutation.AddUpdateAt(u)
	return acsuo
}

// SetDeleteAt sets the "delete_at" field.
func (acsuo *AppCommissionSettingUpdateOne) SetDeleteAt(u uint32) *AppCommissionSettingUpdateOne {
	acsuo.mutation.ResetDeleteAt()
	acsuo.mutation.SetDeleteAt(u)
	return acsuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (acsuo *AppCommissionSettingUpdateOne) SetNillableDeleteAt(u *uint32) *AppCommissionSettingUpdateOne {
	if u != nil {
		acsuo.SetDeleteAt(*u)
	}
	return acsuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (acsuo *AppCommissionSettingUpdateOne) AddDeleteAt(u int32) *AppCommissionSettingUpdateOne {
	acsuo.mutation.AddDeleteAt(u)
	return acsuo
}

// Mutation returns the AppCommissionSettingMutation object of the builder.
func (acsuo *AppCommissionSettingUpdateOne) Mutation() *AppCommissionSettingMutation {
	return acsuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acsuo *AppCommissionSettingUpdateOne) Select(field string, fields ...string) *AppCommissionSettingUpdateOne {
	acsuo.fields = append([]string{field}, fields...)
	return acsuo
}

// Save executes the query and returns the updated AppCommissionSetting entity.
func (acsuo *AppCommissionSettingUpdateOne) Save(ctx context.Context) (*AppCommissionSetting, error) {
	var (
		err  error
		node *AppCommissionSetting
	)
	acsuo.defaults()
	if len(acsuo.hooks) == 0 {
		node, err = acsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppCommissionSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			acsuo.mutation = mutation
			node, err = acsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(acsuo.hooks) - 1; i >= 0; i-- {
			if acsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acsuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, acsuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppCommissionSetting)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppCommissionSettingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (acsuo *AppCommissionSettingUpdateOne) SaveX(ctx context.Context) *AppCommissionSetting {
	node, err := acsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acsuo *AppCommissionSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := acsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acsuo *AppCommissionSettingUpdateOne) ExecX(ctx context.Context) {
	if err := acsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acsuo *AppCommissionSettingUpdateOne) defaults() {
	if _, ok := acsuo.mutation.UpdateAt(); !ok {
		v := appcommissionsetting.UpdateDefaultUpdateAt()
		acsuo.mutation.SetUpdateAt(v)
	}
}

func (acsuo *AppCommissionSettingUpdateOne) sqlSave(ctx context.Context) (_node *AppCommissionSetting, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appcommissionsetting.Table,
			Columns: appcommissionsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcommissionsetting.FieldID,
			},
		},
	}
	id, ok := acsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppCommissionSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appcommissionsetting.FieldID)
		for _, f := range fields {
			if !appcommissionsetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appcommissionsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acsuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcommissionsetting.FieldAppID,
		})
	}
	if value, ok := acsuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appcommissionsetting.FieldType,
		})
	}
	if value, ok := acsuo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldLevel,
		})
	}
	if value, ok := acsuo.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldLevel,
		})
	}
	if value, ok := acsuo.mutation.InvitationDiscount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appcommissionsetting.FieldInvitationDiscount,
		})
	}
	if value, ok := acsuo.mutation.UniqueSetting(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appcommissionsetting.FieldUniqueSetting,
		})
	}
	if value, ok := acsuo.mutation.KpiSetting(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appcommissionsetting.FieldKpiSetting,
		})
	}
	if value, ok := acsuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldCreateAt,
		})
	}
	if value, ok := acsuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldCreateAt,
		})
	}
	if value, ok := acsuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldUpdateAt,
		})
	}
	if value, ok := acsuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldUpdateAt,
		})
	}
	if value, ok := acsuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldDeleteAt,
		})
	}
	if value, ok := acsuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcommissionsetting.FieldDeleteAt,
		})
	}
	_node = &AppCommissionSetting{config: acsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appcommissionsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
