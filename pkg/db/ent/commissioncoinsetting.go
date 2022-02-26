// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/commissioncoinsetting"
	"github.com/google/uuid"
)

// CommissionCoinSetting is the model entity for the CommissionCoinSetting schema.
type CommissionCoinSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CommissionCoinSetting) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case commissioncoinsetting.FieldCreateAt, commissioncoinsetting.FieldUpdateAt, commissioncoinsetting.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case commissioncoinsetting.FieldID, commissioncoinsetting.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CommissionCoinSetting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CommissionCoinSetting fields.
func (ccs *CommissionCoinSetting) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case commissioncoinsetting.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ccs.ID = *value
			}
		case commissioncoinsetting.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				ccs.CoinTypeID = *value
			}
		case commissioncoinsetting.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ccs.CreateAt = uint32(value.Int64)
			}
		case commissioncoinsetting.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ccs.UpdateAt = uint32(value.Int64)
			}
		case commissioncoinsetting.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ccs.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CommissionCoinSetting.
// Note that you need to call CommissionCoinSetting.Unwrap() before calling this method if this CommissionCoinSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (ccs *CommissionCoinSetting) Update() *CommissionCoinSettingUpdateOne {
	return (&CommissionCoinSettingClient{config: ccs.config}).UpdateOne(ccs)
}

// Unwrap unwraps the CommissionCoinSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ccs *CommissionCoinSetting) Unwrap() *CommissionCoinSetting {
	tx, ok := ccs.config.driver.(*txDriver)
	if !ok {
		panic("ent: CommissionCoinSetting is not a transactional entity")
	}
	ccs.config.driver = tx.drv
	return ccs
}

// String implements the fmt.Stringer.
func (ccs *CommissionCoinSetting) String() string {
	var builder strings.Builder
	builder.WriteString("CommissionCoinSetting(")
	builder.WriteString(fmt.Sprintf("id=%v", ccs.ID))
	builder.WriteString(", coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", ccs.CoinTypeID))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ccs.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ccs.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", ccs.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// CommissionCoinSettings is a parsable slice of CommissionCoinSetting.
type CommissionCoinSettings []*CommissionCoinSetting

func (ccs CommissionCoinSettings) config(cfg config) {
	for _i := range ccs {
		ccs[_i].config = cfg
	}
}