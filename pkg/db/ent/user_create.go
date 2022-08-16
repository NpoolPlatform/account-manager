// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/user"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(u uint32) *UserCreate {
	uc.mutation.SetCreatedAt(u)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(u *uint32) *UserCreate {
	if u != nil {
		uc.SetCreatedAt(*u)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(u uint32) *UserCreate {
	uc.mutation.SetUpdatedAt(u)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(u *uint32) *UserCreate {
	if u != nil {
		uc.SetUpdatedAt(*u)
	}
	return uc
}

// SetDeletedAt sets the "deleted_at" field.
func (uc *UserCreate) SetDeletedAt(u uint32) *UserCreate {
	uc.mutation.SetDeletedAt(u)
	return uc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeletedAt(u *uint32) *UserCreate {
	if u != nil {
		uc.SetDeletedAt(*u)
	}
	return uc
}

// SetAppID sets the "app_id" field.
func (uc *UserCreate) SetAppID(u uuid.UUID) *UserCreate {
	uc.mutation.SetAppID(u)
	return uc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableAppID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetAppID(*u)
	}
	return uc
}

// SetUserID sets the "user_id" field.
func (uc *UserCreate) SetUserID(u uuid.UUID) *UserCreate {
	uc.mutation.SetUserID(u)
	return uc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableUserID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetUserID(*u)
	}
	return uc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (uc *UserCreate) SetCoinTypeID(u uuid.UUID) *UserCreate {
	uc.mutation.SetCoinTypeID(u)
	return uc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableCoinTypeID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetCoinTypeID(*u)
	}
	return uc
}

// SetAccountID sets the "account_id" field.
func (uc *UserCreate) SetAccountID(u uuid.UUID) *UserCreate {
	uc.mutation.SetAccountID(u)
	return uc
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableAccountID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetAccountID(*u)
	}
	return uc
}

// SetUsedFor sets the "used_for" field.
func (uc *UserCreate) SetUsedFor(s string) *UserCreate {
	uc.mutation.SetUsedFor(s)
	return uc
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (uc *UserCreate) SetNillableUsedFor(s *string) *UserCreate {
	if s != nil {
		uc.SetUsedFor(*s)
	}
	return uc
}

// SetLabels sets the "labels" field.
func (uc *UserCreate) SetLabels(s []string) *UserCreate {
	uc.mutation.SetLabels(s)
	return uc
}

// SetBalance sets the "balance" field.
func (uc *UserCreate) SetBalance(d decimal.Decimal) *UserCreate {
	uc.mutation.SetBalance(d)
	return uc
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (uc *UserCreate) SetNillableBalance(d *decimal.Decimal) *UserCreate {
	if d != nil {
		uc.SetBalance(*d)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if err := uc.defaults(); err != nil {
		return nil, err
	}
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		if user.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		if user.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.DeletedAt(); !ok {
		if user.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := user.DefaultDeletedAt()
		uc.mutation.SetDeletedAt(v)
	}
	if _, ok := uc.mutation.AppID(); !ok {
		if user.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := user.DefaultAppID()
		uc.mutation.SetAppID(v)
	}
	if _, ok := uc.mutation.UserID(); !ok {
		if user.DefaultUserID == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultUserID (forgotten import ent/runtime?)")
		}
		v := user.DefaultUserID()
		uc.mutation.SetUserID(v)
	}
	if _, ok := uc.mutation.CoinTypeID(); !ok {
		if user.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := user.DefaultCoinTypeID()
		uc.mutation.SetCoinTypeID(v)
	}
	if _, ok := uc.mutation.AccountID(); !ok {
		if user.DefaultAccountID == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultAccountID (forgotten import ent/runtime?)")
		}
		v := user.DefaultAccountID()
		uc.mutation.SetAccountID(v)
	}
	if _, ok := uc.mutation.UsedFor(); !ok {
		v := user.DefaultUsedFor
		uc.mutation.SetUsedFor(v)
	}
	if _, ok := uc.mutation.Labels(); !ok {
		v := user.DefaultLabels
		uc.mutation.SetLabels(v)
	}
	if _, ok := uc.mutation.Balance(); !ok {
		v := user.DefaultBalance
		uc.mutation.SetBalance(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		if user.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultID (forgotten import ent/runtime?)")
		}
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "User.deleted_at"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		}
	)
	_spec.OnConflict = uc.conflict
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := uc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := uc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := uc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := uc.mutation.AccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldAccountID,
		})
		_node.AccountID = value
	}
	if value, ok := uc.mutation.UsedFor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsedFor,
		})
		_node.UsedFor = value
	}
	if value, ok := uc.mutation.Labels(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldLabels,
		})
		_node.Labels = value
	}
	if value, ok := uc.mutation.Balance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: user.FieldBalance,
		})
		_node.Balance = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsert) SetCreatedAt(v uint32) *UserUpsert {
	u.Set(user.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateCreatedAt() *UserUpsert {
	u.SetExcluded(user.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserUpsert) AddCreatedAt(v uint32) *UserUpsert {
	u.Add(user.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsert) SetUpdatedAt(v uint32) *UserUpsert {
	u.Set(user.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateUpdatedAt() *UserUpsert {
	u.SetExcluded(user.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserUpsert) AddUpdatedAt(v uint32) *UserUpsert {
	u.Add(user.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserUpsert) SetDeletedAt(v uint32) *UserUpsert {
	u.Set(user.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateDeletedAt() *UserUpsert {
	u.SetExcluded(user.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserUpsert) AddDeletedAt(v uint32) *UserUpsert {
	u.Add(user.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserUpsert) SetAppID(v uuid.UUID) *UserUpsert {
	u.Set(user.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserUpsert) UpdateAppID() *UserUpsert {
	u.SetExcluded(user.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserUpsert) ClearAppID() *UserUpsert {
	u.SetNull(user.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserUpsert) SetUserID(v uuid.UUID) *UserUpsert {
	u.Set(user.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserUpsert) UpdateUserID() *UserUpsert {
	u.SetExcluded(user.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserUpsert) ClearUserID() *UserUpsert {
	u.SetNull(user.FieldUserID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserUpsert) SetCoinTypeID(v uuid.UUID) *UserUpsert {
	u.Set(user.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserUpsert) UpdateCoinTypeID() *UserUpsert {
	u.SetExcluded(user.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *UserUpsert) ClearCoinTypeID() *UserUpsert {
	u.SetNull(user.FieldCoinTypeID)
	return u
}

// SetAccountID sets the "account_id" field.
func (u *UserUpsert) SetAccountID(v uuid.UUID) *UserUpsert {
	u.Set(user.FieldAccountID, v)
	return u
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *UserUpsert) UpdateAccountID() *UserUpsert {
	u.SetExcluded(user.FieldAccountID)
	return u
}

// ClearAccountID clears the value of the "account_id" field.
func (u *UserUpsert) ClearAccountID() *UserUpsert {
	u.SetNull(user.FieldAccountID)
	return u
}

// SetUsedFor sets the "used_for" field.
func (u *UserUpsert) SetUsedFor(v string) *UserUpsert {
	u.Set(user.FieldUsedFor, v)
	return u
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *UserUpsert) UpdateUsedFor() *UserUpsert {
	u.SetExcluded(user.FieldUsedFor)
	return u
}

// ClearUsedFor clears the value of the "used_for" field.
func (u *UserUpsert) ClearUsedFor() *UserUpsert {
	u.SetNull(user.FieldUsedFor)
	return u
}

// SetLabels sets the "labels" field.
func (u *UserUpsert) SetLabels(v []string) *UserUpsert {
	u.Set(user.FieldLabels, v)
	return u
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *UserUpsert) UpdateLabels() *UserUpsert {
	u.SetExcluded(user.FieldLabels)
	return u
}

// ClearLabels clears the value of the "labels" field.
func (u *UserUpsert) ClearLabels() *UserUpsert {
	u.SetNull(user.FieldLabels)
	return u
}

// SetBalance sets the "balance" field.
func (u *UserUpsert) SetBalance(v decimal.Decimal) *UserUpsert {
	u.Set(user.FieldBalance, v)
	return u
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *UserUpsert) UpdateBalance() *UserUpsert {
	u.SetExcluded(user.FieldBalance)
	return u
}

// ClearBalance clears the value of the "balance" field.
func (u *UserUpsert) ClearBalance() *UserUpsert {
	u.SetNull(user.FieldBalance)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(user.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.User.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsertOne) SetCreatedAt(v uint32) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserUpsertOne) AddCreatedAt(v uint32) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateCreatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertOne) SetUpdatedAt(v uint32) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserUpsertOne) AddUpdatedAt(v uint32) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUpdatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserUpsertOne) SetDeletedAt(v uint32) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserUpsertOne) AddDeletedAt(v uint32) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateDeletedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *UserUpsertOne) SetAppID(v uuid.UUID) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateAppID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserUpsertOne) ClearAppID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserUpsertOne) SetUserID(v uuid.UUID) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUserID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserUpsertOne) ClearUserID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserUpsertOne) SetCoinTypeID(v uuid.UUID) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateCoinTypeID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *UserUpsertOne) ClearCoinTypeID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetAccountID sets the "account_id" field.
func (u *UserUpsertOne) SetAccountID(v uuid.UUID) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateAccountID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAccountID()
	})
}

// ClearAccountID clears the value of the "account_id" field.
func (u *UserUpsertOne) ClearAccountID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearAccountID()
	})
}

// SetUsedFor sets the "used_for" field.
func (u *UserUpsertOne) SetUsedFor(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUsedFor(v)
	})
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUsedFor() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUsedFor()
	})
}

// ClearUsedFor clears the value of the "used_for" field.
func (u *UserUpsertOne) ClearUsedFor() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearUsedFor()
	})
}

// SetLabels sets the "labels" field.
func (u *UserUpsertOne) SetLabels(v []string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateLabels() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLabels()
	})
}

// ClearLabels clears the value of the "labels" field.
func (u *UserUpsertOne) ClearLabels() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearLabels()
	})
}

// SetBalance sets the "balance" field.
func (u *UserUpsertOne) SetBalance(v decimal.Decimal) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetBalance(v)
	})
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateBalance() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBalance()
	})
}

// ClearBalance clears the value of the "balance" field.
func (u *UserUpsertOne) ClearBalance() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearBalance()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserUpsertOne.ID is not supported by MySQL driver. Use UserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(user.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsertBulk) SetCreatedAt(v uint32) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserUpsertBulk) AddCreatedAt(v uint32) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateCreatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertBulk) SetUpdatedAt(v uint32) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserUpsertBulk) AddUpdatedAt(v uint32) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUpdatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserUpsertBulk) SetDeletedAt(v uint32) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserUpsertBulk) AddDeletedAt(v uint32) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateDeletedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *UserUpsertBulk) SetAppID(v uuid.UUID) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateAppID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserUpsertBulk) ClearAppID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserUpsertBulk) SetUserID(v uuid.UUID) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUserID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserUpsertBulk) ClearUserID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserUpsertBulk) SetCoinTypeID(v uuid.UUID) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateCoinTypeID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *UserUpsertBulk) ClearCoinTypeID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetAccountID sets the "account_id" field.
func (u *UserUpsertBulk) SetAccountID(v uuid.UUID) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateAccountID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAccountID()
	})
}

// ClearAccountID clears the value of the "account_id" field.
func (u *UserUpsertBulk) ClearAccountID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearAccountID()
	})
}

// SetUsedFor sets the "used_for" field.
func (u *UserUpsertBulk) SetUsedFor(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUsedFor(v)
	})
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUsedFor() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUsedFor()
	})
}

// ClearUsedFor clears the value of the "used_for" field.
func (u *UserUpsertBulk) ClearUsedFor() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearUsedFor()
	})
}

// SetLabels sets the "labels" field.
func (u *UserUpsertBulk) SetLabels(v []string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateLabels() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLabels()
	})
}

// ClearLabels clears the value of the "labels" field.
func (u *UserUpsertBulk) ClearLabels() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearLabels()
	})
}

// SetBalance sets the "balance" field.
func (u *UserUpsertBulk) SetBalance(v decimal.Decimal) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetBalance(v)
	})
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateBalance() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBalance()
	})
}

// ClearBalance clears the value of the "balance" field.
func (u *UserUpsertBulk) ClearBalance() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearBalance()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
