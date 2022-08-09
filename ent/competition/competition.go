// Code generated by entc, DO NOT EDIT.

package competition

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the competition type in the database.
	Label = "competition"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCompetitionToTeams holds the string denoting the competitiontoteams edge name in mutations.
	EdgeCompetitionToTeams = "CompetitionToTeams"
	// EdgeCompetitionToProvider holds the string denoting the competitiontoprovider edge name in mutations.
	EdgeCompetitionToProvider = "CompetitionToProvider"
	// Table holds the table name of the competition in the database.
	Table = "competitions"
	// CompetitionToTeamsTable is the table that holds the CompetitionToTeams relation/edge.
	CompetitionToTeamsTable = "teams"
	// CompetitionToTeamsInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	CompetitionToTeamsInverseTable = "teams"
	// CompetitionToTeamsColumn is the table column denoting the CompetitionToTeams relation/edge.
	CompetitionToTeamsColumn = "team_team_to_competition"
	// CompetitionToProviderTable is the table that holds the CompetitionToProvider relation/edge.
	CompetitionToProviderTable = "competitions"
	// CompetitionToProviderInverseTable is the table name for the Provider entity.
	// It exists in this package in order to avoid circular dependency with the "provider" package.
	CompetitionToProviderInverseTable = "providers"
	// CompetitionToProviderColumn is the table column denoting the CompetitionToProvider relation/edge.
	CompetitionToProviderColumn = "competition_competition_to_provider"
)

// Columns holds all SQL columns for competition fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "competitions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"competition_competition_to_provider",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
