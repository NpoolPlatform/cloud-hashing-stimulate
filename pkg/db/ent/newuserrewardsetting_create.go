// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/newuserrewardsetting"
)

// NewUserRewardSettingCreate is the builder for creating a NewUserRewardSetting entity.
type NewUserRewardSettingCreate struct {
	config
	mutation *NewUserRewardSettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// Mutation returns the NewUserRewardSettingMutation object of the builder.
func (nursc *NewUserRewardSettingCreate) Mutation() *NewUserRewardSettingMutation {
	return nursc.mutation
}

// Save creates the NewUserRewardSetting in the database.
func (nursc *NewUserRewardSettingCreate) Save(ctx context.Context) (*NewUserRewardSetting, error) {
	var (
		err  error
		node *NewUserRewardSetting
	)
	if len(nursc.hooks) == 0 {
		if err = nursc.check(); err != nil {
			return nil, err
		}
		node, err = nursc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewUserRewardSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nursc.check(); err != nil {
				return nil, err
			}
			nursc.mutation = mutation
			if node, err = nursc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nursc.hooks) - 1; i >= 0; i-- {
			if nursc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nursc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nursc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nursc *NewUserRewardSettingCreate) SaveX(ctx context.Context) *NewUserRewardSetting {
	v, err := nursc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nursc *NewUserRewardSettingCreate) Exec(ctx context.Context) error {
	_, err := nursc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nursc *NewUserRewardSettingCreate) ExecX(ctx context.Context) {
	if err := nursc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nursc *NewUserRewardSettingCreate) check() error {
	return nil
}

func (nursc *NewUserRewardSettingCreate) sqlSave(ctx context.Context) (*NewUserRewardSetting, error) {
	_node, _spec := nursc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nursc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (nursc *NewUserRewardSettingCreate) createSpec() (*NewUserRewardSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &NewUserRewardSetting{config: nursc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: newuserrewardsetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: newuserrewardsetting.FieldID,
			},
		}
	)
	_spec.OnConflict = nursc.conflict
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NewUserRewardSetting.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (nursc *NewUserRewardSettingCreate) OnConflict(opts ...sql.ConflictOption) *NewUserRewardSettingUpsertOne {
	nursc.conflict = opts
	return &NewUserRewardSettingUpsertOne{
		create: nursc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NewUserRewardSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (nursc *NewUserRewardSettingCreate) OnConflictColumns(columns ...string) *NewUserRewardSettingUpsertOne {
	nursc.conflict = append(nursc.conflict, sql.ConflictColumns(columns...))
	return &NewUserRewardSettingUpsertOne{
		create: nursc,
	}
}

type (
	// NewUserRewardSettingUpsertOne is the builder for "upsert"-ing
	//  one NewUserRewardSetting node.
	NewUserRewardSettingUpsertOne struct {
		create *NewUserRewardSettingCreate
	}

	// NewUserRewardSettingUpsert is the "OnConflict" setter.
	NewUserRewardSettingUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.NewUserRewardSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *NewUserRewardSettingUpsertOne) UpdateNewValues() *NewUserRewardSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.NewUserRewardSetting.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *NewUserRewardSettingUpsertOne) Ignore() *NewUserRewardSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NewUserRewardSettingUpsertOne) DoNothing() *NewUserRewardSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NewUserRewardSettingCreate.OnConflict
// documentation for more info.
func (u *NewUserRewardSettingUpsertOne) Update(set func(*NewUserRewardSettingUpsert)) *NewUserRewardSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NewUserRewardSettingUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *NewUserRewardSettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NewUserRewardSettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NewUserRewardSettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NewUserRewardSettingUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NewUserRewardSettingUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NewUserRewardSettingCreateBulk is the builder for creating many NewUserRewardSetting entities in bulk.
type NewUserRewardSettingCreateBulk struct {
	config
	builders []*NewUserRewardSettingCreate
	conflict []sql.ConflictOption
}

// Save creates the NewUserRewardSetting entities in the database.
func (nurscb *NewUserRewardSettingCreateBulk) Save(ctx context.Context) ([]*NewUserRewardSetting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(nurscb.builders))
	nodes := make([]*NewUserRewardSetting, len(nurscb.builders))
	mutators := make([]Mutator, len(nurscb.builders))
	for i := range nurscb.builders {
		func(i int, root context.Context) {
			builder := nurscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NewUserRewardSettingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, nurscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = nurscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nurscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, nurscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nurscb *NewUserRewardSettingCreateBulk) SaveX(ctx context.Context) []*NewUserRewardSetting {
	v, err := nurscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nurscb *NewUserRewardSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := nurscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nurscb *NewUserRewardSettingCreateBulk) ExecX(ctx context.Context) {
	if err := nurscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NewUserRewardSetting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (nurscb *NewUserRewardSettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *NewUserRewardSettingUpsertBulk {
	nurscb.conflict = opts
	return &NewUserRewardSettingUpsertBulk{
		create: nurscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NewUserRewardSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (nurscb *NewUserRewardSettingCreateBulk) OnConflictColumns(columns ...string) *NewUserRewardSettingUpsertBulk {
	nurscb.conflict = append(nurscb.conflict, sql.ConflictColumns(columns...))
	return &NewUserRewardSettingUpsertBulk{
		create: nurscb,
	}
}

// NewUserRewardSettingUpsertBulk is the builder for "upsert"-ing
// a bulk of NewUserRewardSetting nodes.
type NewUserRewardSettingUpsertBulk struct {
	create *NewUserRewardSettingCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.NewUserRewardSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *NewUserRewardSettingUpsertBulk) UpdateNewValues() *NewUserRewardSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NewUserRewardSetting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *NewUserRewardSettingUpsertBulk) Ignore() *NewUserRewardSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NewUserRewardSettingUpsertBulk) DoNothing() *NewUserRewardSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NewUserRewardSettingCreateBulk.OnConflict
// documentation for more info.
func (u *NewUserRewardSettingUpsertBulk) Update(set func(*NewUserRewardSettingUpsert)) *NewUserRewardSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NewUserRewardSettingUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *NewUserRewardSettingUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NewUserRewardSettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NewUserRewardSettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NewUserRewardSettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
