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
	"github.com/m16_z/app/ent/roomtype"
)

// RoomTypeUpdate is the builder for updating RoomType entities.
type RoomTypeUpdate struct {
	config
	hooks      []Hook
	mutation   *RoomTypeMutation
	predicates []predicate.RoomType
}

// Where adds a new predicate for the builder.
func (rtu *RoomTypeUpdate) Where(ps ...predicate.RoomType) *RoomTypeUpdate {
	rtu.predicates = append(rtu.predicates, ps...)
	return rtu
}

// SetROOMTYPEDATA sets the ROOMTYPEDATA field.
func (rtu *RoomTypeUpdate) SetROOMTYPEDATA(s string) *RoomTypeUpdate {
	rtu.mutation.SetROOMTYPEDATA(s)
	return rtu
}

// SetCOSTDATA sets the COSTDATA field.
func (rtu *RoomTypeUpdate) SetCOSTDATA(i int) *RoomTypeUpdate {
	rtu.mutation.ResetCOSTDATA()
	rtu.mutation.SetCOSTDATA(i)
	return rtu
}

// AddCOSTDATA adds i to COSTDATA.
func (rtu *RoomTypeUpdate) AddCOSTDATA(i int) *RoomTypeUpdate {
	rtu.mutation.AddCOSTDATA(i)
	return rtu
}

// AddROOMTYPEROOMIDs adds the ROOMTYPE_ROOM edge to Room by ids.
func (rtu *RoomTypeUpdate) AddROOMTYPEROOMIDs(ids ...int) *RoomTypeUpdate {
	rtu.mutation.AddROOMTYPEROOMIDs(ids...)
	return rtu
}

// AddROOMTYPEROOM adds the ROOMTYPE_ROOM edges to Room.
func (rtu *RoomTypeUpdate) AddROOMTYPEROOM(r ...*Room) *RoomTypeUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtu.AddROOMTYPEROOMIDs(ids...)
}

// Mutation returns the RoomTypeMutation object of the builder.
func (rtu *RoomTypeUpdate) Mutation() *RoomTypeMutation {
	return rtu.mutation
}

// RemoveROOMTYPEROOMIDs removes the ROOMTYPE_ROOM edge to Room by ids.
func (rtu *RoomTypeUpdate) RemoveROOMTYPEROOMIDs(ids ...int) *RoomTypeUpdate {
	rtu.mutation.RemoveROOMTYPEROOMIDs(ids...)
	return rtu
}

// RemoveROOMTYPEROOM removes ROOMTYPE_ROOM edges to Room.
func (rtu *RoomTypeUpdate) RemoveROOMTYPEROOM(r ...*Room) *RoomTypeUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtu.RemoveROOMTYPEROOMIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (rtu *RoomTypeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := rtu.mutation.ROOMTYPEDATA(); ok {
		if err := roomtype.ROOMTYPEDATAValidator(v); err != nil {
			return 0, &ValidationError{Name: "ROOMTYPEDATA", err: fmt.Errorf("ent: validator failed for field \"ROOMTYPEDATA\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(rtu.hooks) == 0 {
		affected, err = rtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rtu.mutation = mutation
			affected, err = rtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rtu.hooks) - 1; i >= 0; i-- {
			mut = rtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtu *RoomTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := rtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rtu *RoomTypeUpdate) Exec(ctx context.Context) error {
	_, err := rtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtu *RoomTypeUpdate) ExecX(ctx context.Context) {
	if err := rtu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rtu *RoomTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roomtype.Table,
			Columns: roomtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomtype.FieldID,
			},
		},
	}
	if ps := rtu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtu.mutation.ROOMTYPEDATA(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomtype.FieldROOMTYPEDATA,
		})
	}
	if value, ok := rtu.mutation.COSTDATA(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roomtype.FieldCOSTDATA,
		})
	}
	if value, ok := rtu.mutation.AddedCOSTDATA(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roomtype.FieldCOSTDATA,
		})
	}
	if nodes := rtu.mutation.RemovedROOMTYPEROOMIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roomtype.ROOMTYPEROOMTable,
			Columns: []string{roomtype.ROOMTYPEROOMColumn},
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
	if nodes := rtu.mutation.ROOMTYPEROOMIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roomtype.ROOMTYPEROOMTable,
			Columns: []string{roomtype.ROOMTYPEROOMColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, rtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roomtype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoomTypeUpdateOne is the builder for updating a single RoomType entity.
type RoomTypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *RoomTypeMutation
}

// SetROOMTYPEDATA sets the ROOMTYPEDATA field.
func (rtuo *RoomTypeUpdateOne) SetROOMTYPEDATA(s string) *RoomTypeUpdateOne {
	rtuo.mutation.SetROOMTYPEDATA(s)
	return rtuo
}

// SetCOSTDATA sets the COSTDATA field.
func (rtuo *RoomTypeUpdateOne) SetCOSTDATA(i int) *RoomTypeUpdateOne {
	rtuo.mutation.ResetCOSTDATA()
	rtuo.mutation.SetCOSTDATA(i)
	return rtuo
}

// AddCOSTDATA adds i to COSTDATA.
func (rtuo *RoomTypeUpdateOne) AddCOSTDATA(i int) *RoomTypeUpdateOne {
	rtuo.mutation.AddCOSTDATA(i)
	return rtuo
}

// AddROOMTYPEROOMIDs adds the ROOMTYPE_ROOM edge to Room by ids.
func (rtuo *RoomTypeUpdateOne) AddROOMTYPEROOMIDs(ids ...int) *RoomTypeUpdateOne {
	rtuo.mutation.AddROOMTYPEROOMIDs(ids...)
	return rtuo
}

// AddROOMTYPEROOM adds the ROOMTYPE_ROOM edges to Room.
func (rtuo *RoomTypeUpdateOne) AddROOMTYPEROOM(r ...*Room) *RoomTypeUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtuo.AddROOMTYPEROOMIDs(ids...)
}

// Mutation returns the RoomTypeMutation object of the builder.
func (rtuo *RoomTypeUpdateOne) Mutation() *RoomTypeMutation {
	return rtuo.mutation
}

// RemoveROOMTYPEROOMIDs removes the ROOMTYPE_ROOM edge to Room by ids.
func (rtuo *RoomTypeUpdateOne) RemoveROOMTYPEROOMIDs(ids ...int) *RoomTypeUpdateOne {
	rtuo.mutation.RemoveROOMTYPEROOMIDs(ids...)
	return rtuo
}

// RemoveROOMTYPEROOM removes ROOMTYPE_ROOM edges to Room.
func (rtuo *RoomTypeUpdateOne) RemoveROOMTYPEROOM(r ...*Room) *RoomTypeUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtuo.RemoveROOMTYPEROOMIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (rtuo *RoomTypeUpdateOne) Save(ctx context.Context) (*RoomType, error) {
	if v, ok := rtuo.mutation.ROOMTYPEDATA(); ok {
		if err := roomtype.ROOMTYPEDATAValidator(v); err != nil {
			return nil, &ValidationError{Name: "ROOMTYPEDATA", err: fmt.Errorf("ent: validator failed for field \"ROOMTYPEDATA\": %w", err)}
		}
	}

	var (
		err  error
		node *RoomType
	)
	if len(rtuo.hooks) == 0 {
		node, err = rtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rtuo.mutation = mutation
			node, err = rtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rtuo.hooks) - 1; i >= 0; i-- {
			mut = rtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtuo *RoomTypeUpdateOne) SaveX(ctx context.Context) *RoomType {
	rt, err := rtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return rt
}

// Exec executes the query on the entity.
func (rtuo *RoomTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := rtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtuo *RoomTypeUpdateOne) ExecX(ctx context.Context) {
	if err := rtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rtuo *RoomTypeUpdateOne) sqlSave(ctx context.Context) (rt *RoomType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roomtype.Table,
			Columns: roomtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomtype.FieldID,
			},
		},
	}
	id, ok := rtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing RoomType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := rtuo.mutation.ROOMTYPEDATA(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomtype.FieldROOMTYPEDATA,
		})
	}
	if value, ok := rtuo.mutation.COSTDATA(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roomtype.FieldCOSTDATA,
		})
	}
	if value, ok := rtuo.mutation.AddedCOSTDATA(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roomtype.FieldCOSTDATA,
		})
	}
	if nodes := rtuo.mutation.RemovedROOMTYPEROOMIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roomtype.ROOMTYPEROOMTable,
			Columns: []string{roomtype.ROOMTYPEROOMColumn},
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
	if nodes := rtuo.mutation.ROOMTYPEROOMIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roomtype.ROOMTYPEROOMTable,
			Columns: []string{roomtype.ROOMTYPEROOMColumn},
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
	rt = &RoomType{config: rtuo.config}
	_spec.Assign = rt.assignValues
	_spec.ScanValues = rt.scanValues()
	if err = sqlgraph.UpdateNode(ctx, rtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roomtype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return rt, nil
}
