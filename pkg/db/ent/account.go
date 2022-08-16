// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	"github.com/google/uuid"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// UsedFor holds the value of the "used_for" field.
	UsedFor string `json:"used_for,omitempty"`
	// PlatformHoldPrivateKey holds the value of the "platform_hold_private_key" field.
	PlatformHoldPrivateKey bool `json:"platform_hold_private_key,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Locked holds the value of the "locked" field.
	Locked bool `json:"locked,omitempty"`
	// Blocked holds the value of the "blocked" field.
	Blocked bool `json:"blocked,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldPlatformHoldPrivateKey, account.FieldActive, account.FieldLocked, account.FieldBlocked:
			values[i] = new(sql.NullBool)
		case account.FieldCreatedAt, account.FieldUpdatedAt, account.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case account.FieldAddress, account.FieldUsedFor:
			values[i] = new(sql.NullString)
		case account.FieldID, account.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Account", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case account.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = uint32(value.Int64)
			}
		case account.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = uint32(value.Int64)
			}
		case account.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = uint32(value.Int64)
			}
		case account.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				a.CoinTypeID = *value
			}
		case account.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				a.Address = value.String
			}
		case account.FieldUsedFor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field used_for", values[i])
			} else if value.Valid {
				a.UsedFor = value.String
			}
		case account.FieldPlatformHoldPrivateKey:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field platform_hold_private_key", values[i])
			} else if value.Valid {
				a.PlatformHoldPrivateKey = value.Bool
			}
		case account.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				a.Active = value.Bool
			}
		case account.FieldLocked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value.Valid {
				a.Locked = value.Bool
			}
		case account.FieldBlocked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field blocked", values[i])
			} else if value.Valid {
				a.Blocked = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedAt))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", a.DeletedAt))
	builder.WriteString(", coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", a.CoinTypeID))
	builder.WriteString(", address=")
	builder.WriteString(a.Address)
	builder.WriteString(", used_for=")
	builder.WriteString(a.UsedFor)
	builder.WriteString(", platform_hold_private_key=")
	builder.WriteString(fmt.Sprintf("%v", a.PlatformHoldPrivateKey))
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", a.Active))
	builder.WriteString(", locked=")
	builder.WriteString(fmt.Sprintf("%v", a.Locked))
	builder.WriteString(", blocked=")
	builder.WriteString(fmt.Sprintf("%v", a.Blocked))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
