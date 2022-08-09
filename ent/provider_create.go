// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/compsole/ent/competition"
	"github.com/BradHacker/compsole/ent/provider"
	"github.com/google/uuid"
)

// ProviderCreate is the builder for creating a Provider entity.
type ProviderCreate struct {
	config
	mutation *ProviderMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProviderCreate) SetName(s string) *ProviderCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetType sets the "type" field.
func (pc *ProviderCreate) SetType(s string) *ProviderCreate {
	pc.mutation.SetType(s)
	return pc
}

// SetConfig sets the "config" field.
func (pc *ProviderCreate) SetConfig(s string) *ProviderCreate {
	pc.mutation.SetConfig(s)
	return pc
}

// SetID sets the "id" field.
func (pc *ProviderCreate) SetID(u uuid.UUID) *ProviderCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *ProviderCreate) SetNillableID(u *uuid.UUID) *ProviderCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// AddProviderToCompetitionIDs adds the "ProviderToCompetition" edge to the Competition entity by IDs.
func (pc *ProviderCreate) AddProviderToCompetitionIDs(ids ...uuid.UUID) *ProviderCreate {
	pc.mutation.AddProviderToCompetitionIDs(ids...)
	return pc
}

// AddProviderToCompetition adds the "ProviderToCompetition" edges to the Competition entity.
func (pc *ProviderCreate) AddProviderToCompetition(c ...*Competition) *ProviderCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddProviderToCompetitionIDs(ids...)
}

// Mutation returns the ProviderMutation object of the builder.
func (pc *ProviderCreate) Mutation() *ProviderMutation {
	return pc.mutation
}

// Save creates the Provider in the database.
func (pc *ProviderCreate) Save(ctx context.Context) (*Provider, error) {
	var (
		err  error
		node *Provider
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProviderCreate) SaveX(ctx context.Context) *Provider {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProviderCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProviderCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProviderCreate) defaults() {
	if _, ok := pc.mutation.ID(); !ok {
		v := provider.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProviderCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Provider.name"`)}
	}
	if _, ok := pc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Provider.type"`)}
	}
	if _, ok := pc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New(`ent: missing required field "Provider.config"`)}
	}
	return nil
}

func (pc *ProviderCreate) sqlSave(ctx context.Context) (*Provider, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (pc *ProviderCreate) createSpec() (*Provider, *sqlgraph.CreateSpec) {
	var (
		_node = &Provider{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: provider.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: provider.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldType,
		})
		_node.Type = value
	}
	if value, ok := pc.mutation.Config(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldConfig,
		})
		_node.Config = value
	}
	if nodes := pc.mutation.ProviderToCompetitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provider.ProviderToCompetitionTable,
			Columns: []string{provider.ProviderToCompetitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: competition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProviderCreateBulk is the builder for creating many Provider entities in bulk.
type ProviderCreateBulk struct {
	config
	builders []*ProviderCreate
}

// Save creates the Provider entities in the database.
func (pcb *ProviderCreateBulk) Save(ctx context.Context) ([]*Provider, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Provider, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProviderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProviderCreateBulk) SaveX(ctx context.Context) []*Provider {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProviderCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
