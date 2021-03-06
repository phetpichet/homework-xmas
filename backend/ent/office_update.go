// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Piichet/app/ent/office"
	"github.com/Piichet/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// OfficeUpdate is the builder for updating Office entities.
type OfficeUpdate struct {
	config
	hooks      []Hook
	mutation   *OfficeMutation
	predicates []predicate.Office
}

// Where adds a new predicate for the builder.
func (ou *OfficeUpdate) Where(ps ...predicate.Office) *OfficeUpdate {
	ou.predicates = append(ou.predicates, ps...)
	return ou
}

// SetFname sets the fname field.
func (ou *OfficeUpdate) SetFname(s string) *OfficeUpdate {
	ou.mutation.SetFname(s)
	return ou
}

// Mutation returns the OfficeMutation object of the builder.
func (ou *OfficeUpdate) Mutation() *OfficeMutation {
	return ou.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ou *OfficeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ou.mutation.Fname(); ok {
		if err := office.FnameValidator(v); err != nil {
			return 0, &ValidationError{Name: "fname", err: fmt.Errorf("ent: validator failed for field \"fname\": %w", err)}
		}
	}
	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OfficeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OfficeUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OfficeUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OfficeUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OfficeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   office.Table,
			Columns: office.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: office.FieldID,
			},
		},
	}
	if ps := ou.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Fname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: office.FieldFname,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{office.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OfficeUpdateOne is the builder for updating a single Office entity.
type OfficeUpdateOne struct {
	config
	hooks    []Hook
	mutation *OfficeMutation
}

// SetFname sets the fname field.
func (ouo *OfficeUpdateOne) SetFname(s string) *OfficeUpdateOne {
	ouo.mutation.SetFname(s)
	return ouo
}

// Mutation returns the OfficeMutation object of the builder.
func (ouo *OfficeUpdateOne) Mutation() *OfficeMutation {
	return ouo.mutation
}

// Save executes the query and returns the updated entity.
func (ouo *OfficeUpdateOne) Save(ctx context.Context) (*Office, error) {
	if v, ok := ouo.mutation.Fname(); ok {
		if err := office.FnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "fname", err: fmt.Errorf("ent: validator failed for field \"fname\": %w", err)}
		}
	}
	var (
		err  error
		node *Office
	)
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OfficeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OfficeUpdateOne) SaveX(ctx context.Context) *Office {
	o, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return o
}

// Exec executes the query on the entity.
func (ouo *OfficeUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OfficeUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OfficeUpdateOne) sqlSave(ctx context.Context) (o *Office, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   office.Table,
			Columns: office.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: office.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Office.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ouo.mutation.Fname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: office.FieldFname,
		})
	}
	o = &Office{config: ouo.config}
	_spec.Assign = o.assignValues
	_spec.ScanValues = o.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{office.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return o, nil
}
