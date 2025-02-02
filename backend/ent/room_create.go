// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/m16_z/app/ent/book"
	"github.com/m16_z/app/ent/room"
	"github.com/m16_z/app/ent/roominfo"
	"github.com/m16_z/app/ent/roomstatus"
	"github.com/m16_z/app/ent/roomtype"
)

// RoomCreate is the builder for creating a Room entity.
type RoomCreate struct {
	config
	mutation *RoomMutation
	hooks    []Hook
}

// SetROOMNAME sets the ROOMNAME field.
func (rc *RoomCreate) SetROOMNAME(s string) *RoomCreate {
	rc.mutation.SetROOMNAME(s)
	return rc
}

// AddROOMBOOKIDs adds the ROOM_BOOK edge to Book by ids.
func (rc *RoomCreate) AddROOMBOOKIDs(ids ...int) *RoomCreate {
	rc.mutation.AddROOMBOOKIDs(ids...)
	return rc
}

// AddROOMBOOK adds the ROOM_BOOK edges to Book.
func (rc *RoomCreate) AddROOMBOOK(b ...*Book) *RoomCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return rc.AddROOMBOOKIDs(ids...)
}

// SetROOMROOMTYPEID sets the ROOM_ROOMTYPE edge to RoomType by id.
func (rc *RoomCreate) SetROOMROOMTYPEID(id int) *RoomCreate {
	rc.mutation.SetROOMROOMTYPEID(id)
	return rc
}

// SetNillableROOMROOMTYPEID sets the ROOM_ROOMTYPE edge to RoomType by id if the given value is not nil.
func (rc *RoomCreate) SetNillableROOMROOMTYPEID(id *int) *RoomCreate {
	if id != nil {
		rc = rc.SetROOMROOMTYPEID(*id)
	}
	return rc
}

// SetROOMROOMTYPE sets the ROOM_ROOMTYPE edge to RoomType.
func (rc *RoomCreate) SetROOMROOMTYPE(r *RoomType) *RoomCreate {
	return rc.SetROOMROOMTYPEID(r.ID)
}

// SetROOMSTATUSID sets the ROOM_STATUS edge to RoomStatus by id.
func (rc *RoomCreate) SetROOMSTATUSID(id int) *RoomCreate {
	rc.mutation.SetROOMSTATUSID(id)
	return rc
}

// SetNillableROOMSTATUSID sets the ROOM_STATUS edge to RoomStatus by id if the given value is not nil.
func (rc *RoomCreate) SetNillableROOMSTATUSID(id *int) *RoomCreate {
	if id != nil {
		rc = rc.SetROOMSTATUSID(*id)
	}
	return rc
}

// SetROOMSTATUS sets the ROOM_STATUS edge to RoomStatus.
func (rc *RoomCreate) SetROOMSTATUS(r *RoomStatus) *RoomCreate {
	return rc.SetROOMSTATUSID(r.ID)
}

// SetROOMINFOID sets the ROOM_INFO edge to RoomInfo by id.
func (rc *RoomCreate) SetROOMINFOID(id int) *RoomCreate {
	rc.mutation.SetROOMINFOID(id)
	return rc
}

// SetNillableROOMINFOID sets the ROOM_INFO edge to RoomInfo by id if the given value is not nil.
func (rc *RoomCreate) SetNillableROOMINFOID(id *int) *RoomCreate {
	if id != nil {
		rc = rc.SetROOMINFOID(*id)
	}
	return rc
}

// SetROOMINFO sets the ROOM_INFO edge to RoomInfo.
func (rc *RoomCreate) SetROOMINFO(r *RoomInfo) *RoomCreate {
	return rc.SetROOMINFOID(r.ID)
}

// Mutation returns the RoomMutation object of the builder.
func (rc *RoomCreate) Mutation() *RoomMutation {
	return rc.mutation
}

// Save creates the Room in the database.
func (rc *RoomCreate) Save(ctx context.Context) (*Room, error) {
	if _, ok := rc.mutation.ROOMNAME(); !ok {
		return nil, &ValidationError{Name: "ROOMNAME", err: errors.New("ent: missing required field \"ROOMNAME\"")}
	}
	if v, ok := rc.mutation.ROOMNAME(); ok {
		if err := room.ROOMNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "ROOMNAME", err: fmt.Errorf("ent: validator failed for field \"ROOMNAME\": %w", err)}
		}
	}
	var (
		err  error
		node *Room
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoomCreate) SaveX(ctx context.Context) *Room {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RoomCreate) sqlSave(ctx context.Context) (*Room, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}

func (rc *RoomCreate) createSpec() (*Room, *sqlgraph.CreateSpec) {
	var (
		r     = &Room{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: room.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: room.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.ROOMNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldROOMNAME,
		})
		r.ROOMNAME = value
	}
	if nodes := rc.mutation.ROOMBOOKIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.ROOMBOOKTable,
			Columns: []string{room.ROOMBOOKColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: book.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ROOMROOMTYPEIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.ROOMROOMTYPETable,
			Columns: []string{room.ROOMROOMTYPEColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ROOMSTATUSIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.ROOMSTATUSTable,
			Columns: []string{room.ROOMSTATUSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ROOMINFOIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.ROOMINFOTable,
			Columns: []string{room.ROOMINFOColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roominfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return r, _spec
}
