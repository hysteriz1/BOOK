// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/m16_z/app/ent/predicate"
	"github.com/m16_z/app/ent/userstatus"
)

// UserStatusDelete is the builder for deleting a UserStatus entity.
type UserStatusDelete struct {
	config
	hooks      []Hook
	mutation   *UserStatusMutation
	predicates []predicate.UserStatus
}

// Where adds a new predicate to the delete builder.
func (usd *UserStatusDelete) Where(ps ...predicate.UserStatus) *UserStatusDelete {
	usd.predicates = append(usd.predicates, ps...)
	return usd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (usd *UserStatusDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(usd.hooks) == 0 {
		affected, err = usd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			usd.mutation = mutation
			affected, err = usd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(usd.hooks) - 1; i >= 0; i-- {
			mut = usd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (usd *UserStatusDelete) ExecX(ctx context.Context) int {
	n, err := usd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (usd *UserStatusDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userstatus.FieldID,
			},
		},
	}
	if ps := usd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, usd.driver, _spec)
}

// UserStatusDeleteOne is the builder for deleting a single UserStatus entity.
type UserStatusDeleteOne struct {
	usd *UserStatusDelete
}

// Exec executes the deletion query.
func (usdo *UserStatusDeleteOne) Exec(ctx context.Context) error {
	n, err := usdo.usd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userstatus.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (usdo *UserStatusDeleteOne) ExecX(ctx context.Context) {
	usdo.usd.ExecX(ctx)
}
