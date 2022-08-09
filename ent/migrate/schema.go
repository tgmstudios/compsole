// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompetitionsColumns holds the columns for the "competitions" table.
	CompetitionsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "competition_competition_to_provider", Type: field.TypeUUID},
	}
	// CompetitionsTable holds the schema information for the "competitions" table.
	CompetitionsTable = &schema.Table{
		Name:       "competitions",
		Columns:    CompetitionsColumns,
		PrimaryKey: []*schema.Column{CompetitionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "competitions_providers_CompetitionToProvider",
				Columns:    []*schema.Column{CompetitionsColumns[2]},
				RefColumns: []*schema.Column{ProvidersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProvidersColumns holds the columns for the "providers" table.
	ProvidersColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeString},
		{Name: "config", Type: field.TypeString},
	}
	// ProvidersTable holds the schema information for the "providers" table.
	ProvidersTable = &schema.Table{
		Name:       "providers",
		Columns:    ProvidersColumns,
		PrimaryKey: []*schema.Column{ProvidersColumns[0]},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "team_number", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "team_team_to_competition", Type: field.TypeUUID},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "teams_competitions_TeamToCompetition",
				Columns:    []*schema.Column{TeamsColumns[3]},
				RefColumns: []*schema.Column{CompetitionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "token", Type: field.TypeString},
		{Name: "expire_at", Type: field.TypeInt64},
		{Name: "user_user_to_token", Type: field.TypeUUID},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tokens_users_UserToToken",
				Columns:    []*schema.Column{TokensColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString, Default: ""},
		{Name: "last_name", Type: field.TypeString, Default: ""},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"USER", "ADMIN"}},
		{Name: "provider", Type: field.TypeEnum, Enums: []string{"LOCAL", "GITLAB"}},
		{Name: "user_user_to_team", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_teams_UserToTeam",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VMObjectsColumns holds the columns for the "vm_objects" table.
	VMObjectsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "identifier", Type: field.TypeString},
		{Name: "ip_addresses", Type: field.TypeJSON, Nullable: true},
		{Name: "vm_object_vm_object_to_team", Type: field.TypeUUID, Nullable: true},
	}
	// VMObjectsTable holds the schema information for the "vm_objects" table.
	VMObjectsTable = &schema.Table{
		Name:       "vm_objects",
		Columns:    VMObjectsColumns,
		PrimaryKey: []*schema.Column{VMObjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vm_objects_teams_VmObjectToTeam",
				Columns:    []*schema.Column{VMObjectsColumns[4]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompetitionsTable,
		ProvidersTable,
		TeamsTable,
		TokensTable,
		UsersTable,
		VMObjectsTable,
	}
)

func init() {
	CompetitionsTable.ForeignKeys[0].RefTable = ProvidersTable
	TeamsTable.ForeignKeys[0].RefTable = CompetitionsTable
	TokensTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = TeamsTable
	VMObjectsTable.ForeignKeys[0].RefTable = TeamsTable
}
