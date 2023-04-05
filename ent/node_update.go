// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airdrop-bot/ent/node"
	"airdrop-bot/ent/predicate"
	"airdrop-bot/ent/steprun"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NodeUpdate is the builder for updating Node entities.
type NodeUpdate struct {
	config
	hooks    []Hook
	mutation *NodeMutation
}

// Where appends a list predicates to the NodeUpdate builder.
func (nu *NodeUpdate) Where(ps ...predicate.Node) *NodeUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetName sets the "name" field.
func (nu *NodeUpdate) SetName(s string) *NodeUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetRegion sets the "region" field.
func (nu *NodeUpdate) SetRegion(s string) *NodeUpdate {
	nu.mutation.SetRegion(s)
	return nu
}

// SetDnsName sets the "dnsName" field.
func (nu *NodeUpdate) SetDnsName(s string) *NodeUpdate {
	nu.mutation.SetDnsName(s)
	return nu
}

// SetAddress sets the "address" field.
func (nu *NodeUpdate) SetAddress(s string) *NodeUpdate {
	nu.mutation.SetAddress(s)
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NodeUpdate) SetUpdatedAt(t time.Time) *NodeUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// AddStepRunIDs adds the "step_runs" edge to the StepRun entity by IDs.
func (nu *NodeUpdate) AddStepRunIDs(ids ...int) *NodeUpdate {
	nu.mutation.AddStepRunIDs(ids...)
	return nu
}

// AddStepRuns adds the "step_runs" edges to the StepRun entity.
func (nu *NodeUpdate) AddStepRuns(s ...*StepRun) *NodeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nu.AddStepRunIDs(ids...)
}

// Mutation returns the NodeMutation object of the builder.
func (nu *NodeUpdate) Mutation() *NodeMutation {
	return nu.mutation
}

// ClearStepRuns clears all "step_runs" edges to the StepRun entity.
func (nu *NodeUpdate) ClearStepRuns() *NodeUpdate {
	nu.mutation.ClearStepRuns()
	return nu
}

// RemoveStepRunIDs removes the "step_runs" edge to StepRun entities by IDs.
func (nu *NodeUpdate) RemoveStepRunIDs(ids ...int) *NodeUpdate {
	nu.mutation.RemoveStepRunIDs(ids...)
	return nu
}

// RemoveStepRuns removes "step_runs" edges to StepRun entities.
func (nu *NodeUpdate) RemoveStepRuns(s ...*StepRun) *NodeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nu.RemoveStepRunIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nu.hooks) == 0 {
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			if nu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NodeUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NodeUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NodeUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   node.Table,
			Columns: node.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: node.FieldID,
			},
		},
	}
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldName,
		})
	}
	if value, ok := nu.mutation.Region(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldRegion,
		})
	}
	if value, ok := nu.mutation.DnsName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldDnsName,
		})
	}
	if value, ok := nu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldAddress,
		})
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: node.FieldUpdatedAt,
		})
	}
	if nu.mutation.StepRunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.StepRunsTable,
			Columns: []string{node.StepRunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: steprun.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedStepRunsIDs(); len(nodes) > 0 && !nu.mutation.StepRunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.StepRunsTable,
			Columns: []string{node.StepRunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: steprun.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.StepRunsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.StepRunsTable,
			Columns: []string{node.StepRunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: steprun.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NodeUpdateOne is the builder for updating a single Node entity.
type NodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NodeMutation
}

// SetName sets the "name" field.
func (nuo *NodeUpdateOne) SetName(s string) *NodeUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetRegion sets the "region" field.
func (nuo *NodeUpdateOne) SetRegion(s string) *NodeUpdateOne {
	nuo.mutation.SetRegion(s)
	return nuo
}

// SetDnsName sets the "dnsName" field.
func (nuo *NodeUpdateOne) SetDnsName(s string) *NodeUpdateOne {
	nuo.mutation.SetDnsName(s)
	return nuo
}

// SetAddress sets the "address" field.
func (nuo *NodeUpdateOne) SetAddress(s string) *NodeUpdateOne {
	nuo.mutation.SetAddress(s)
	return nuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NodeUpdateOne) SetUpdatedAt(t time.Time) *NodeUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// AddStepRunIDs adds the "step_runs" edge to the StepRun entity by IDs.
func (nuo *NodeUpdateOne) AddStepRunIDs(ids ...int) *NodeUpdateOne {
	nuo.mutation.AddStepRunIDs(ids...)
	return nuo
}

// AddStepRuns adds the "step_runs" edges to the StepRun entity.
func (nuo *NodeUpdateOne) AddStepRuns(s ...*StepRun) *NodeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nuo.AddStepRunIDs(ids...)
}

// Mutation returns the NodeMutation object of the builder.
func (nuo *NodeUpdateOne) Mutation() *NodeMutation {
	return nuo.mutation
}

// ClearStepRuns clears all "step_runs" edges to the StepRun entity.
func (nuo *NodeUpdateOne) ClearStepRuns() *NodeUpdateOne {
	nuo.mutation.ClearStepRuns()
	return nuo
}

// RemoveStepRunIDs removes the "step_runs" edge to StepRun entities by IDs.
func (nuo *NodeUpdateOne) RemoveStepRunIDs(ids ...int) *NodeUpdateOne {
	nuo.mutation.RemoveStepRunIDs(ids...)
	return nuo
}

// RemoveStepRuns removes "step_runs" edges to StepRun entities.
func (nuo *NodeUpdateOne) RemoveStepRuns(s ...*StepRun) *NodeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nuo.RemoveStepRunIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NodeUpdateOne) Select(field string, fields ...string) *NodeUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Node entity.
func (nuo *NodeUpdateOne) Save(ctx context.Context) (*Node, error) {
	var (
		err  error
		node *Node
	)
	if len(nuo.hooks) == 0 {
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			if nuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NodeUpdateOne) SaveX(ctx context.Context) *Node {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NodeUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NodeUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NodeUpdateOne) sqlSave(ctx context.Context) (_node *Node, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   node.Table,
			Columns: node.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: node.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Node.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, node.FieldID)
		for _, f := range fields {
			if !node.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != node.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldName,
		})
	}
	if value, ok := nuo.mutation.Region(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldRegion,
		})
	}
	if value, ok := nuo.mutation.DnsName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldDnsName,
		})
	}
	if value, ok := nuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldAddress,
		})
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: node.FieldUpdatedAt,
		})
	}
	if nuo.mutation.StepRunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.StepRunsTable,
			Columns: []string{node.StepRunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: steprun.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedStepRunsIDs(); len(nodes) > 0 && !nuo.mutation.StepRunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.StepRunsTable,
			Columns: []string{node.StepRunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: steprun.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.StepRunsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.StepRunsTable,
			Columns: []string{node.StepRunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: steprun.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Node{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
