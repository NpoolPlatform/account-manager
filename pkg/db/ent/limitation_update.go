// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/limitation"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// LimitationUpdate is the builder for updating Limitation entities.
type LimitationUpdate struct {
	config
	hooks     []Hook
	mutation  *LimitationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LimitationUpdate builder.
func (lu *LimitationUpdate) Where(ps ...predicate.Limitation) *LimitationUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LimitationUpdate) SetCreatedAt(u uint32) *LimitationUpdate {
	lu.mutation.ResetCreatedAt()
	lu.mutation.SetCreatedAt(u)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LimitationUpdate) SetNillableCreatedAt(u *uint32) *LimitationUpdate {
	if u != nil {
		lu.SetCreatedAt(*u)
	}
	return lu
}

// AddCreatedAt adds u to the "created_at" field.
func (lu *LimitationUpdate) AddCreatedAt(u int32) *LimitationUpdate {
	lu.mutation.AddCreatedAt(u)
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LimitationUpdate) SetUpdatedAt(u uint32) *LimitationUpdate {
	lu.mutation.ResetUpdatedAt()
	lu.mutation.SetUpdatedAt(u)
	return lu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (lu *LimitationUpdate) AddUpdatedAt(u int32) *LimitationUpdate {
	lu.mutation.AddUpdatedAt(u)
	return lu
}

// SetDeletedAt sets the "deleted_at" field.
func (lu *LimitationUpdate) SetDeletedAt(u uint32) *LimitationUpdate {
	lu.mutation.ResetDeletedAt()
	lu.mutation.SetDeletedAt(u)
	return lu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lu *LimitationUpdate) SetNillableDeletedAt(u *uint32) *LimitationUpdate {
	if u != nil {
		lu.SetDeletedAt(*u)
	}
	return lu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (lu *LimitationUpdate) AddDeletedAt(u int32) *LimitationUpdate {
	lu.mutation.AddDeletedAt(u)
	return lu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (lu *LimitationUpdate) SetCoinTypeID(u uuid.UUID) *LimitationUpdate {
	lu.mutation.SetCoinTypeID(u)
	return lu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (lu *LimitationUpdate) SetNillableCoinTypeID(u *uuid.UUID) *LimitationUpdate {
	if u != nil {
		lu.SetCoinTypeID(*u)
	}
	return lu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (lu *LimitationUpdate) ClearCoinTypeID() *LimitationUpdate {
	lu.mutation.ClearCoinTypeID()
	return lu
}

// SetLimitation sets the "limitation" field.
func (lu *LimitationUpdate) SetLimitation(s string) *LimitationUpdate {
	lu.mutation.SetLimitation(s)
	return lu
}

// SetNillableLimitation sets the "limitation" field if the given value is not nil.
func (lu *LimitationUpdate) SetNillableLimitation(s *string) *LimitationUpdate {
	if s != nil {
		lu.SetLimitation(*s)
	}
	return lu
}

// ClearLimitation clears the value of the "limitation" field.
func (lu *LimitationUpdate) ClearLimitation() *LimitationUpdate {
	lu.mutation.ClearLimitation()
	return lu
}

// SetAmount sets the "amount" field.
func (lu *LimitationUpdate) SetAmount(d decimal.Decimal) *LimitationUpdate {
	lu.mutation.SetAmount(d)
	return lu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (lu *LimitationUpdate) SetNillableAmount(d *decimal.Decimal) *LimitationUpdate {
	if d != nil {
		lu.SetAmount(*d)
	}
	return lu
}

// ClearAmount clears the value of the "amount" field.
func (lu *LimitationUpdate) ClearAmount() *LimitationUpdate {
	lu.mutation.ClearAmount()
	return lu
}

// Mutation returns the LimitationMutation object of the builder.
func (lu *LimitationUpdate) Mutation() *LimitationMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LimitationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := lu.defaults(); err != nil {
		return 0, err
	}
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LimitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LimitationUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LimitationUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LimitationUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LimitationUpdate) defaults() error {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		if limitation.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized limitation.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := limitation.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lu *LimitationUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LimitationUpdate {
	lu.modifiers = append(lu.modifiers, modifiers...)
	return lu
}

func (lu *LimitationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   limitation.Table,
			Columns: limitation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: limitation.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldCreatedAt,
		})
	}
	if value, ok := lu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldCreatedAt,
		})
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldUpdatedAt,
		})
	}
	if value, ok := lu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldUpdatedAt,
		})
	}
	if value, ok := lu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldDeletedAt,
		})
	}
	if value, ok := lu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldDeletedAt,
		})
	}
	if value, ok := lu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: limitation.FieldCoinTypeID,
		})
	}
	if lu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: limitation.FieldCoinTypeID,
		})
	}
	if value, ok := lu.mutation.Limitation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: limitation.FieldLimitation,
		})
	}
	if lu.mutation.LimitationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: limitation.FieldLimitation,
		})
	}
	if value, ok := lu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: limitation.FieldAmount,
		})
	}
	if lu.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: limitation.FieldAmount,
		})
	}
	_spec.Modifiers = lu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{limitation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LimitationUpdateOne is the builder for updating a single Limitation entity.
type LimitationUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LimitationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (luo *LimitationUpdateOne) SetCreatedAt(u uint32) *LimitationUpdateOne {
	luo.mutation.ResetCreatedAt()
	luo.mutation.SetCreatedAt(u)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LimitationUpdateOne) SetNillableCreatedAt(u *uint32) *LimitationUpdateOne {
	if u != nil {
		luo.SetCreatedAt(*u)
	}
	return luo
}

// AddCreatedAt adds u to the "created_at" field.
func (luo *LimitationUpdateOne) AddCreatedAt(u int32) *LimitationUpdateOne {
	luo.mutation.AddCreatedAt(u)
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LimitationUpdateOne) SetUpdatedAt(u uint32) *LimitationUpdateOne {
	luo.mutation.ResetUpdatedAt()
	luo.mutation.SetUpdatedAt(u)
	return luo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (luo *LimitationUpdateOne) AddUpdatedAt(u int32) *LimitationUpdateOne {
	luo.mutation.AddUpdatedAt(u)
	return luo
}

// SetDeletedAt sets the "deleted_at" field.
func (luo *LimitationUpdateOne) SetDeletedAt(u uint32) *LimitationUpdateOne {
	luo.mutation.ResetDeletedAt()
	luo.mutation.SetDeletedAt(u)
	return luo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (luo *LimitationUpdateOne) SetNillableDeletedAt(u *uint32) *LimitationUpdateOne {
	if u != nil {
		luo.SetDeletedAt(*u)
	}
	return luo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (luo *LimitationUpdateOne) AddDeletedAt(u int32) *LimitationUpdateOne {
	luo.mutation.AddDeletedAt(u)
	return luo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (luo *LimitationUpdateOne) SetCoinTypeID(u uuid.UUID) *LimitationUpdateOne {
	luo.mutation.SetCoinTypeID(u)
	return luo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (luo *LimitationUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *LimitationUpdateOne {
	if u != nil {
		luo.SetCoinTypeID(*u)
	}
	return luo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (luo *LimitationUpdateOne) ClearCoinTypeID() *LimitationUpdateOne {
	luo.mutation.ClearCoinTypeID()
	return luo
}

// SetLimitation sets the "limitation" field.
func (luo *LimitationUpdateOne) SetLimitation(s string) *LimitationUpdateOne {
	luo.mutation.SetLimitation(s)
	return luo
}

// SetNillableLimitation sets the "limitation" field if the given value is not nil.
func (luo *LimitationUpdateOne) SetNillableLimitation(s *string) *LimitationUpdateOne {
	if s != nil {
		luo.SetLimitation(*s)
	}
	return luo
}

// ClearLimitation clears the value of the "limitation" field.
func (luo *LimitationUpdateOne) ClearLimitation() *LimitationUpdateOne {
	luo.mutation.ClearLimitation()
	return luo
}

// SetAmount sets the "amount" field.
func (luo *LimitationUpdateOne) SetAmount(d decimal.Decimal) *LimitationUpdateOne {
	luo.mutation.SetAmount(d)
	return luo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (luo *LimitationUpdateOne) SetNillableAmount(d *decimal.Decimal) *LimitationUpdateOne {
	if d != nil {
		luo.SetAmount(*d)
	}
	return luo
}

// ClearAmount clears the value of the "amount" field.
func (luo *LimitationUpdateOne) ClearAmount() *LimitationUpdateOne {
	luo.mutation.ClearAmount()
	return luo
}

// Mutation returns the LimitationMutation object of the builder.
func (luo *LimitationUpdateOne) Mutation() *LimitationMutation {
	return luo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LimitationUpdateOne) Select(field string, fields ...string) *LimitationUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Limitation entity.
func (luo *LimitationUpdateOne) Save(ctx context.Context) (*Limitation, error) {
	var (
		err  error
		node *Limitation
	)
	if err := luo.defaults(); err != nil {
		return nil, err
	}
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LimitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, luo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Limitation)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LimitationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LimitationUpdateOne) SaveX(ctx context.Context) *Limitation {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LimitationUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LimitationUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LimitationUpdateOne) defaults() error {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		if limitation.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized limitation.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := limitation.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (luo *LimitationUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LimitationUpdateOne {
	luo.modifiers = append(luo.modifiers, modifiers...)
	return luo
}

func (luo *LimitationUpdateOne) sqlSave(ctx context.Context) (_node *Limitation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   limitation.Table,
			Columns: limitation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: limitation.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Limitation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, limitation.FieldID)
		for _, f := range fields {
			if !limitation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != limitation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldCreatedAt,
		})
	}
	if value, ok := luo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldCreatedAt,
		})
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldUpdatedAt,
		})
	}
	if value, ok := luo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldUpdatedAt,
		})
	}
	if value, ok := luo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldDeletedAt,
		})
	}
	if value, ok := luo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: limitation.FieldDeletedAt,
		})
	}
	if value, ok := luo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: limitation.FieldCoinTypeID,
		})
	}
	if luo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: limitation.FieldCoinTypeID,
		})
	}
	if value, ok := luo.mutation.Limitation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: limitation.FieldLimitation,
		})
	}
	if luo.mutation.LimitationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: limitation.FieldLimitation,
		})
	}
	if value, ok := luo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: limitation.FieldAmount,
		})
	}
	if luo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: limitation.FieldAmount,
		})
	}
	_spec.Modifiers = luo.modifiers
	_node = &Limitation{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{limitation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
