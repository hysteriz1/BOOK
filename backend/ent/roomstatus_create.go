// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/m16_z/app/ent/room"
	"github.com/m16_z/app/ent/roomstatus"
)

// RoomStatusCreate is the builder for creating a RoomStatus entity.
type RoomStatusCreate struct {
	config
	mutation *RoomStatusMutation
	hooks    []Hook
}

// SetSTATUSDATA sets the STATUSDATA field.
func (rsc *RoomStatusCreate) SetSTATUSDATA(s string) *RoomStatusCreate {
	rsc.mutation.SetSTATUSDATA(s)
	return rsc
}

// AddSTATUSROOMIDs adds the STATUS_ROOM edge to Room by ids.
func (rsc *RoomStatusCreate) AddSTATUSROOMIDs(ids ...int) *RoomStatusCreate {
	rsc.mutation.AddSTATUSROOMIDs(ids...)
	return rsc
}

// AddSTATUSROOM adds the STATUS_ROOM edges to Room.
func (rsc *RoomStatusCreate) AddSTATUSROOM(r ...*Room) *RoomStatusCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rsc.AddSTATUSROOMIDs(ids...)
}

// Mutation returns the RoomStatusMutation object of the builder.
func (rsc *RoomStatusCreate) Mutation() *RoomStatusMutation {
	return rsc.mutation
}

// Save creates the RoomStatus in the database.
func (rsc *RoomStatusCreate) Save(ctx context.Context) (*RoomStatus, error) {
	if _, ok := rsc.mutation.STATUSDATA(); !ok {
		return nil, &ValidationError{Name: "STATUSDATA", err: errors.New("ent: missing required field \"STATUSDATA\"")}
	}
	if v, ok := rsc.mutation.STATUSDATA(); ok {
		if err := roomstatus.STATUSDATAValidator(v); err != nil {
			return nil, &ValidationError{Name: "STATUSDATA", err: fmt.Errorf("ent: validator failed for field \"STATUSDATA\": %w", err)}
		}
	}
	var (
		err  error
		node *RoomStatus
	)
	if len(rsc.hooks) == 0 {
		node, err = rsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rsc.mutation = mutation
			node, err = rsc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rsc.hooks) - 1; i >= 0; i-- {
			mut = rsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rsc *RoomStatusCreate) SaveX(ctx context.Context) *RoomStatus {
	v, err := rsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rsc *RoomStatusCreate) sqlSave(ctx context.Context) (*RoomStatus, error) {
	rs, _spec := rsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rsc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	rs.ID = int(id)
	return rs, nil
}

func (rsc *RoomStatusCreate) createSpec() (*RoomStatus, *sqlgraph.CreateSpec) {
	var (
		rs    = &RoomStatus{config: rsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: roomstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomstatus.FieldID,
			},
		}
	)
	if value, ok := rsc.mutation.STATUSDATA(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomstatus.FieldSTATUSDATA,
		})
		rs.STATUSDATA = value
	}
	if nodes := rsc.mutation.STATUSROOMIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roomstatus.STATUSROOMTable,
			Columns: []string{roomstatus.STATUSROOMColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return rs, _spec
}
