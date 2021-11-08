// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-stimulate/pkg/db/ent/registrationinvitation"
)

// RegistrationInvitationCreate is the builder for creating a RegistrationInvitation entity.
type RegistrationInvitationCreate struct {
	config
	mutation *RegistrationInvitationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// Mutation returns the RegistrationInvitationMutation object of the builder.
func (ric *RegistrationInvitationCreate) Mutation() *RegistrationInvitationMutation {
	return ric.mutation
}

// Save creates the RegistrationInvitation in the database.
func (ric *RegistrationInvitationCreate) Save(ctx context.Context) (*RegistrationInvitation, error) {
	var (
		err  error
		node *RegistrationInvitation
	)
	if len(ric.hooks) == 0 {
		if err = ric.check(); err != nil {
			return nil, err
		}
		node, err = ric.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegistrationInvitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ric.check(); err != nil {
				return nil, err
			}
			ric.mutation = mutation
			if node, err = ric.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ric.hooks) - 1; i >= 0; i-- {
			if ric.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ric.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ric.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ric *RegistrationInvitationCreate) SaveX(ctx context.Context) *RegistrationInvitation {
	v, err := ric.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ric *RegistrationInvitationCreate) Exec(ctx context.Context) error {
	_, err := ric.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ric *RegistrationInvitationCreate) ExecX(ctx context.Context) {
	if err := ric.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ric *RegistrationInvitationCreate) check() error {
	return nil
}

func (ric *RegistrationInvitationCreate) sqlSave(ctx context.Context) (*RegistrationInvitation, error) {
	_node, _spec := ric.createSpec()
	if err := sqlgraph.CreateNode(ctx, ric.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ric *RegistrationInvitationCreate) createSpec() (*RegistrationInvitation, *sqlgraph.CreateSpec) {
	var (
		_node = &RegistrationInvitation{config: ric.config}
		_spec = &sqlgraph.CreateSpec{
			Table: registrationinvitation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: registrationinvitation.FieldID,
			},
		}
	)
	_spec.OnConflict = ric.conflict
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.RegistrationInvitation.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (ric *RegistrationInvitationCreate) OnConflict(opts ...sql.ConflictOption) *RegistrationInvitationUpsertOne {
	ric.conflict = opts
	return &RegistrationInvitationUpsertOne{
		create: ric,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.RegistrationInvitation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ric *RegistrationInvitationCreate) OnConflictColumns(columns ...string) *RegistrationInvitationUpsertOne {
	ric.conflict = append(ric.conflict, sql.ConflictColumns(columns...))
	return &RegistrationInvitationUpsertOne{
		create: ric,
	}
}

type (
	// RegistrationInvitationUpsertOne is the builder for "upsert"-ing
	//  one RegistrationInvitation node.
	RegistrationInvitationUpsertOne struct {
		create *RegistrationInvitationCreate
	}

	// RegistrationInvitationUpsert is the "OnConflict" setter.
	RegistrationInvitationUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.RegistrationInvitation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *RegistrationInvitationUpsertOne) UpdateNewValues() *RegistrationInvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.RegistrationInvitation.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *RegistrationInvitationUpsertOne) Ignore() *RegistrationInvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RegistrationInvitationUpsertOne) DoNothing() *RegistrationInvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RegistrationInvitationCreate.OnConflict
// documentation for more info.
func (u *RegistrationInvitationUpsertOne) Update(set func(*RegistrationInvitationUpsert)) *RegistrationInvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RegistrationInvitationUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *RegistrationInvitationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RegistrationInvitationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RegistrationInvitationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RegistrationInvitationUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RegistrationInvitationUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RegistrationInvitationCreateBulk is the builder for creating many RegistrationInvitation entities in bulk.
type RegistrationInvitationCreateBulk struct {
	config
	builders []*RegistrationInvitationCreate
	conflict []sql.ConflictOption
}

// Save creates the RegistrationInvitation entities in the database.
func (ricb *RegistrationInvitationCreateBulk) Save(ctx context.Context) ([]*RegistrationInvitation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ricb.builders))
	nodes := make([]*RegistrationInvitation, len(ricb.builders))
	mutators := make([]Mutator, len(ricb.builders))
	for i := range ricb.builders {
		func(i int, root context.Context) {
			builder := ricb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RegistrationInvitationMutation)
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
					_, err = mutators[i+1].Mutate(root, ricb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ricb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ricb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ricb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ricb *RegistrationInvitationCreateBulk) SaveX(ctx context.Context) []*RegistrationInvitation {
	v, err := ricb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ricb *RegistrationInvitationCreateBulk) Exec(ctx context.Context) error {
	_, err := ricb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ricb *RegistrationInvitationCreateBulk) ExecX(ctx context.Context) {
	if err := ricb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.RegistrationInvitation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (ricb *RegistrationInvitationCreateBulk) OnConflict(opts ...sql.ConflictOption) *RegistrationInvitationUpsertBulk {
	ricb.conflict = opts
	return &RegistrationInvitationUpsertBulk{
		create: ricb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.RegistrationInvitation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ricb *RegistrationInvitationCreateBulk) OnConflictColumns(columns ...string) *RegistrationInvitationUpsertBulk {
	ricb.conflict = append(ricb.conflict, sql.ConflictColumns(columns...))
	return &RegistrationInvitationUpsertBulk{
		create: ricb,
	}
}

// RegistrationInvitationUpsertBulk is the builder for "upsert"-ing
// a bulk of RegistrationInvitation nodes.
type RegistrationInvitationUpsertBulk struct {
	create *RegistrationInvitationCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.RegistrationInvitation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *RegistrationInvitationUpsertBulk) UpdateNewValues() *RegistrationInvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.RegistrationInvitation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *RegistrationInvitationUpsertBulk) Ignore() *RegistrationInvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RegistrationInvitationUpsertBulk) DoNothing() *RegistrationInvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RegistrationInvitationCreateBulk.OnConflict
// documentation for more info.
func (u *RegistrationInvitationUpsertBulk) Update(set func(*RegistrationInvitationUpsert)) *RegistrationInvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RegistrationInvitationUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *RegistrationInvitationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RegistrationInvitationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RegistrationInvitationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RegistrationInvitationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
