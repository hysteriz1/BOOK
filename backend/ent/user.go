// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/m16_z/app/ent/user"
	"github.com/m16_z/app/ent/userstatus"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// USEREMAIL holds the value of the "USEREMAIL" field.
	USEREMAIL string `json:"USEREMAIL,omitempty"`
	// USERNAME holds the value of the "USERNAME" field.
	USERNAME string `json:"USERNAME,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges                       UserEdges `json:"edges"`
	user_status_userstatus_user *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// USERBOOK holds the value of the USER_BOOK edge.
	USERBOOK []*Book
	// USERUSERSTATUS holds the value of the USER_USERSTATUS edge.
	USERUSERSTATUS *UserStatus
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// USERBOOKOrErr returns the USERBOOK value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) USERBOOKOrErr() ([]*Book, error) {
	if e.loadedTypes[0] {
		return e.USERBOOK, nil
	}
	return nil, &NotLoadedError{edge: "USER_BOOK"}
}

// USERUSERSTATUSOrErr returns the USERUSERSTATUS value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) USERUSERSTATUSOrErr() (*UserStatus, error) {
	if e.loadedTypes[1] {
		if e.USERUSERSTATUS == nil {
			// The edge USER_USERSTATUS was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: userstatus.Label}
		}
		return e.USERUSERSTATUS, nil
	}
	return nil, &NotLoadedError{edge: "USER_USERSTATUS"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // USEREMAIL
		&sql.NullString{}, // USERNAME
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_status_userstatus_user
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field USEREMAIL", values[0])
	} else if value.Valid {
		u.USEREMAIL = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field USERNAME", values[1])
	} else if value.Valid {
		u.USERNAME = value.String
	}
	values = values[2:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_status_userstatus_user", value)
		} else if value.Valid {
			u.user_status_userstatus_user = new(int)
			*u.user_status_userstatus_user = int(value.Int64)
		}
	}
	return nil
}

// QueryUSERBOOK queries the USER_BOOK edge of the User.
func (u *User) QueryUSERBOOK() *BookQuery {
	return (&UserClient{config: u.config}).QueryUSERBOOK(u)
}

// QueryUSERUSERSTATUS queries the USER_USERSTATUS edge of the User.
func (u *User) QueryUSERUSERSTATUS() *UserStatusQuery {
	return (&UserClient{config: u.config}).QueryUSERUSERSTATUS(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", USEREMAIL=")
	builder.WriteString(u.USEREMAIL)
	builder.WriteString(", USERNAME=")
	builder.WriteString(u.USERNAME)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
