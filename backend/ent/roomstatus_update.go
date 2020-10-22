// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/m16_z/app/ent/predicate"
	"github.com/m16_z/app/ent/room"
	"github.com/m16_z/app/ent/roomstatus"
)

// RoomStatusUpdate is the builder for updating RoomStatus entities.
type RoomStatusUpdate struct {
	config
	hooks      []Hook
	mutation   *RoomStatusMutation
	predicates []predicate.RoomStatus
}

// Where adds a new predicate for the builder.
func (rsu *RoomStatusUpdate) Where(ps ...predicate.RoomStatus) *RoomStatusUpdate {
	rsu.predicates = append(rsu.predicates, ps...)
	return rsu
}

// SetSTATUSDATA sets the STATUSDATA field.
func (rsu *RoomStatusUpdate) SetSTATUSDATA(s string) *RoomStatusUpdate {
	rsu.mutation.SetSTATUSDATA(s)
	return rsu
}

// AddSTATUSROOMIDs adds the STATUS_ROOM edge to Room by ids.
func (rsu *RoomStatusUpdate) AddSTATUSROOMIDs(ids ...int) *RoomStatusUpdate {
	rsu.mutation.AddSTATUSROOMIDs(ids...)
	return rsu
}

// AddSTATUSROOM adds the STATUS_ROOM edges to Room.
func (rsu *RoomStatusUpdate) AddSTATUSROOM(r ...*Room) *RoomStatusUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rsu.AddSTATUSROOMIDs(ids...)
}

// Mutation returns the RoomStatusMutation object of the builder.
func (rsu *RoomStatusUpdate) Mutation() *RoomStatusMutation {
	return rsu.mutation
}

// RemoveSTATUSROOMIDs removes the STATUS_ROOM edge to Room by ids.
func (rsu *RoomStatusUpdate) RemoveSTATUSROOMIDs(ids ...int) *RoomStatusUpdate {
	rsu.mutation.RemoveSTATUSROOMIDs(ids...)
	return rsu
}

// RemoveSTATUSROOM removes STATUS_ROOM edges to Room.
func (rsu *RoomStatusUpdate) RemoveSTATUSROOM(r ...*Room) *RoomStatusUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rsu.RemoveSTATUSROOMIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (rsu *RoomStatusUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := rsu.mutation.STATUSDATA(); ok {
		if err := roomstatus.STATUSDATAValidator(v); err != nil {
			return 0, &ValidationError{Name: "STATUSDATA", err: fmt.Errorf("ent: validator failed for field \"STATUSDATA\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(rsu.hooks) == 0 {
		affected, err = rsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rsu.mutation = mutation
			affected, err = rsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rsu.hooks) - 1; i >= 0; i-- {
			mut = rsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rsu *RoomStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := rsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rsu *RoomStatusUpdate) Exec(ctx context.Context) error {
	_, err := rsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsu *RoomStatusUpdate) ExecX(ctx context.Context) {
	if err := rsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rsu *RoomStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roomstatus.Table,
			Columns: roomstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomstatus.FieldID,
			},
		},
	}
	if ps := rsu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rsu.mutation.STATUSDATA(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomstatus.FieldSTATUSDATA,
		})
	}
	if nodes := rsu.mutation.RemovedSTATUSROOMIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsu.mutation.STATUSROOMIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roomstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoomStatusUpdateOne is the builder for updating a single RoomStatus entity.
type RoomStatusUpdateOne struct {
	config
	hooks    []Hook
	mutation *RoomStatusMutation
}

// SetSTATUSDATA sets the STATUSDATA field.
func (rsuo *RoomStatusUpdateOne) SetSTATUSDATA(s string) *RoomStatusUpdateOne {
	rsuo.mutation.SetSTATUSDATA(s)
	return rsuo
}

// AddSTATUSROOMIDs adds the STATUS_ROOM edge to Room by ids.
func (rsuo *RoomStatusUpdateOne) AddSTATUSROOMIDs(ids ...int) *RoomStatusUpdateOne {
	rsuo.mutation.AddSTATUSROOMIDs(ids...)
	return rsuo
}

// AddSTATUSROOM adds the STATUS_ROOM edges to Room.
func (rsuo *RoomStatusUpdateOne) AddSTATUSROOM(r ...*Room) *RoomStatusUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rsuo.AddSTATUSROOMIDs(ids...)
}

// Mutation returns the RoomStatusMutation object of the builder.
func (rsuo *RoomStatusUpdateOne) Mutation() *RoomStatusMutation {
	return rsuo.mutation
}

// RemoveSTATUSROOMIDs removes the STATUS_ROOM edge to Room by ids.
func (rsuo *RoomStatusUpdateOne) RemoveSTATUSROOMIDs(ids ...int) *RoomStatusUpdateOne {
	rsuo.mutation.RemoveSTATUSROOMIDs(ids...)
	return rsuo
}

// RemoveSTATUSROOM removes STATUS_ROOM edges to Room.
func (rsuo *RoomStatusUpdateOne) RemoveSTATUSROOM(r ...*Room) *RoomStatusUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rsuo.RemoveSTATUSROOMIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (rsuo *RoomStatusUpdateOne) Save(ctx context.Context) (*RoomStatus, error) {
	if v, ok := rsuo.mutation.STATUSDATA(); ok {
		if err := roomstatus.STATUSDATAValidator(v); err != nil {
			return nil, &ValidationError{Name: "STATUSDATA", err: fmt.Errorf("ent: validator failed for field \"STATUSDATA\": %w", err)}
		}
	}

	var (
		err  error
		node *RoomStatus
	)
	if len(rsuo.hooks) == 0 {
		node, err = rsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rsuo.mutation = mutation
			node, err = rsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rsuo.hooks) - 1; i >= 0; i-- {
			mut = rsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rsuo *RoomStatusUpdateOne) SaveX(ctx context.Context) *RoomStatus {
	rs, err := rsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return rs
}

// Exec executes the query on the entity.
func (rsuo *RoomStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := rsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsuo *RoomStatusUpdateOne) ExecX(ctx context.Context) {
	if err := rsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rsuo *RoomStatusUpdateOne) sqlSave(ctx context.Context) (rs *RoomStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roomstatus.Table,
			Columns: roomstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomstatus.FieldID,
			},
		},
	}
	id, ok := rsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing RoomStatus.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := rsuo.mutation.STATUSDATA(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomstatus.FieldSTATUSDATA,
		})
	}
	if nodes := rsuo.mutation.RemovedSTATUSROOMIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsuo.mutation.STATUSROOMIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	rs = &RoomStatus{config: rsuo.config}
	_spec.Assign = rs.assignValues
	_spec.ScanValues = rs.scanValues()
	if err = sqlgraph.UpdateNode(ctx, rsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roomstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return rs, nil
}
