// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vorteil/direktiv/ent/cloudevents"
	"github.com/vorteil/direktiv/ent/predicate"
)

// CloudEventsDelete is the builder for deleting a CloudEvents entity.
type CloudEventsDelete struct {
	config
	hooks    []Hook
	mutation *CloudEventsMutation
}

// Where appends a list predicates to the CloudEventsDelete builder.
func (ced *CloudEventsDelete) Where(ps ...predicate.CloudEvents) *CloudEventsDelete {
	ced.mutation.Where(ps...)
	return ced
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ced *CloudEventsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ced.hooks) == 0 {
		affected, err = ced.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CloudEventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ced.mutation = mutation
			affected, err = ced.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ced.hooks) - 1; i >= 0; i-- {
			if ced.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ced.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ced.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ced *CloudEventsDelete) ExecX(ctx context.Context) int {
	n, err := ced.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ced *CloudEventsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: cloudevents.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: cloudevents.FieldID,
			},
		},
	}
	if ps := ced.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ced.driver, _spec)
}

// CloudEventsDeleteOne is the builder for deleting a single CloudEvents entity.
type CloudEventsDeleteOne struct {
	ced *CloudEventsDelete
}

// Exec executes the deletion query.
func (cedo *CloudEventsDeleteOne) Exec(ctx context.Context) error {
	n, err := cedo.ced.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cloudevents.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cedo *CloudEventsDeleteOne) ExecX(ctx context.Context) {
	cedo.ced.ExecX(ctx)
}