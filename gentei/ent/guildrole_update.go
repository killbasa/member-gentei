// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/member-gentei/member-gentei/gentei/ent/guild"
	"github.com/member-gentei/member-gentei/gentei/ent/guildrole"
	"github.com/member-gentei/member-gentei/gentei/ent/predicate"
	"github.com/member-gentei/member-gentei/gentei/ent/user"
)

// GuildRoleUpdate is the builder for updating GuildRole entities.
type GuildRoleUpdate struct {
	config
	hooks    []Hook
	mutation *GuildRoleMutation
}

// Where appends a list predicates to the GuildRoleUpdate builder.
func (gru *GuildRoleUpdate) Where(ps ...predicate.GuildRole) *GuildRoleUpdate {
	gru.mutation.Where(ps...)
	return gru
}

// SetName sets the "name" field.
func (gru *GuildRoleUpdate) SetName(s string) *GuildRoleUpdate {
	gru.mutation.SetName(s)
	return gru
}

// SetLastUpdated sets the "last_updated" field.
func (gru *GuildRoleUpdate) SetLastUpdated(t time.Time) *GuildRoleUpdate {
	gru.mutation.SetLastUpdated(t)
	return gru
}

// SetNillableLastUpdated sets the "last_updated" field if the given value is not nil.
func (gru *GuildRoleUpdate) SetNillableLastUpdated(t *time.Time) *GuildRoleUpdate {
	if t != nil {
		gru.SetLastUpdated(*t)
	}
	return gru
}

// SetGuildID sets the "guild" edge to the Guild entity by ID.
func (gru *GuildRoleUpdate) SetGuildID(id uint64) *GuildRoleUpdate {
	gru.mutation.SetGuildID(id)
	return gru
}

// SetGuild sets the "guild" edge to the Guild entity.
func (gru *GuildRoleUpdate) SetGuild(g *Guild) *GuildRoleUpdate {
	return gru.SetGuildID(g.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (gru *GuildRoleUpdate) AddUserIDs(ids ...uint64) *GuildRoleUpdate {
	gru.mutation.AddUserIDs(ids...)
	return gru
}

// AddUsers adds the "users" edges to the User entity.
func (gru *GuildRoleUpdate) AddUsers(u ...*User) *GuildRoleUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gru.AddUserIDs(ids...)
}

// Mutation returns the GuildRoleMutation object of the builder.
func (gru *GuildRoleUpdate) Mutation() *GuildRoleMutation {
	return gru.mutation
}

// ClearGuild clears the "guild" edge to the Guild entity.
func (gru *GuildRoleUpdate) ClearGuild() *GuildRoleUpdate {
	gru.mutation.ClearGuild()
	return gru
}

// ClearUsers clears all "users" edges to the User entity.
func (gru *GuildRoleUpdate) ClearUsers() *GuildRoleUpdate {
	gru.mutation.ClearUsers()
	return gru
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (gru *GuildRoleUpdate) RemoveUserIDs(ids ...uint64) *GuildRoleUpdate {
	gru.mutation.RemoveUserIDs(ids...)
	return gru
}

// RemoveUsers removes "users" edges to User entities.
func (gru *GuildRoleUpdate) RemoveUsers(u ...*User) *GuildRoleUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gru.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gru *GuildRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gru.hooks) == 0 {
		if err = gru.check(); err != nil {
			return 0, err
		}
		affected, err = gru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GuildRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gru.check(); err != nil {
				return 0, err
			}
			gru.mutation = mutation
			affected, err = gru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gru.hooks) - 1; i >= 0; i-- {
			if gru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gru *GuildRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := gru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gru *GuildRoleUpdate) Exec(ctx context.Context) error {
	_, err := gru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gru *GuildRoleUpdate) ExecX(ctx context.Context) {
	if err := gru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gru *GuildRoleUpdate) check() error {
	if _, ok := gru.mutation.GuildID(); gru.mutation.GuildCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"guild\"")
	}
	return nil
}

func (gru *GuildRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   guildrole.Table,
			Columns: guildrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: guildrole.FieldID,
			},
		},
	}
	if ps := gru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: guildrole.FieldName,
		})
	}
	if value, ok := gru.mutation.LastUpdated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: guildrole.FieldLastUpdated,
		})
	}
	if gru.mutation.GuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   guildrole.GuildTable,
			Columns: []string{guildrole.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   guildrole.GuildTable,
			Columns: []string{guildrole.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   guildrole.UsersTable,
			Columns: guildrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.RemovedUsersIDs(); len(nodes) > 0 && !gru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   guildrole.UsersTable,
			Columns: guildrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   guildrole.UsersTable,
			Columns: guildrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{guildrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GuildRoleUpdateOne is the builder for updating a single GuildRole entity.
type GuildRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GuildRoleMutation
}

// SetName sets the "name" field.
func (gruo *GuildRoleUpdateOne) SetName(s string) *GuildRoleUpdateOne {
	gruo.mutation.SetName(s)
	return gruo
}

// SetLastUpdated sets the "last_updated" field.
func (gruo *GuildRoleUpdateOne) SetLastUpdated(t time.Time) *GuildRoleUpdateOne {
	gruo.mutation.SetLastUpdated(t)
	return gruo
}

// SetNillableLastUpdated sets the "last_updated" field if the given value is not nil.
func (gruo *GuildRoleUpdateOne) SetNillableLastUpdated(t *time.Time) *GuildRoleUpdateOne {
	if t != nil {
		gruo.SetLastUpdated(*t)
	}
	return gruo
}

// SetGuildID sets the "guild" edge to the Guild entity by ID.
func (gruo *GuildRoleUpdateOne) SetGuildID(id uint64) *GuildRoleUpdateOne {
	gruo.mutation.SetGuildID(id)
	return gruo
}

// SetGuild sets the "guild" edge to the Guild entity.
func (gruo *GuildRoleUpdateOne) SetGuild(g *Guild) *GuildRoleUpdateOne {
	return gruo.SetGuildID(g.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (gruo *GuildRoleUpdateOne) AddUserIDs(ids ...uint64) *GuildRoleUpdateOne {
	gruo.mutation.AddUserIDs(ids...)
	return gruo
}

// AddUsers adds the "users" edges to the User entity.
func (gruo *GuildRoleUpdateOne) AddUsers(u ...*User) *GuildRoleUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gruo.AddUserIDs(ids...)
}

// Mutation returns the GuildRoleMutation object of the builder.
func (gruo *GuildRoleUpdateOne) Mutation() *GuildRoleMutation {
	return gruo.mutation
}

// ClearGuild clears the "guild" edge to the Guild entity.
func (gruo *GuildRoleUpdateOne) ClearGuild() *GuildRoleUpdateOne {
	gruo.mutation.ClearGuild()
	return gruo
}

// ClearUsers clears all "users" edges to the User entity.
func (gruo *GuildRoleUpdateOne) ClearUsers() *GuildRoleUpdateOne {
	gruo.mutation.ClearUsers()
	return gruo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (gruo *GuildRoleUpdateOne) RemoveUserIDs(ids ...uint64) *GuildRoleUpdateOne {
	gruo.mutation.RemoveUserIDs(ids...)
	return gruo
}

// RemoveUsers removes "users" edges to User entities.
func (gruo *GuildRoleUpdateOne) RemoveUsers(u ...*User) *GuildRoleUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gruo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gruo *GuildRoleUpdateOne) Select(field string, fields ...string) *GuildRoleUpdateOne {
	gruo.fields = append([]string{field}, fields...)
	return gruo
}

// Save executes the query and returns the updated GuildRole entity.
func (gruo *GuildRoleUpdateOne) Save(ctx context.Context) (*GuildRole, error) {
	var (
		err  error
		node *GuildRole
	)
	if len(gruo.hooks) == 0 {
		if err = gruo.check(); err != nil {
			return nil, err
		}
		node, err = gruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GuildRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gruo.check(); err != nil {
				return nil, err
			}
			gruo.mutation = mutation
			node, err = gruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gruo.hooks) - 1; i >= 0; i-- {
			if gruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gruo *GuildRoleUpdateOne) SaveX(ctx context.Context) *GuildRole {
	node, err := gruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gruo *GuildRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := gruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gruo *GuildRoleUpdateOne) ExecX(ctx context.Context) {
	if err := gruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gruo *GuildRoleUpdateOne) check() error {
	if _, ok := gruo.mutation.GuildID(); gruo.mutation.GuildCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"guild\"")
	}
	return nil
}

func (gruo *GuildRoleUpdateOne) sqlSave(ctx context.Context) (_node *GuildRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   guildrole.Table,
			Columns: guildrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: guildrole.FieldID,
			},
		},
	}
	id, ok := gruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GuildRole.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := gruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, guildrole.FieldID)
		for _, f := range fields {
			if !guildrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != guildrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: guildrole.FieldName,
		})
	}
	if value, ok := gruo.mutation.LastUpdated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: guildrole.FieldLastUpdated,
		})
	}
	if gruo.mutation.GuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   guildrole.GuildTable,
			Columns: []string{guildrole.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   guildrole.GuildTable,
			Columns: []string{guildrole.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   guildrole.UsersTable,
			Columns: guildrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !gruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   guildrole.UsersTable,
			Columns: guildrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   guildrole.UsersTable,
			Columns: guildrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GuildRole{config: gruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{guildrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
