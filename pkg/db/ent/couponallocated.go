// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponallocated"
	"github.com/google/uuid"
)

// CouponAllocated is the model entity for the CouponAllocated schema.
type CouponAllocated struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// Used holds the value of the "used" field.
	Used bool `json:"used,omitempty"`
	// CouponID holds the value of the "coupon_id" field.
	CouponID uuid.UUID `json:"coupon_id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CouponAllocated) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case couponallocated.FieldUsed:
			values[i] = new(sql.NullBool)
		case couponallocated.FieldCreateAt, couponallocated.FieldUpdateAt, couponallocated.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case couponallocated.FieldID, couponallocated.FieldUserID, couponallocated.FieldAppID, couponallocated.FieldCouponID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CouponAllocated", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CouponAllocated fields.
func (ca *CouponAllocated) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case couponallocated.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ca.ID = *value
			}
		case couponallocated.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ca.UserID = *value
			}
		case couponallocated.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ca.AppID = *value
			}
		case couponallocated.FieldUsed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field used", values[i])
			} else if value.Valid {
				ca.Used = value.Bool
			}
		case couponallocated.FieldCouponID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coupon_id", values[i])
			} else if value != nil {
				ca.CouponID = *value
			}
		case couponallocated.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ca.CreateAt = uint32(value.Int64)
			}
		case couponallocated.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ca.UpdateAt = uint32(value.Int64)
			}
		case couponallocated.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ca.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CouponAllocated.
// Note that you need to call CouponAllocated.Unwrap() before calling this method if this CouponAllocated
// was returned from a transaction, and the transaction was committed or rolled back.
func (ca *CouponAllocated) Update() *CouponAllocatedUpdateOne {
	return (&CouponAllocatedClient{config: ca.config}).UpdateOne(ca)
}

// Unwrap unwraps the CouponAllocated entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ca *CouponAllocated) Unwrap() *CouponAllocated {
	tx, ok := ca.config.driver.(*txDriver)
	if !ok {
		panic("ent: CouponAllocated is not a transactional entity")
	}
	ca.config.driver = tx.drv
	return ca
}

// String implements the fmt.Stringer.
func (ca *CouponAllocated) String() string {
	var builder strings.Builder
	builder.WriteString("CouponAllocated(")
	builder.WriteString(fmt.Sprintf("id=%v", ca.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.UserID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.AppID))
	builder.WriteString(", used=")
	builder.WriteString(fmt.Sprintf("%v", ca.Used))
	builder.WriteString(", coupon_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.CouponID))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ca.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ca.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", ca.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// CouponAllocateds is a parsable slice of CouponAllocated.
type CouponAllocateds []*CouponAllocated

func (ca CouponAllocateds) config(cfg config) {
	for _i := range ca {
		ca[_i].config = cfg
	}
}
