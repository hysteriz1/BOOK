// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/m16_z/app/ent/roomstatus"
)

// RoomStatus is the model entity for the RoomStatus schema.
type RoomStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// STATUSDATA holds the value of the "STATUSDATA" field.
	STATUSDATA string `json:"STATUSDATA,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomStatusQuery when eager-loading is set.
	Edges RoomStatusEdges `json:"edges"`
}

// RoomStatusEdges holds the relations/edges for other nodes in the graph.
type RoomStatusEdges struct {
	// STATUSROOM holds the value of the STATUS_ROOM edge.
	STATUSROOM []*Room
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// STATUSROOMOrErr returns the STATUSROOM value or an error if the edge
// was not loaded in eager-loading.
func (e RoomStatusEdges) STATUSROOMOrErr() ([]*Room, error) {
	if e.loadedTypes[0] {
		return e.STATUSROOM, nil
	}
	return nil, &NotLoadedError{edge: "STATUS_ROOM"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoomStatus) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // STATUSDATA
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoomStatus fields.
func (rs *RoomStatus) assignValues(values ...interface{}) error {
	if m, n := len(values), len(roomstatus.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	rs.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field STATUSDATA", values[0])
	} else if value.Valid {
		rs.STATUSDATA = value.String
	}
	return nil
}

// QuerySTATUSROOM queries the STATUS_ROOM edge of the RoomStatus.
func (rs *RoomStatus) QuerySTATUSROOM() *RoomQuery {
	return (&RoomStatusClient{config: rs.config}).QuerySTATUSROOM(rs)
}

// Update returns a builder for updating this RoomStatus.
// Note that, you need to call RoomStatus.Unwrap() before calling this method, if this RoomStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (rs *RoomStatus) Update() *RoomStatusUpdateOne {
	return (&RoomStatusClient{config: rs.config}).UpdateOne(rs)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (rs *RoomStatus) Unwrap() *RoomStatus {
	tx, ok := rs.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoomStatus is not a transactional entity")
	}
	rs.config.driver = tx.drv
	return rs
}

// String implements the fmt.Stringer.
func (rs *RoomStatus) String() string {
	var builder strings.Builder
	builder.WriteString("RoomStatus(")
	builder.WriteString(fmt.Sprintf("id=%v", rs.ID))
	builder.WriteString(", STATUSDATA=")
	builder.WriteString(rs.STATUSDATA)
	builder.WriteByte(')')
	return builder.String()
}

// RoomStatusSlice is a parsable slice of RoomStatus.
type RoomStatusSlice []*RoomStatus

func (rs RoomStatusSlice) config(cfg config) {
	for _i := range rs {
		rs[_i].config = cfg
	}
}
