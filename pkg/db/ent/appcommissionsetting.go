// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcommissionsetting"
	"github.com/google/uuid"
)

// AppCommissionSetting is the model entity for the AppCommissionSetting schema.
type AppCommissionSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Level holds the value of the "level" field.
	Level uint32 `json:"level,omitempty"`
	// InvitationDiscount holds the value of the "invitation_discount" field.
	InvitationDiscount bool `json:"invitation_discount,omitempty"`
	// UniqueSetting holds the value of the "unique_setting" field.
	UniqueSetting bool `json:"unique_setting,omitempty"`
	// KpiSetting holds the value of the "kpi_setting" field.
	KpiSetting bool `json:"kpi_setting,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppCommissionSetting) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appcommissionsetting.FieldInvitationDiscount, appcommissionsetting.FieldUniqueSetting, appcommissionsetting.FieldKpiSetting:
			values[i] = new(sql.NullBool)
		case appcommissionsetting.FieldLevel, appcommissionsetting.FieldCreateAt, appcommissionsetting.FieldUpdateAt, appcommissionsetting.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case appcommissionsetting.FieldType:
			values[i] = new(sql.NullString)
		case appcommissionsetting.FieldID, appcommissionsetting.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppCommissionSetting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppCommissionSetting fields.
func (acs *AppCommissionSetting) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appcommissionsetting.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				acs.ID = *value
			}
		case appcommissionsetting.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				acs.AppID = *value
			}
		case appcommissionsetting.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				acs.Type = value.String
			}
		case appcommissionsetting.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				acs.Level = uint32(value.Int64)
			}
		case appcommissionsetting.FieldInvitationDiscount:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field invitation_discount", values[i])
			} else if value.Valid {
				acs.InvitationDiscount = value.Bool
			}
		case appcommissionsetting.FieldUniqueSetting:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field unique_setting", values[i])
			} else if value.Valid {
				acs.UniqueSetting = value.Bool
			}
		case appcommissionsetting.FieldKpiSetting:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field kpi_setting", values[i])
			} else if value.Valid {
				acs.KpiSetting = value.Bool
			}
		case appcommissionsetting.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				acs.CreateAt = uint32(value.Int64)
			}
		case appcommissionsetting.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				acs.UpdateAt = uint32(value.Int64)
			}
		case appcommissionsetting.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				acs.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppCommissionSetting.
// Note that you need to call AppCommissionSetting.Unwrap() before calling this method if this AppCommissionSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (acs *AppCommissionSetting) Update() *AppCommissionSettingUpdateOne {
	return (&AppCommissionSettingClient{config: acs.config}).UpdateOne(acs)
}

// Unwrap unwraps the AppCommissionSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (acs *AppCommissionSetting) Unwrap() *AppCommissionSetting {
	tx, ok := acs.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppCommissionSetting is not a transactional entity")
	}
	acs.config.driver = tx.drv
	return acs
}

// String implements the fmt.Stringer.
func (acs *AppCommissionSetting) String() string {
	var builder strings.Builder
	builder.WriteString("AppCommissionSetting(")
	builder.WriteString(fmt.Sprintf("id=%v", acs.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", acs.AppID))
	builder.WriteString(", type=")
	builder.WriteString(acs.Type)
	builder.WriteString(", level=")
	builder.WriteString(fmt.Sprintf("%v", acs.Level))
	builder.WriteString(", invitation_discount=")
	builder.WriteString(fmt.Sprintf("%v", acs.InvitationDiscount))
	builder.WriteString(", unique_setting=")
	builder.WriteString(fmt.Sprintf("%v", acs.UniqueSetting))
	builder.WriteString(", kpi_setting=")
	builder.WriteString(fmt.Sprintf("%v", acs.KpiSetting))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", acs.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", acs.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", acs.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// AppCommissionSettings is a parsable slice of AppCommissionSetting.
type AppCommissionSettings []*AppCommissionSetting

func (acs AppCommissionSettings) config(cfg config) {
	for _i := range acs {
		acs[_i].config = cfg
	}
}
