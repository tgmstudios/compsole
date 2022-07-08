// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/BradHacker/compsole/ent/competition"
	"github.com/BradHacker/compsole/ent/schema"
	"github.com/BradHacker/compsole/ent/team"
	"github.com/BradHacker/compsole/ent/token"
	"github.com/BradHacker/compsole/ent/user"
	"github.com/BradHacker/compsole/ent/vmobject"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	competitionFields := schema.Competition{}.Fields()
	_ = competitionFields
	// competitionDescID is the schema descriptor for id field.
	competitionDescID := competitionFields[0].Descriptor()
	// competition.DefaultID holds the default value on creation for the id field.
	competition.DefaultID = competitionDescID.Default.(func() uuid.UUID)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescID is the schema descriptor for id field.
	teamDescID := teamFields[0].Descriptor()
	// team.DefaultID holds the default value on creation for the id field.
	team.DefaultID = teamDescID.Default.(func() uuid.UUID)
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescID is the schema descriptor for id field.
	tokenDescID := tokenFields[0].Descriptor()
	// token.DefaultID holds the default value on creation for the id field.
	token.DefaultID = tokenDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[3].Descriptor()
	// user.DefaultFirstName holds the default value on creation for the first_name field.
	user.DefaultFirstName = userDescFirstName.Default.(string)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[4].Descriptor()
	// user.DefaultLastName holds the default value on creation for the last_name field.
	user.DefaultLastName = userDescLastName.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	vmobjectFields := schema.VmObject{}.Fields()
	_ = vmobjectFields
	// vmobjectDescID is the schema descriptor for id field.
	vmobjectDescID := vmobjectFields[0].Descriptor()
	// vmobject.DefaultID holds the default value on creation for the id field.
	vmobject.DefaultID = vmobjectDescID.Default.(func() uuid.UUID)
}
