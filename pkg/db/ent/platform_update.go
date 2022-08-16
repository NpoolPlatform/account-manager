// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/platform"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// PlatformUpdate is the builder for updating Platform entities.
type PlatformUpdate struct {
	config
	hooks    []Hook
	mutation *PlatformMutation
}

// Where appends a list predicates to the PlatformUpdate builder.
func (pu *PlatformUpdate) Where(ps ...predicate.Platform) *PlatformUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PlatformUpdate) SetCreatedAt(u uint32) *PlatformUpdate {
	pu.mutation.ResetCreatedAt()
	pu.mutation.SetCreatedAt(u)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PlatformUpdate) SetNillableCreatedAt(u *uint32) *PlatformUpdate {
	if u != nil {
		pu.SetCreatedAt(*u)
	}
	return pu
}

// AddCreatedAt adds u to the "created_at" field.
func (pu *PlatformUpdate) AddCreatedAt(u int32) *PlatformUpdate {
	pu.mutation.AddCreatedAt(u)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PlatformUpdate) SetUpdatedAt(u uint32) *PlatformUpdate {
	pu.mutation.ResetUpdatedAt()
	pu.mutation.SetUpdatedAt(u)
	return pu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (pu *PlatformUpdate) AddUpdatedAt(u int32) *PlatformUpdate {
	pu.mutation.AddUpdatedAt(u)
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *PlatformUpdate) SetDeletedAt(u uint32) *PlatformUpdate {
	pu.mutation.ResetDeletedAt()
	pu.mutation.SetDeletedAt(u)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *PlatformUpdate) SetNillableDeletedAt(u *uint32) *PlatformUpdate {
	if u != nil {
		pu.SetDeletedAt(*u)
	}
	return pu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (pu *PlatformUpdate) AddDeletedAt(u int32) *PlatformUpdate {
	pu.mutation.AddDeletedAt(u)
	return pu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (pu *PlatformUpdate) SetCoinTypeID(u uuid.UUID) *PlatformUpdate {
	pu.mutation.SetCoinTypeID(u)
	return pu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (pu *PlatformUpdate) SetNillableCoinTypeID(u *uuid.UUID) *PlatformUpdate {
	if u != nil {
		pu.SetCoinTypeID(*u)
	}
	return pu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (pu *PlatformUpdate) ClearCoinTypeID() *PlatformUpdate {
	pu.mutation.ClearCoinTypeID()
	return pu
}

// SetAccountID sets the "account_id" field.
func (pu *PlatformUpdate) SetAccountID(u uuid.UUID) *PlatformUpdate {
	pu.mutation.SetAccountID(u)
	return pu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (pu *PlatformUpdate) SetNillableAccountID(u *uuid.UUID) *PlatformUpdate {
	if u != nil {
		pu.SetAccountID(*u)
	}
	return pu
}

// ClearAccountID clears the value of the "account_id" field.
func (pu *PlatformUpdate) ClearAccountID() *PlatformUpdate {
	pu.mutation.ClearAccountID()
	return pu
}

// SetUsedFor sets the "used_for" field.
func (pu *PlatformUpdate) SetUsedFor(s string) *PlatformUpdate {
	pu.mutation.SetUsedFor(s)
	return pu
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (pu *PlatformUpdate) SetNillableUsedFor(s *string) *PlatformUpdate {
	if s != nil {
		pu.SetUsedFor(*s)
	}
	return pu
}

// ClearUsedFor clears the value of the "used_for" field.
func (pu *PlatformUpdate) ClearUsedFor() *PlatformUpdate {
	pu.mutation.ClearUsedFor()
	return pu
}

// SetBackup sets the "backup" field.
func (pu *PlatformUpdate) SetBackup(b bool) *PlatformUpdate {
	pu.mutation.SetBackup(b)
	return pu
}

// SetNillableBackup sets the "backup" field if the given value is not nil.
func (pu *PlatformUpdate) SetNillableBackup(b *bool) *PlatformUpdate {
	if b != nil {
		pu.SetBackup(*b)
	}
	return pu
}

// ClearBackup clears the value of the "backup" field.
func (pu *PlatformUpdate) ClearBackup() *PlatformUpdate {
	pu.mutation.ClearBackup()
	return pu
}

// Mutation returns the PlatformMutation object of the builder.
func (pu *PlatformUpdate) Mutation() *PlatformMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlatformUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlatformMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlatformUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlatformUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlatformUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PlatformUpdate) defaults() error {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		if platform.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized platform.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := platform.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (pu *PlatformUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   platform.Table,
			Columns: platform.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: platform.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldDeletedAt,
		})
	}
	if value, ok := pu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldDeletedAt,
		})
	}
	if value, ok := pu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: platform.FieldCoinTypeID,
		})
	}
	if pu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: platform.FieldCoinTypeID,
		})
	}
	if value, ok := pu.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: platform.FieldAccountID,
		})
	}
	if pu.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: platform.FieldAccountID,
		})
	}
	if value, ok := pu.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: platform.FieldUsedFor,
		})
	}
	if pu.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: platform.FieldUsedFor,
		})
	}
	if value, ok := pu.mutation.Backup(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: platform.FieldBackup,
		})
	}
	if pu.mutation.BackupCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: platform.FieldBackup,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{platform.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PlatformUpdateOne is the builder for updating a single Platform entity.
type PlatformUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlatformMutation
}

// SetCreatedAt sets the "created_at" field.
func (puo *PlatformUpdateOne) SetCreatedAt(u uint32) *PlatformUpdateOne {
	puo.mutation.ResetCreatedAt()
	puo.mutation.SetCreatedAt(u)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PlatformUpdateOne) SetNillableCreatedAt(u *uint32) *PlatformUpdateOne {
	if u != nil {
		puo.SetCreatedAt(*u)
	}
	return puo
}

// AddCreatedAt adds u to the "created_at" field.
func (puo *PlatformUpdateOne) AddCreatedAt(u int32) *PlatformUpdateOne {
	puo.mutation.AddCreatedAt(u)
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PlatformUpdateOne) SetUpdatedAt(u uint32) *PlatformUpdateOne {
	puo.mutation.ResetUpdatedAt()
	puo.mutation.SetUpdatedAt(u)
	return puo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (puo *PlatformUpdateOne) AddUpdatedAt(u int32) *PlatformUpdateOne {
	puo.mutation.AddUpdatedAt(u)
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *PlatformUpdateOne) SetDeletedAt(u uint32) *PlatformUpdateOne {
	puo.mutation.ResetDeletedAt()
	puo.mutation.SetDeletedAt(u)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *PlatformUpdateOne) SetNillableDeletedAt(u *uint32) *PlatformUpdateOne {
	if u != nil {
		puo.SetDeletedAt(*u)
	}
	return puo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (puo *PlatformUpdateOne) AddDeletedAt(u int32) *PlatformUpdateOne {
	puo.mutation.AddDeletedAt(u)
	return puo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (puo *PlatformUpdateOne) SetCoinTypeID(u uuid.UUID) *PlatformUpdateOne {
	puo.mutation.SetCoinTypeID(u)
	return puo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (puo *PlatformUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *PlatformUpdateOne {
	if u != nil {
		puo.SetCoinTypeID(*u)
	}
	return puo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (puo *PlatformUpdateOne) ClearCoinTypeID() *PlatformUpdateOne {
	puo.mutation.ClearCoinTypeID()
	return puo
}

// SetAccountID sets the "account_id" field.
func (puo *PlatformUpdateOne) SetAccountID(u uuid.UUID) *PlatformUpdateOne {
	puo.mutation.SetAccountID(u)
	return puo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (puo *PlatformUpdateOne) SetNillableAccountID(u *uuid.UUID) *PlatformUpdateOne {
	if u != nil {
		puo.SetAccountID(*u)
	}
	return puo
}

// ClearAccountID clears the value of the "account_id" field.
func (puo *PlatformUpdateOne) ClearAccountID() *PlatformUpdateOne {
	puo.mutation.ClearAccountID()
	return puo
}

// SetUsedFor sets the "used_for" field.
func (puo *PlatformUpdateOne) SetUsedFor(s string) *PlatformUpdateOne {
	puo.mutation.SetUsedFor(s)
	return puo
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (puo *PlatformUpdateOne) SetNillableUsedFor(s *string) *PlatformUpdateOne {
	if s != nil {
		puo.SetUsedFor(*s)
	}
	return puo
}

// ClearUsedFor clears the value of the "used_for" field.
func (puo *PlatformUpdateOne) ClearUsedFor() *PlatformUpdateOne {
	puo.mutation.ClearUsedFor()
	return puo
}

// SetBackup sets the "backup" field.
func (puo *PlatformUpdateOne) SetBackup(b bool) *PlatformUpdateOne {
	puo.mutation.SetBackup(b)
	return puo
}

// SetNillableBackup sets the "backup" field if the given value is not nil.
func (puo *PlatformUpdateOne) SetNillableBackup(b *bool) *PlatformUpdateOne {
	if b != nil {
		puo.SetBackup(*b)
	}
	return puo
}

// ClearBackup clears the value of the "backup" field.
func (puo *PlatformUpdateOne) ClearBackup() *PlatformUpdateOne {
	puo.mutation.ClearBackup()
	return puo
}

// Mutation returns the PlatformMutation object of the builder.
func (puo *PlatformUpdateOne) Mutation() *PlatformMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlatformUpdateOne) Select(field string, fields ...string) *PlatformUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Platform entity.
func (puo *PlatformUpdateOne) Save(ctx context.Context) (*Platform, error) {
	var (
		err  error
		node *Platform
	)
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlatformMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlatformUpdateOne) SaveX(ctx context.Context) *Platform {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlatformUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlatformUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PlatformUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		if platform.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized platform.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := platform.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (puo *PlatformUpdateOne) sqlSave(ctx context.Context) (_node *Platform, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   platform.Table,
			Columns: platform.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: platform.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Platform.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, platform.FieldID)
		for _, f := range fields {
			if !platform.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != platform.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldDeletedAt,
		})
	}
	if value, ok := puo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platform.FieldDeletedAt,
		})
	}
	if value, ok := puo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: platform.FieldCoinTypeID,
		})
	}
	if puo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: platform.FieldCoinTypeID,
		})
	}
	if value, ok := puo.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: platform.FieldAccountID,
		})
	}
	if puo.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: platform.FieldAccountID,
		})
	}
	if value, ok := puo.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: platform.FieldUsedFor,
		})
	}
	if puo.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: platform.FieldUsedFor,
		})
	}
	if value, ok := puo.mutation.Backup(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: platform.FieldBackup,
		})
	}
	if puo.mutation.BackupCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: platform.FieldBackup,
		})
	}
	_node = &Platform{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{platform.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}