// Code generated by ent, DO NOT EDIT.

package registrationinvitation

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the registrationinvitation type in the database.
	Label = "registration_invitation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// FieldInviterID holds the string denoting the inviter_id field in the database.
	FieldInviterID = "inviter_id"
	// FieldInviteeID holds the string denoting the invitee_id field in the database.
	FieldInviteeID = "invitee_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// Table holds the table name of the registrationinvitation in the database.
	Table = "registration_invitations"
)

// Columns holds all SQL columns for registrationinvitation fields.
var Columns = []string{
	FieldID,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
	FieldInviterID,
	FieldInviteeID,
	FieldAppID,
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
//	import _ "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
