// Code generated by ent, DO NOT EDIT.

package platform

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the platform type in the database.
	Label = "platform"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldUsedFor holds the string denoting the used_for field in the database.
	FieldUsedFor = "used_for"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldBackup holds the string denoting the backup field in the database.
	FieldBackup = "backup"
	// Table holds the table name of the platform in the database.
	Table = "platforms"
)

// Columns holds all SQL columns for platform fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAccountID,
	FieldUsedFor,
	FieldGoodID,
	FieldBackup,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/account-manager/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultAccountID holds the default value on creation for the "account_id" field.
	DefaultAccountID func() uuid.UUID
	// DefaultUsedFor holds the default value on creation for the "used_for" field.
	DefaultUsedFor string
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultBackup holds the default value on creation for the "backup" field.
	DefaultBackup bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
