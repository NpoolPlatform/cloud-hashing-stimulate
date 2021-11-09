// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/registrationinvitation"
	"github.com/google/uuid"
)

// RegistrationInvitation is the model entity for the RegistrationInvitation schema.
type RegistrationInvitation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
	// InviterID holds the value of the "inviter_id" field.
	InviterID uuid.UUID `json:"inviter_id,omitempty"`
	// InviteeID holds the value of the "invitee_id" field.
	InviteeID uuid.UUID `json:"invitee_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// Fulfilled holds the value of the "fulfilled" field.
	Fulfilled bool `json:"fulfilled,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RegistrationInvitation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case registrationinvitation.FieldFulfilled:
			values[i] = new(sql.NullBool)
		case registrationinvitation.FieldCreateAt, registrationinvitation.FieldUpdateAt, registrationinvitation.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case registrationinvitation.FieldID, registrationinvitation.FieldInviterID, registrationinvitation.FieldInviteeID, registrationinvitation.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RegistrationInvitation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RegistrationInvitation fields.
func (ri *RegistrationInvitation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case registrationinvitation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ri.ID = *value
			}
		case registrationinvitation.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ri.CreateAt = uint32(value.Int64)
			}
		case registrationinvitation.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ri.UpdateAt = uint32(value.Int64)
			}
		case registrationinvitation.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ri.DeleteAt = uint32(value.Int64)
			}
		case registrationinvitation.FieldInviterID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field inviter_id", values[i])
			} else if value != nil {
				ri.InviterID = *value
			}
		case registrationinvitation.FieldInviteeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field invitee_id", values[i])
			} else if value != nil {
				ri.InviteeID = *value
			}
		case registrationinvitation.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ri.AppID = *value
			}
		case registrationinvitation.FieldFulfilled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field fulfilled", values[i])
			} else if value.Valid {
				ri.Fulfilled = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this RegistrationInvitation.
// Note that you need to call RegistrationInvitation.Unwrap() before calling this method if this RegistrationInvitation
// was returned from a transaction, and the transaction was committed or rolled back.
func (ri *RegistrationInvitation) Update() *RegistrationInvitationUpdateOne {
	return (&RegistrationInvitationClient{config: ri.config}).UpdateOne(ri)
}

// Unwrap unwraps the RegistrationInvitation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ri *RegistrationInvitation) Unwrap() *RegistrationInvitation {
	tx, ok := ri.config.driver.(*txDriver)
	if !ok {
		panic("ent: RegistrationInvitation is not a transactional entity")
	}
	ri.config.driver = tx.drv
	return ri
}

// String implements the fmt.Stringer.
func (ri *RegistrationInvitation) String() string {
	var builder strings.Builder
	builder.WriteString("RegistrationInvitation(")
	builder.WriteString(fmt.Sprintf("id=%v", ri.ID))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ri.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ri.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", ri.DeleteAt))
	builder.WriteString(", inviter_id=")
	builder.WriteString(fmt.Sprintf("%v", ri.InviterID))
	builder.WriteString(", invitee_id=")
	builder.WriteString(fmt.Sprintf("%v", ri.InviteeID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", ri.AppID))
	builder.WriteString(", fulfilled=")
	builder.WriteString(fmt.Sprintf("%v", ri.Fulfilled))
	builder.WriteByte(')')
	return builder.String()
}

// RegistrationInvitations is a parsable slice of RegistrationInvitation.
type RegistrationInvitations []*RegistrationInvitation

func (ri RegistrationInvitations) config(cfg config) {
	for _i := range ri {
		ri[_i].config = cfg
	}
}
