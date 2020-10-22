// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/m16_z/app/ent/predicate"
	"github.com/m16_z/app/ent/user"
	"github.com/m16_z/app/ent/userstatus"
)

// UserStatusUpdate is the builder for updating UserStatus entities.
type UserStatusUpdate struct {
	config
	hooks      []Hook
	mutation   *UserStatusMutation
	predicates []predicate.UserStatus
}

// Where adds a new predicate for the builder.
func (usu *UserStatusUpdate) Where(ps ...predicate.UserStatus) *UserStatusUpdate {
	usu.predicates = append(usu.predicates, ps...)
	return usu
}

// SetSTATUS sets the STATUS field.
func (usu *UserStatusUpdate) SetSTATUS(s string) *UserStatusUpdate {
	usu.mutation.SetSTATUS(s)
	return usu
}

// AddUSERSTATUSUSERIDs adds the USERSTATUS_USER edge to User by ids.
func (usu *UserStatusUpdate) AddUSERSTATUSUSERIDs(ids ...int) *UserStatusUpdate {
	usu.mutation.AddUSERSTATUSUSERIDs(ids...)
	return usu
}

// AddUSERSTATUSUSER adds the USERSTATUS_USER edges to User.
func (usu *UserStatusUpdate) AddUSERSTATUSUSER(u ...*User) *UserStatusUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return usu.AddUSERSTATUSUSERIDs(ids...)
}

// Mutation returns the UserStatusMutation object of the builder.
func (usu *UserStatusUpdate) Mutation() *UserStatusMutation {
	return usu.mutation
}

// RemoveUSERSTATUSUSERIDs removes the USERSTATUS_USER edge to User by ids.
func (usu *UserStatusUpdate) RemoveUSERSTATUSUSERIDs(ids ...int) *UserStatusUpdate {
	usu.mutation.RemoveUSERSTATUSUSERIDs(ids...)
	return usu
}

// RemoveUSERSTATUSUSER removes USERSTATUS_USER edges to User.
func (usu *UserStatusUpdate) RemoveUSERSTATUSUSER(u ...*User) *UserStatusUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return usu.RemoveUSERSTATUSUSERIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (usu *UserStatusUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := usu.mutation.STATUS(); ok {
		if err := userstatus.STATUSValidator(v); err != nil {
			return 0, &ValidationError{Name: "STATUS", err: fmt.Errorf("ent: validator failed for field \"STATUS\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(usu.hooks) == 0 {
		affected, err = usu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			usu.mutation = mutation
			affected, err = usu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(usu.hooks) - 1; i >= 0; i-- {
			mut = usu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (usu *UserStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := usu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usu *UserStatusUpdate) Exec(ctx context.Context) error {
	_, err := usu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usu *UserStatusUpdate) ExecX(ctx context.Context) {
	if err := usu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (usu *UserStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userstatus.Table,
			Columns: userstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userstatus.FieldID,
			},
		},
	}
	if ps := usu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usu.mutation.STATUS(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userstatus.FieldSTATUS,
		})
	}
	if nodes := usu.mutation.RemovedUSERSTATUSUSERIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userstatus.USERSTATUSUSERTable,
			Columns: []string{userstatus.USERSTATUSUSERColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usu.mutation.USERSTATUSUSERIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userstatus.USERSTATUSUSERTable,
			Columns: []string{userstatus.USERSTATUSUSERColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, usu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserStatusUpdateOne is the builder for updating a single UserStatus entity.
type UserStatusUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserStatusMutation
}

// SetSTATUS sets the STATUS field.
func (usuo *UserStatusUpdateOne) SetSTATUS(s string) *UserStatusUpdateOne {
	usuo.mutation.SetSTATUS(s)
	return usuo
}

// AddUSERSTATUSUSERIDs adds the USERSTATUS_USER edge to User by ids.
func (usuo *UserStatusUpdateOne) AddUSERSTATUSUSERIDs(ids ...int) *UserStatusUpdateOne {
	usuo.mutation.AddUSERSTATUSUSERIDs(ids...)
	return usuo
}

// AddUSERSTATUSUSER adds the USERSTATUS_USER edges to User.
func (usuo *UserStatusUpdateOne) AddUSERSTATUSUSER(u ...*User) *UserStatusUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return usuo.AddUSERSTATUSUSERIDs(ids...)
}

// Mutation returns the UserStatusMutation object of the builder.
func (usuo *UserStatusUpdateOne) Mutation() *UserStatusMutation {
	return usuo.mutation
}

// RemoveUSERSTATUSUSERIDs removes the USERSTATUS_USER edge to User by ids.
func (usuo *UserStatusUpdateOne) RemoveUSERSTATUSUSERIDs(ids ...int) *UserStatusUpdateOne {
	usuo.mutation.RemoveUSERSTATUSUSERIDs(ids...)
	return usuo
}

// RemoveUSERSTATUSUSER removes USERSTATUS_USER edges to User.
func (usuo *UserStatusUpdateOne) RemoveUSERSTATUSUSER(u ...*User) *UserStatusUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return usuo.RemoveUSERSTATUSUSERIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (usuo *UserStatusUpdateOne) Save(ctx context.Context) (*UserStatus, error) {
	if v, ok := usuo.mutation.STATUS(); ok {
		if err := userstatus.STATUSValidator(v); err != nil {
			return nil, &ValidationError{Name: "STATUS", err: fmt.Errorf("ent: validator failed for field \"STATUS\": %w", err)}
		}
	}

	var (
		err  error
		node *UserStatus
	)
	if len(usuo.hooks) == 0 {
		node, err = usuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			usuo.mutation = mutation
			node, err = usuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(usuo.hooks) - 1; i >= 0; i-- {
			mut = usuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (usuo *UserStatusUpdateOne) SaveX(ctx context.Context) *UserStatus {
	us, err := usuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return us
}

// Exec executes the query on the entity.
func (usuo *UserStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := usuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usuo *UserStatusUpdateOne) ExecX(ctx context.Context) {
	if err := usuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (usuo *UserStatusUpdateOne) sqlSave(ctx context.Context) (us *UserStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userstatus.Table,
			Columns: userstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userstatus.FieldID,
			},
		},
	}
	id, ok := usuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserStatus.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := usuo.mutation.STATUS(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userstatus.FieldSTATUS,
		})
	}
	if nodes := usuo.mutation.RemovedUSERSTATUSUSERIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userstatus.USERSTATUSUSERTable,
			Columns: []string{userstatus.USERSTATUSUSERColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usuo.mutation.USERSTATUSUSERIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userstatus.USERSTATUSUSERTable,
			Columns: []string{userstatus.USERSTATUSUSERColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	us = &UserStatus{config: usuo.config}
	_spec.Assign = us.assignValues
	_spec.ScanValues = us.scanValues()
	if err = sqlgraph.UpdateNode(ctx, usuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return us, nil
}
