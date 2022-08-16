// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AccountUpdate) SetCreatedAt(u uint32) *AccountUpdate {
	au.mutation.ResetCreatedAt()
	au.mutation.SetCreatedAt(u)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableCreatedAt(u *uint32) *AccountUpdate {
	if u != nil {
		au.SetCreatedAt(*u)
	}
	return au
}

// AddCreatedAt adds u to the "created_at" field.
func (au *AccountUpdate) AddCreatedAt(u int32) *AccountUpdate {
	au.mutation.AddCreatedAt(u)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(u uint32) *AccountUpdate {
	au.mutation.ResetUpdatedAt()
	au.mutation.SetUpdatedAt(u)
	return au
}

// AddUpdatedAt adds u to the "updated_at" field.
func (au *AccountUpdate) AddUpdatedAt(u int32) *AccountUpdate {
	au.mutation.AddUpdatedAt(u)
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AccountUpdate) SetDeletedAt(u uint32) *AccountUpdate {
	au.mutation.ResetDeletedAt()
	au.mutation.SetDeletedAt(u)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableDeletedAt(u *uint32) *AccountUpdate {
	if u != nil {
		au.SetDeletedAt(*u)
	}
	return au
}

// AddDeletedAt adds u to the "deleted_at" field.
func (au *AccountUpdate) AddDeletedAt(u int32) *AccountUpdate {
	au.mutation.AddDeletedAt(u)
	return au
}

// SetCoinTypeID sets the "coin_type_id" field.
func (au *AccountUpdate) SetCoinTypeID(u uuid.UUID) *AccountUpdate {
	au.mutation.SetCoinTypeID(u)
	return au
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (au *AccountUpdate) SetNillableCoinTypeID(u *uuid.UUID) *AccountUpdate {
	if u != nil {
		au.SetCoinTypeID(*u)
	}
	return au
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (au *AccountUpdate) ClearCoinTypeID() *AccountUpdate {
	au.mutation.ClearCoinTypeID()
	return au
}

// SetAddress sets the "address" field.
func (au *AccountUpdate) SetAddress(s string) *AccountUpdate {
	au.mutation.SetAddress(s)
	return au
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (au *AccountUpdate) SetNillableAddress(s *string) *AccountUpdate {
	if s != nil {
		au.SetAddress(*s)
	}
	return au
}

// ClearAddress clears the value of the "address" field.
func (au *AccountUpdate) ClearAddress() *AccountUpdate {
	au.mutation.ClearAddress()
	return au
}

// SetUsedFor sets the "used_for" field.
func (au *AccountUpdate) SetUsedFor(s string) *AccountUpdate {
	au.mutation.SetUsedFor(s)
	return au
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (au *AccountUpdate) SetNillableUsedFor(s *string) *AccountUpdate {
	if s != nil {
		au.SetUsedFor(*s)
	}
	return au
}

// ClearUsedFor clears the value of the "used_for" field.
func (au *AccountUpdate) ClearUsedFor() *AccountUpdate {
	au.mutation.ClearUsedFor()
	return au
}

// SetPlatformHoldPrivateKey sets the "platform_hold_private_key" field.
func (au *AccountUpdate) SetPlatformHoldPrivateKey(b bool) *AccountUpdate {
	au.mutation.SetPlatformHoldPrivateKey(b)
	return au
}

// SetNillablePlatformHoldPrivateKey sets the "platform_hold_private_key" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePlatformHoldPrivateKey(b *bool) *AccountUpdate {
	if b != nil {
		au.SetPlatformHoldPrivateKey(*b)
	}
	return au
}

// ClearPlatformHoldPrivateKey clears the value of the "platform_hold_private_key" field.
func (au *AccountUpdate) ClearPlatformHoldPrivateKey() *AccountUpdate {
	au.mutation.ClearPlatformHoldPrivateKey()
	return au
}

// SetActive sets the "active" field.
func (au *AccountUpdate) SetActive(b bool) *AccountUpdate {
	au.mutation.SetActive(b)
	return au
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (au *AccountUpdate) SetNillableActive(b *bool) *AccountUpdate {
	if b != nil {
		au.SetActive(*b)
	}
	return au
}

// ClearActive clears the value of the "active" field.
func (au *AccountUpdate) ClearActive() *AccountUpdate {
	au.mutation.ClearActive()
	return au
}

// SetLocked sets the "locked" field.
func (au *AccountUpdate) SetLocked(b bool) *AccountUpdate {
	au.mutation.SetLocked(b)
	return au
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (au *AccountUpdate) SetNillableLocked(b *bool) *AccountUpdate {
	if b != nil {
		au.SetLocked(*b)
	}
	return au
}

// ClearLocked clears the value of the "locked" field.
func (au *AccountUpdate) ClearLocked() *AccountUpdate {
	au.mutation.ClearLocked()
	return au
}

// SetLockedBy sets the "locked_by" field.
func (au *AccountUpdate) SetLockedBy(s string) *AccountUpdate {
	au.mutation.SetLockedBy(s)
	return au
}

// SetNillableLockedBy sets the "locked_by" field if the given value is not nil.
func (au *AccountUpdate) SetNillableLockedBy(s *string) *AccountUpdate {
	if s != nil {
		au.SetLockedBy(*s)
	}
	return au
}

// ClearLockedBy clears the value of the "locked_by" field.
func (au *AccountUpdate) ClearLockedBy() *AccountUpdate {
	au.mutation.ClearLockedBy()
	return au
}

// SetBlocked sets the "blocked" field.
func (au *AccountUpdate) SetBlocked(b bool) *AccountUpdate {
	au.mutation.SetBlocked(b)
	return au
}

// SetNillableBlocked sets the "blocked" field if the given value is not nil.
func (au *AccountUpdate) SetNillableBlocked(b *bool) *AccountUpdate {
	if b != nil {
		au.SetBlocked(*b)
	}
	return au
}

// ClearBlocked clears the value of the "blocked" field.
func (au *AccountUpdate) ClearBlocked() *AccountUpdate {
	au.mutation.ClearBlocked()
	return au
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := au.defaults(); err != nil {
		return 0, err
	}
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() error {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		if account.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized account.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldDeletedAt,
		})
	}
	if value, ok := au.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldDeletedAt,
		})
	}
	if value, ok := au.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: account.FieldCoinTypeID,
		})
	}
	if au.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: account.FieldCoinTypeID,
		})
	}
	if value, ok := au.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAddress,
		})
	}
	if au.mutation.AddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldAddress,
		})
	}
	if value, ok := au.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldUsedFor,
		})
	}
	if au.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldUsedFor,
		})
	}
	if value, ok := au.mutation.PlatformHoldPrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldPlatformHoldPrivateKey,
		})
	}
	if au.mutation.PlatformHoldPrivateKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldPlatformHoldPrivateKey,
		})
	}
	if value, ok := au.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldActive,
		})
	}
	if au.mutation.ActiveCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldActive,
		})
	}
	if value, ok := au.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldLocked,
		})
	}
	if au.mutation.LockedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldLocked,
		})
	}
	if value, ok := au.mutation.LockedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldLockedBy,
		})
	}
	if au.mutation.LockedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldLockedBy,
		})
	}
	if value, ok := au.mutation.Blocked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldBlocked,
		})
	}
	if au.mutation.BlockedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldBlocked,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetCreatedAt sets the "created_at" field.
func (auo *AccountUpdateOne) SetCreatedAt(u uint32) *AccountUpdateOne {
	auo.mutation.ResetCreatedAt()
	auo.mutation.SetCreatedAt(u)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableCreatedAt(u *uint32) *AccountUpdateOne {
	if u != nil {
		auo.SetCreatedAt(*u)
	}
	return auo
}

// AddCreatedAt adds u to the "created_at" field.
func (auo *AccountUpdateOne) AddCreatedAt(u int32) *AccountUpdateOne {
	auo.mutation.AddCreatedAt(u)
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(u uint32) *AccountUpdateOne {
	auo.mutation.ResetUpdatedAt()
	auo.mutation.SetUpdatedAt(u)
	return auo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (auo *AccountUpdateOne) AddUpdatedAt(u int32) *AccountUpdateOne {
	auo.mutation.AddUpdatedAt(u)
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AccountUpdateOne) SetDeletedAt(u uint32) *AccountUpdateOne {
	auo.mutation.ResetDeletedAt()
	auo.mutation.SetDeletedAt(u)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableDeletedAt(u *uint32) *AccountUpdateOne {
	if u != nil {
		auo.SetDeletedAt(*u)
	}
	return auo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (auo *AccountUpdateOne) AddDeletedAt(u int32) *AccountUpdateOne {
	auo.mutation.AddDeletedAt(u)
	return auo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (auo *AccountUpdateOne) SetCoinTypeID(u uuid.UUID) *AccountUpdateOne {
	auo.mutation.SetCoinTypeID(u)
	return auo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *AccountUpdateOne {
	if u != nil {
		auo.SetCoinTypeID(*u)
	}
	return auo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (auo *AccountUpdateOne) ClearCoinTypeID() *AccountUpdateOne {
	auo.mutation.ClearCoinTypeID()
	return auo
}

// SetAddress sets the "address" field.
func (auo *AccountUpdateOne) SetAddress(s string) *AccountUpdateOne {
	auo.mutation.SetAddress(s)
	return auo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableAddress(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetAddress(*s)
	}
	return auo
}

// ClearAddress clears the value of the "address" field.
func (auo *AccountUpdateOne) ClearAddress() *AccountUpdateOne {
	auo.mutation.ClearAddress()
	return auo
}

// SetUsedFor sets the "used_for" field.
func (auo *AccountUpdateOne) SetUsedFor(s string) *AccountUpdateOne {
	auo.mutation.SetUsedFor(s)
	return auo
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableUsedFor(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetUsedFor(*s)
	}
	return auo
}

// ClearUsedFor clears the value of the "used_for" field.
func (auo *AccountUpdateOne) ClearUsedFor() *AccountUpdateOne {
	auo.mutation.ClearUsedFor()
	return auo
}

// SetPlatformHoldPrivateKey sets the "platform_hold_private_key" field.
func (auo *AccountUpdateOne) SetPlatformHoldPrivateKey(b bool) *AccountUpdateOne {
	auo.mutation.SetPlatformHoldPrivateKey(b)
	return auo
}

// SetNillablePlatformHoldPrivateKey sets the "platform_hold_private_key" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePlatformHoldPrivateKey(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetPlatformHoldPrivateKey(*b)
	}
	return auo
}

// ClearPlatformHoldPrivateKey clears the value of the "platform_hold_private_key" field.
func (auo *AccountUpdateOne) ClearPlatformHoldPrivateKey() *AccountUpdateOne {
	auo.mutation.ClearPlatformHoldPrivateKey()
	return auo
}

// SetActive sets the "active" field.
func (auo *AccountUpdateOne) SetActive(b bool) *AccountUpdateOne {
	auo.mutation.SetActive(b)
	return auo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableActive(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetActive(*b)
	}
	return auo
}

// ClearActive clears the value of the "active" field.
func (auo *AccountUpdateOne) ClearActive() *AccountUpdateOne {
	auo.mutation.ClearActive()
	return auo
}

// SetLocked sets the "locked" field.
func (auo *AccountUpdateOne) SetLocked(b bool) *AccountUpdateOne {
	auo.mutation.SetLocked(b)
	return auo
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableLocked(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetLocked(*b)
	}
	return auo
}

// ClearLocked clears the value of the "locked" field.
func (auo *AccountUpdateOne) ClearLocked() *AccountUpdateOne {
	auo.mutation.ClearLocked()
	return auo
}

// SetLockedBy sets the "locked_by" field.
func (auo *AccountUpdateOne) SetLockedBy(s string) *AccountUpdateOne {
	auo.mutation.SetLockedBy(s)
	return auo
}

// SetNillableLockedBy sets the "locked_by" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableLockedBy(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetLockedBy(*s)
	}
	return auo
}

// ClearLockedBy clears the value of the "locked_by" field.
func (auo *AccountUpdateOne) ClearLockedBy() *AccountUpdateOne {
	auo.mutation.ClearLockedBy()
	return auo
}

// SetBlocked sets the "blocked" field.
func (auo *AccountUpdateOne) SetBlocked(b bool) *AccountUpdateOne {
	auo.mutation.SetBlocked(b)
	return auo
}

// SetNillableBlocked sets the "blocked" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableBlocked(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetBlocked(*b)
	}
	return auo
}

// ClearBlocked clears the value of the "blocked" field.
func (auo *AccountUpdateOne) ClearBlocked() *AccountUpdateOne {
	auo.mutation.ClearBlocked()
	return auo
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	if err := auo.defaults(); err != nil {
		return nil, err
	}
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() error {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		if account.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized account.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldDeletedAt,
		})
	}
	if value, ok := auo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: account.FieldDeletedAt,
		})
	}
	if value, ok := auo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: account.FieldCoinTypeID,
		})
	}
	if auo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: account.FieldCoinTypeID,
		})
	}
	if value, ok := auo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAddress,
		})
	}
	if auo.mutation.AddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldAddress,
		})
	}
	if value, ok := auo.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldUsedFor,
		})
	}
	if auo.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldUsedFor,
		})
	}
	if value, ok := auo.mutation.PlatformHoldPrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldPlatformHoldPrivateKey,
		})
	}
	if auo.mutation.PlatformHoldPrivateKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldPlatformHoldPrivateKey,
		})
	}
	if value, ok := auo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldActive,
		})
	}
	if auo.mutation.ActiveCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldActive,
		})
	}
	if value, ok := auo.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldLocked,
		})
	}
	if auo.mutation.LockedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldLocked,
		})
	}
	if value, ok := auo.mutation.LockedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldLockedBy,
		})
	}
	if auo.mutation.LockedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldLockedBy,
		})
	}
	if value, ok := auo.mutation.Blocked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldBlocked,
		})
	}
	if auo.mutation.BlockedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldBlocked,
		})
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
