// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/BradHacker/compsole/ent/token"
	"github.com/BradHacker/compsole/ent/user"
	"github.com/google/uuid"
)

// Token is the model entity for the Token schema.
type Token struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Token holds the value of the "token" field.
	// [REQUIRED] The auth-token cookie value for the user session.
	Token string `json:"token,omitempty"`
	// ExpireAt holds the value of the "expire_at" field.
	// [REQUIRED] The time the token should expire.
	ExpireAt int64 `json:"expire_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TokenQuery when eager-loading is set.
	Edges              TokenEdges `json:"edges"`
	user_user_to_token *uuid.UUID
}

// TokenEdges holds the relations/edges for other nodes in the graph.
type TokenEdges struct {
	// TokenToUser holds the value of the TokenToUser edge.
	TokenToUser *User `json:"TokenToUser,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TokenToUserOrErr returns the TokenToUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TokenEdges) TokenToUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.TokenToUser == nil {
			// The edge TokenToUser was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.TokenToUser, nil
	}
	return nil, &NotLoadedError{edge: "TokenToUser"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Token) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case token.FieldExpireAt:
			values[i] = new(sql.NullInt64)
		case token.FieldToken:
			values[i] = new(sql.NullString)
		case token.FieldID:
			values[i] = new(uuid.UUID)
		case token.ForeignKeys[0]: // user_user_to_token
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Token", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Token fields.
func (t *Token) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case token.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case token.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				t.Token = value.String
			}
		case token.FieldExpireAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field expire_at", values[i])
			} else if value.Valid {
				t.ExpireAt = value.Int64
			}
		case token.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_user_to_token", values[i])
			} else if value.Valid {
				t.user_user_to_token = new(uuid.UUID)
				*t.user_user_to_token = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryTokenToUser queries the "TokenToUser" edge of the Token entity.
func (t *Token) QueryTokenToUser() *UserQuery {
	return (&TokenClient{config: t.config}).QueryTokenToUser(t)
}

// Update returns a builder for updating this Token.
// Note that you need to call Token.Unwrap() before calling this method if this Token
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Token) Update() *TokenUpdateOne {
	return (&TokenClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Token entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Token) Unwrap() *Token {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Token is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Token) String() string {
	var builder strings.Builder
	builder.WriteString("Token(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", token=")
	builder.WriteString(t.Token)
	builder.WriteString(", expire_at=")
	builder.WriteString(fmt.Sprintf("%v", t.ExpireAt))
	builder.WriteByte(')')
	return builder.String()
}

// Tokens is a parsable slice of Token.
type Tokens []*Token

func (t Tokens) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
