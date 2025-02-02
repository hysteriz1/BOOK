// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/m16_z/app/ent/roomtype"
)

// RoomType is the model entity for the RoomType schema.
type RoomType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ROOMTYPEDATA holds the value of the "ROOMTYPEDATA" field.
	ROOMTYPEDATA string `json:"ROOMTYPEDATA,omitempty"`
	// COSTDATA holds the value of the "COSTDATA" field.
	COSTDATA int `json:"COSTDATA,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomTypeQuery when eager-loading is set.
	Edges RoomTypeEdges `json:"edges"`
}

// RoomTypeEdges holds the relations/edges for other nodes in the graph.
type RoomTypeEdges struct {
	// ROOMTYPEROOM holds the value of the ROOMTYPE_ROOM edge.
	ROOMTYPEROOM []*Room
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ROOMTYPEROOMOrErr returns the ROOMTYPEROOM value or an error if the edge
// was not loaded in eager-loading.
func (e RoomTypeEdges) ROOMTYPEROOMOrErr() ([]*Room, error) {
	if e.loadedTypes[0] {
		return e.ROOMTYPEROOM, nil
	}
	return nil, &NotLoadedError{edge: "ROOMTYPE_ROOM"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoomType) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // ROOMTYPEDATA
		&sql.NullInt64{},  // COSTDATA
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoomType fields.
func (rt *RoomType) assignValues(values ...interface{}) error {
	if m, n := len(values), len(roomtype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	rt.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field ROOMTYPEDATA", values[0])
	} else if value.Valid {
		rt.ROOMTYPEDATA = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field COSTDATA", values[1])
	} else if value.Valid {
		rt.COSTDATA = int(value.Int64)
	}
	return nil
}

// QueryROOMTYPEROOM queries the ROOMTYPE_ROOM edge of the RoomType.
func (rt *RoomType) QueryROOMTYPEROOM() *RoomQuery {
	return (&RoomTypeClient{config: rt.config}).QueryROOMTYPEROOM(rt)
}

// Update returns a builder for updating this RoomType.
// Note that, you need to call RoomType.Unwrap() before calling this method, if this RoomType
// was returned from a transaction, and the transaction was committed or rolled back.
func (rt *RoomType) Update() *RoomTypeUpdateOne {
	return (&RoomTypeClient{config: rt.config}).UpdateOne(rt)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (rt *RoomType) Unwrap() *RoomType {
	tx, ok := rt.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoomType is not a transactional entity")
	}
	rt.config.driver = tx.drv
	return rt
}

// String implements the fmt.Stringer.
func (rt *RoomType) String() string {
	var builder strings.Builder
	builder.WriteString("RoomType(")
	builder.WriteString(fmt.Sprintf("id=%v", rt.ID))
	builder.WriteString(", ROOMTYPEDATA=")
	builder.WriteString(rt.ROOMTYPEDATA)
	builder.WriteString(", COSTDATA=")
	builder.WriteString(fmt.Sprintf("%v", rt.COSTDATA))
	builder.WriteByte(')')
	return builder.String()
}

// RoomTypes is a parsable slice of RoomType.
type RoomTypes []*RoomType

func (rt RoomTypes) config(cfg config) {
	for _i := range rt {
		rt[_i].config = cfg
	}
}
