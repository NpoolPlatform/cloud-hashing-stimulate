// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userinvitationcode"
)

// UserInvitationCode is the model entity for the UserInvitationCode schema.
type UserInvitationCode struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserInvitationCode) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userinvitationcode.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserInvitationCode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserInvitationCode fields.
func (uic *UserInvitationCode) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userinvitationcode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uic.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this UserInvitationCode.
// Note that you need to call UserInvitationCode.Unwrap() before calling this method if this UserInvitationCode
// was returned from a transaction, and the transaction was committed or rolled back.
func (uic *UserInvitationCode) Update() *UserInvitationCodeUpdateOne {
	return (&UserInvitationCodeClient{config: uic.config}).UpdateOne(uic)
}

// Unwrap unwraps the UserInvitationCode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uic *UserInvitationCode) Unwrap() *UserInvitationCode {
	tx, ok := uic.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserInvitationCode is not a transactional entity")
	}
	uic.config.driver = tx.drv
	return uic
}

// String implements the fmt.Stringer.
func (uic *UserInvitationCode) String() string {
	var builder strings.Builder
	builder.WriteString("UserInvitationCode(")
	builder.WriteString(fmt.Sprintf("id=%v", uic.ID))
	builder.WriteByte(')')
	return builder.String()
}

// UserInvitationCodes is a parsable slice of UserInvitationCode.
type UserInvitationCodes []*UserInvitationCode

func (uic UserInvitationCodes) config(cfg config) {
	for _i := range uic {
		uic[_i].config = cfg
	}
}
