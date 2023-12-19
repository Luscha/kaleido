// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.pitagora/pkg/ent/automation"
)

// AutomationCreate is the builder for creating a Automation entity.
type AutomationCreate struct {
	config
	mutation *AutomationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ac *AutomationCreate) SetName(s string) *AutomationCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *AutomationCreate) SetDescription(s string) *AutomationCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetTrigger sets the "trigger" field.
func (ac *AutomationCreate) SetTrigger(s string) *AutomationCreate {
	ac.mutation.SetTrigger(s)
	return ac
}

// SetManifest sets the "manifest" field.
func (ac *AutomationCreate) SetManifest(s string) *AutomationCreate {
	ac.mutation.SetManifest(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AutomationCreate) SetID(i int64) *AutomationCreate {
	ac.mutation.SetID(i)
	return ac
}

// Mutation returns the AutomationMutation object of the builder.
func (ac *AutomationCreate) Mutation() *AutomationMutation {
	return ac.mutation
}

// Save creates the Automation in the database.
func (ac *AutomationCreate) Save(ctx context.Context) (*Automation, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AutomationCreate) SaveX(ctx context.Context) *Automation {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AutomationCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AutomationCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AutomationCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Automation.name"`)}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Automation.description"`)}
	}
	if _, ok := ac.mutation.Trigger(); !ok {
		return &ValidationError{Name: "trigger", err: errors.New(`ent: missing required field "Automation.trigger"`)}
	}
	if _, ok := ac.mutation.Manifest(); !ok {
		return &ValidationError{Name: "manifest", err: errors.New(`ent: missing required field "Automation.manifest"`)}
	}
	return nil
}

func (ac *AutomationCreate) sqlSave(ctx context.Context) (*Automation, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AutomationCreate) createSpec() (*Automation, *sqlgraph.CreateSpec) {
	var (
		_node = &Automation{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(automation.Table, sqlgraph.NewFieldSpec(automation.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(automation.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(automation.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.Trigger(); ok {
		_spec.SetField(automation.FieldTrigger, field.TypeString, value)
		_node.Trigger = value
	}
	if value, ok := ac.mutation.Manifest(); ok {
		_spec.SetField(automation.FieldManifest, field.TypeString, value)
		_node.Manifest = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Automation.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AutomationUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ac *AutomationCreate) OnConflict(opts ...sql.ConflictOption) *AutomationUpsertOne {
	ac.conflict = opts
	return &AutomationUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Automation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AutomationCreate) OnConflictColumns(columns ...string) *AutomationUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AutomationUpsertOne{
		create: ac,
	}
}

type (
	// AutomationUpsertOne is the builder for "upsert"-ing
	//  one Automation node.
	AutomationUpsertOne struct {
		create *AutomationCreate
	}

	// AutomationUpsert is the "OnConflict" setter.
	AutomationUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *AutomationUpsert) SetDescription(v string) *AutomationUpsert {
	u.Set(automation.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AutomationUpsert) UpdateDescription() *AutomationUpsert {
	u.SetExcluded(automation.FieldDescription)
	return u
}

// SetTrigger sets the "trigger" field.
func (u *AutomationUpsert) SetTrigger(v string) *AutomationUpsert {
	u.Set(automation.FieldTrigger, v)
	return u
}

// UpdateTrigger sets the "trigger" field to the value that was provided on create.
func (u *AutomationUpsert) UpdateTrigger() *AutomationUpsert {
	u.SetExcluded(automation.FieldTrigger)
	return u
}

// SetManifest sets the "manifest" field.
func (u *AutomationUpsert) SetManifest(v string) *AutomationUpsert {
	u.Set(automation.FieldManifest, v)
	return u
}

// UpdateManifest sets the "manifest" field to the value that was provided on create.
func (u *AutomationUpsert) UpdateManifest() *AutomationUpsert {
	u.SetExcluded(automation.FieldManifest)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Automation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(automation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AutomationUpsertOne) UpdateNewValues() *AutomationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(automation.FieldID)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(automation.FieldName)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Automation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AutomationUpsertOne) Ignore() *AutomationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AutomationUpsertOne) DoNothing() *AutomationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AutomationCreate.OnConflict
// documentation for more info.
func (u *AutomationUpsertOne) Update(set func(*AutomationUpsert)) *AutomationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AutomationUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *AutomationUpsertOne) SetDescription(v string) *AutomationUpsertOne {
	return u.Update(func(s *AutomationUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AutomationUpsertOne) UpdateDescription() *AutomationUpsertOne {
	return u.Update(func(s *AutomationUpsert) {
		s.UpdateDescription()
	})
}

// SetTrigger sets the "trigger" field.
func (u *AutomationUpsertOne) SetTrigger(v string) *AutomationUpsertOne {
	return u.Update(func(s *AutomationUpsert) {
		s.SetTrigger(v)
	})
}

// UpdateTrigger sets the "trigger" field to the value that was provided on create.
func (u *AutomationUpsertOne) UpdateTrigger() *AutomationUpsertOne {
	return u.Update(func(s *AutomationUpsert) {
		s.UpdateTrigger()
	})
}

// SetManifest sets the "manifest" field.
func (u *AutomationUpsertOne) SetManifest(v string) *AutomationUpsertOne {
	return u.Update(func(s *AutomationUpsert) {
		s.SetManifest(v)
	})
}

// UpdateManifest sets the "manifest" field to the value that was provided on create.
func (u *AutomationUpsertOne) UpdateManifest() *AutomationUpsertOne {
	return u.Update(func(s *AutomationUpsert) {
		s.UpdateManifest()
	})
}

// Exec executes the query.
func (u *AutomationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AutomationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AutomationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AutomationUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AutomationUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AutomationCreateBulk is the builder for creating many Automation entities in bulk.
type AutomationCreateBulk struct {
	config
	err      error
	builders []*AutomationCreate
	conflict []sql.ConflictOption
}

// Save creates the Automation entities in the database.
func (acb *AutomationCreateBulk) Save(ctx context.Context) ([]*Automation, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Automation, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AutomationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AutomationCreateBulk) SaveX(ctx context.Context) []*Automation {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AutomationCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AutomationCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Automation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AutomationUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (acb *AutomationCreateBulk) OnConflict(opts ...sql.ConflictOption) *AutomationUpsertBulk {
	acb.conflict = opts
	return &AutomationUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Automation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AutomationCreateBulk) OnConflictColumns(columns ...string) *AutomationUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AutomationUpsertBulk{
		create: acb,
	}
}

// AutomationUpsertBulk is the builder for "upsert"-ing
// a bulk of Automation nodes.
type AutomationUpsertBulk struct {
	create *AutomationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Automation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(automation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AutomationUpsertBulk) UpdateNewValues() *AutomationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(automation.FieldID)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(automation.FieldName)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Automation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AutomationUpsertBulk) Ignore() *AutomationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AutomationUpsertBulk) DoNothing() *AutomationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AutomationCreateBulk.OnConflict
// documentation for more info.
func (u *AutomationUpsertBulk) Update(set func(*AutomationUpsert)) *AutomationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AutomationUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *AutomationUpsertBulk) SetDescription(v string) *AutomationUpsertBulk {
	return u.Update(func(s *AutomationUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AutomationUpsertBulk) UpdateDescription() *AutomationUpsertBulk {
	return u.Update(func(s *AutomationUpsert) {
		s.UpdateDescription()
	})
}

// SetTrigger sets the "trigger" field.
func (u *AutomationUpsertBulk) SetTrigger(v string) *AutomationUpsertBulk {
	return u.Update(func(s *AutomationUpsert) {
		s.SetTrigger(v)
	})
}

// UpdateTrigger sets the "trigger" field to the value that was provided on create.
func (u *AutomationUpsertBulk) UpdateTrigger() *AutomationUpsertBulk {
	return u.Update(func(s *AutomationUpsert) {
		s.UpdateTrigger()
	})
}

// SetManifest sets the "manifest" field.
func (u *AutomationUpsertBulk) SetManifest(v string) *AutomationUpsertBulk {
	return u.Update(func(s *AutomationUpsert) {
		s.SetManifest(v)
	})
}

// UpdateManifest sets the "manifest" field to the value that was provided on create.
func (u *AutomationUpsertBulk) UpdateManifest() *AutomationUpsertBulk {
	return u.Update(func(s *AutomationUpsert) {
		s.UpdateManifest()
	})
}

// Exec executes the query.
func (u *AutomationUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AutomationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AutomationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AutomationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}