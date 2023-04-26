// Code generated by ent, DO NOT EDIT.

package ent

import (
	"airdrop-bot/ent/account"
	"airdrop-bot/ent/node"
	"airdrop-bot/ent/predicate"
	"airdrop-bot/ent/steprun"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StepRunUpdate is the builder for updating StepRun entities.
type StepRunUpdate struct {
	config
	hooks    []Hook
	mutation *StepRunMutation
}

// Where appends a list predicates to the StepRunUpdate builder.
func (sru *StepRunUpdate) Where(ps ...predicate.StepRun) *StepRunUpdate {
	sru.mutation.Where(ps...)
	return sru
}

// SetStep sets the "step" field.
func (sru *StepRunUpdate) SetStep(s string) *StepRunUpdate {
	sru.mutation.SetStep(s)
	return sru
}

// SetStatus sets the "status" field.
func (sru *StepRunUpdate) SetStatus(s steprun.Status) *StepRunUpdate {
	sru.mutation.SetStatus(s)
	return sru
}

// SetReason sets the "reason" field.
func (sru *StepRunUpdate) SetReason(s string) *StepRunUpdate {
	sru.mutation.SetReason(s)
	return sru
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (sru *StepRunUpdate) SetAccountID(id int) *StepRunUpdate {
	sru.mutation.SetAccountID(id)
	return sru
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (sru *StepRunUpdate) SetNillableAccountID(id *int) *StepRunUpdate {
	if id != nil {
		sru = sru.SetAccountID(*id)
	}
	return sru
}

// SetAccount sets the "account" edge to the Account entity.
func (sru *StepRunUpdate) SetAccount(a *Account) *StepRunUpdate {
	return sru.SetAccountID(a.ID)
}

// SetNodeID sets the "node" edge to the Node entity by ID.
func (sru *StepRunUpdate) SetNodeID(id int) *StepRunUpdate {
	sru.mutation.SetNodeID(id)
	return sru
}

// SetNillableNodeID sets the "node" edge to the Node entity by ID if the given value is not nil.
func (sru *StepRunUpdate) SetNillableNodeID(id *int) *StepRunUpdate {
	if id != nil {
		sru = sru.SetNodeID(*id)
	}
	return sru
}

// SetNode sets the "node" edge to the Node entity.
func (sru *StepRunUpdate) SetNode(n *Node) *StepRunUpdate {
	return sru.SetNodeID(n.ID)
}

// Mutation returns the StepRunMutation object of the builder.
func (sru *StepRunUpdate) Mutation() *StepRunMutation {
	return sru.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (sru *StepRunUpdate) ClearAccount() *StepRunUpdate {
	sru.mutation.ClearAccount()
	return sru
}

// ClearNode clears the "node" edge to the Node entity.
func (sru *StepRunUpdate) ClearNode() *StepRunUpdate {
	sru.mutation.ClearNode()
	return sru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *StepRunUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, StepRunMutation](ctx, sru.sqlSave, sru.mutation, sru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sru *StepRunUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *StepRunUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *StepRunUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sru *StepRunUpdate) check() error {
	if v, ok := sru.mutation.Status(); ok {
		if err := steprun.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "StepRun.status": %w`, err)}
		}
	}
	return nil
}

func (sru *StepRunUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(steprun.Table, steprun.Columns, sqlgraph.NewFieldSpec(steprun.FieldID, field.TypeInt))
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sru.mutation.Step(); ok {
		_spec.SetField(steprun.FieldStep, field.TypeString, value)
	}
	if value, ok := sru.mutation.Status(); ok {
		_spec.SetField(steprun.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := sru.mutation.Reason(); ok {
		_spec.SetField(steprun.FieldReason, field.TypeString, value)
	}
	if sru.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.AccountTable,
			Columns: []string{steprun.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.AccountTable,
			Columns: []string{steprun.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sru.mutation.NodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.NodeTable,
			Columns: []string{steprun.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.NodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.NodeTable,
			Columns: []string{steprun.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{steprun.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sru.mutation.done = true
	return n, nil
}

// StepRunUpdateOne is the builder for updating a single StepRun entity.
type StepRunUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StepRunMutation
}

// SetStep sets the "step" field.
func (sruo *StepRunUpdateOne) SetStep(s string) *StepRunUpdateOne {
	sruo.mutation.SetStep(s)
	return sruo
}

// SetStatus sets the "status" field.
func (sruo *StepRunUpdateOne) SetStatus(s steprun.Status) *StepRunUpdateOne {
	sruo.mutation.SetStatus(s)
	return sruo
}

// SetReason sets the "reason" field.
func (sruo *StepRunUpdateOne) SetReason(s string) *StepRunUpdateOne {
	sruo.mutation.SetReason(s)
	return sruo
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (sruo *StepRunUpdateOne) SetAccountID(id int) *StepRunUpdateOne {
	sruo.mutation.SetAccountID(id)
	return sruo
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (sruo *StepRunUpdateOne) SetNillableAccountID(id *int) *StepRunUpdateOne {
	if id != nil {
		sruo = sruo.SetAccountID(*id)
	}
	return sruo
}

// SetAccount sets the "account" edge to the Account entity.
func (sruo *StepRunUpdateOne) SetAccount(a *Account) *StepRunUpdateOne {
	return sruo.SetAccountID(a.ID)
}

// SetNodeID sets the "node" edge to the Node entity by ID.
func (sruo *StepRunUpdateOne) SetNodeID(id int) *StepRunUpdateOne {
	sruo.mutation.SetNodeID(id)
	return sruo
}

// SetNillableNodeID sets the "node" edge to the Node entity by ID if the given value is not nil.
func (sruo *StepRunUpdateOne) SetNillableNodeID(id *int) *StepRunUpdateOne {
	if id != nil {
		sruo = sruo.SetNodeID(*id)
	}
	return sruo
}

// SetNode sets the "node" edge to the Node entity.
func (sruo *StepRunUpdateOne) SetNode(n *Node) *StepRunUpdateOne {
	return sruo.SetNodeID(n.ID)
}

// Mutation returns the StepRunMutation object of the builder.
func (sruo *StepRunUpdateOne) Mutation() *StepRunMutation {
	return sruo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (sruo *StepRunUpdateOne) ClearAccount() *StepRunUpdateOne {
	sruo.mutation.ClearAccount()
	return sruo
}

// ClearNode clears the "node" edge to the Node entity.
func (sruo *StepRunUpdateOne) ClearNode() *StepRunUpdateOne {
	sruo.mutation.ClearNode()
	return sruo
}

// Where appends a list predicates to the StepRunUpdate builder.
func (sruo *StepRunUpdateOne) Where(ps ...predicate.StepRun) *StepRunUpdateOne {
	sruo.mutation.Where(ps...)
	return sruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *StepRunUpdateOne) Select(field string, fields ...string) *StepRunUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated StepRun entity.
func (sruo *StepRunUpdateOne) Save(ctx context.Context) (*StepRun, error) {
	return withHooks[*StepRun, StepRunMutation](ctx, sruo.sqlSave, sruo.mutation, sruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *StepRunUpdateOne) SaveX(ctx context.Context) *StepRun {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *StepRunUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *StepRunUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sruo *StepRunUpdateOne) check() error {
	if v, ok := sruo.mutation.Status(); ok {
		if err := steprun.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "StepRun.status": %w`, err)}
		}
	}
	return nil
}

func (sruo *StepRunUpdateOne) sqlSave(ctx context.Context) (_node *StepRun, err error) {
	if err := sruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(steprun.Table, steprun.Columns, sqlgraph.NewFieldSpec(steprun.FieldID, field.TypeInt))
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StepRun.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, steprun.FieldID)
		for _, f := range fields {
			if !steprun.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != steprun.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sruo.mutation.Step(); ok {
		_spec.SetField(steprun.FieldStep, field.TypeString, value)
	}
	if value, ok := sruo.mutation.Status(); ok {
		_spec.SetField(steprun.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := sruo.mutation.Reason(); ok {
		_spec.SetField(steprun.FieldReason, field.TypeString, value)
	}
	if sruo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.AccountTable,
			Columns: []string{steprun.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.AccountTable,
			Columns: []string{steprun.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sruo.mutation.NodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.NodeTable,
			Columns: []string{steprun.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.NodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.NodeTable,
			Columns: []string{steprun.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &StepRun{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{steprun.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sruo.mutation.done = true
	return _node, nil
}
