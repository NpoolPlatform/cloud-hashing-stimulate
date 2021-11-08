// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AgencySettingsColumns holds the columns for the "agency_settings" table.
	AgencySettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// AgencySettingsTable holds the schema information for the "agency_settings" table.
	AgencySettingsTable = &schema.Table{
		Name:       "agency_settings",
		Columns:    AgencySettingsColumns,
		PrimaryKey: []*schema.Column{AgencySettingsColumns[0]},
	}
	// NewUserRewardSettingsColumns holds the columns for the "new_user_reward_settings" table.
	NewUserRewardSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// NewUserRewardSettingsTable holds the schema information for the "new_user_reward_settings" table.
	NewUserRewardSettingsTable = &schema.Table{
		Name:       "new_user_reward_settings",
		Columns:    NewUserRewardSettingsColumns,
		PrimaryKey: []*schema.Column{NewUserRewardSettingsColumns[0]},
	}
	// PurchaseInvitationsColumns holds the columns for the "purchase_invitations" table.
	PurchaseInvitationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// PurchaseInvitationsTable holds the schema information for the "purchase_invitations" table.
	PurchaseInvitationsTable = &schema.Table{
		Name:       "purchase_invitations",
		Columns:    PurchaseInvitationsColumns,
		PrimaryKey: []*schema.Column{PurchaseInvitationsColumns[0]},
	}
	// RegistrationInvitationsColumns holds the columns for the "registration_invitations" table.
	RegistrationInvitationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// RegistrationInvitationsTable holds the schema information for the "registration_invitations" table.
	RegistrationInvitationsTable = &schema.Table{
		Name:       "registration_invitations",
		Columns:    RegistrationInvitationsColumns,
		PrimaryKey: []*schema.Column{RegistrationInvitationsColumns[0]},
	}
	// UserInvitationCodesColumns holds the columns for the "user_invitation_codes" table.
	UserInvitationCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UserInvitationCodesTable holds the schema information for the "user_invitation_codes" table.
	UserInvitationCodesTable = &schema.Table{
		Name:       "user_invitation_codes",
		Columns:    UserInvitationCodesColumns,
		PrimaryKey: []*schema.Column{UserInvitationCodesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AgencySettingsTable,
		NewUserRewardSettingsTable,
		PurchaseInvitationsTable,
		RegistrationInvitationsTable,
		UserInvitationCodesTable,
	}
)

func init() {
}
