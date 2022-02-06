// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/member-gentei/member-gentei/gentei/ent/guild"
	"github.com/member-gentei/member-gentei/gentei/ent/guildrole"
	"github.com/member-gentei/member-gentei/gentei/ent/predicate"
	"github.com/member-gentei/member-gentei/gentei/ent/user"
	"github.com/member-gentei/member-gentei/gentei/ent/youtubetalent"
)

// GuildQuery is the builder for querying Guild entities.
type GuildQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Guild
	// eager-loading edges.
	withMembers        *UserQuery
	withAdmins         *UserQuery
	withRoles          *GuildRoleQuery
	withYoutubeTalents *YouTubeTalentQuery
	modifiers          []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GuildQuery builder.
func (gq *GuildQuery) Where(ps ...predicate.Guild) *GuildQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GuildQuery) Limit(limit int) *GuildQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GuildQuery) Offset(offset int) *GuildQuery {
	gq.offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GuildQuery) Unique(unique bool) *GuildQuery {
	gq.unique = &unique
	return gq
}

// Order adds an order step to the query.
func (gq *GuildQuery) Order(o ...OrderFunc) *GuildQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryMembers chains the current query on the "members" edge.
func (gq *GuildQuery) QueryMembers() *UserQuery {
	query := &UserQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(guild.Table, guild.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, guild.MembersTable, guild.MembersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAdmins chains the current query on the "admins" edge.
func (gq *GuildQuery) QueryAdmins() *UserQuery {
	query := &UserQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(guild.Table, guild.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, guild.AdminsTable, guild.AdminsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoles chains the current query on the "roles" edge.
func (gq *GuildQuery) QueryRoles() *GuildRoleQuery {
	query := &GuildRoleQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(guild.Table, guild.FieldID, selector),
			sqlgraph.To(guildrole.Table, guildrole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, guild.RolesTable, guild.RolesColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryYoutubeTalents chains the current query on the "youtube_talents" edge.
func (gq *GuildQuery) QueryYoutubeTalents() *YouTubeTalentQuery {
	query := &YouTubeTalentQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(guild.Table, guild.FieldID, selector),
			sqlgraph.To(youtubetalent.Table, youtubetalent.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, guild.YoutubeTalentsTable, guild.YoutubeTalentsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Guild entity from the query.
// Returns a *NotFoundError when no Guild was found.
func (gq *GuildQuery) First(ctx context.Context) (*Guild, error) {
	nodes, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{guild.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GuildQuery) FirstX(ctx context.Context) *Guild {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Guild ID from the query.
// Returns a *NotFoundError when no Guild ID was found.
func (gq *GuildQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{guild.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GuildQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Guild entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Guild entity is not found.
// Returns a *NotFoundError when no Guild entities are found.
func (gq *GuildQuery) Only(ctx context.Context) (*Guild, error) {
	nodes, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{guild.Label}
	default:
		return nil, &NotSingularError{guild.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GuildQuery) OnlyX(ctx context.Context) *Guild {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Guild ID in the query.
// Returns a *NotSingularError when exactly one Guild ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gq *GuildQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{guild.Label}
	default:
		err = &NotSingularError{guild.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GuildQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Guilds.
func (gq *GuildQuery) All(ctx context.Context) ([]*Guild, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gq *GuildQuery) AllX(ctx context.Context) []*Guild {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Guild IDs.
func (gq *GuildQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := gq.Select(guild.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GuildQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GuildQuery) Count(ctx context.Context) (int, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GuildQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GuildQuery) Exist(ctx context.Context) (bool, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GuildQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GuildQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GuildQuery) Clone() *GuildQuery {
	if gq == nil {
		return nil
	}
	return &GuildQuery{
		config:             gq.config,
		limit:              gq.limit,
		offset:             gq.offset,
		order:              append([]OrderFunc{}, gq.order...),
		predicates:         append([]predicate.Guild{}, gq.predicates...),
		withMembers:        gq.withMembers.Clone(),
		withAdmins:         gq.withAdmins.Clone(),
		withRoles:          gq.withRoles.Clone(),
		withYoutubeTalents: gq.withYoutubeTalents.Clone(),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

// WithMembers tells the query-builder to eager-load the nodes that are connected to
// the "members" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GuildQuery) WithMembers(opts ...func(*UserQuery)) *GuildQuery {
	query := &UserQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withMembers = query
	return gq
}

// WithAdmins tells the query-builder to eager-load the nodes that are connected to
// the "admins" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GuildQuery) WithAdmins(opts ...func(*UserQuery)) *GuildQuery {
	query := &UserQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withAdmins = query
	return gq
}

// WithRoles tells the query-builder to eager-load the nodes that are connected to
// the "roles" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GuildQuery) WithRoles(opts ...func(*GuildRoleQuery)) *GuildQuery {
	query := &GuildRoleQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withRoles = query
	return gq
}

// WithYoutubeTalents tells the query-builder to eager-load the nodes that are connected to
// the "youtube_talents" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GuildQuery) WithYoutubeTalents(opts ...func(*YouTubeTalentQuery)) *GuildQuery {
	query := &YouTubeTalentQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withYoutubeTalents = query
	return gq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Guild.Query().
//		GroupBy(guild.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gq *GuildQuery) GroupBy(field string, fields ...string) *GuildGroupBy {
	group := &GuildGroupBy{config: gq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Guild.Query().
//		Select(guild.FieldName).
//		Scan(ctx, &v)
//
func (gq *GuildQuery) Select(fields ...string) *GuildSelect {
	gq.fields = append(gq.fields, fields...)
	return &GuildSelect{GuildQuery: gq}
}

func (gq *GuildQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gq.fields {
		if !guild.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GuildQuery) sqlAll(ctx context.Context) ([]*Guild, error) {
	var (
		nodes       = []*Guild{}
		_spec       = gq.querySpec()
		loadedTypes = [4]bool{
			gq.withMembers != nil,
			gq.withAdmins != nil,
			gq.withRoles != nil,
			gq.withYoutubeTalents != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Guild{config: gq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gq.withMembers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[uint64]*Guild, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Members = []*User{}
		}
		var (
			edgeids []uint64
			edges   = make(map[uint64][]*Guild)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   guild.MembersTable,
				Columns: guild.MembersPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(guild.MembersPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := uint64(eout.Int64)
				inValue := uint64(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, gq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "members": %w`, err)
		}
		query.Where(user.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "members" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Members = append(nodes[i].Edges.Members, n)
			}
		}
	}

	if query := gq.withAdmins; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[uint64]*Guild, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Admins = []*User{}
		}
		var (
			edgeids []uint64
			edges   = make(map[uint64][]*Guild)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   guild.AdminsTable,
				Columns: guild.AdminsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(guild.AdminsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := uint64(eout.Int64)
				inValue := uint64(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, gq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "admins": %w`, err)
		}
		query.Where(user.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "admins" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Admins = append(nodes[i].Edges.Admins, n)
			}
		}
	}

	if query := gq.withRoles; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uint64]*Guild)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Roles = []*GuildRole{}
		}
		query.withFKs = true
		query.Where(predicate.GuildRole(func(s *sql.Selector) {
			s.Where(sql.InValues(guild.RolesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.guild_roles
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "guild_roles" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "guild_roles" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Roles = append(node.Edges.Roles, n)
		}
	}

	if query := gq.withYoutubeTalents; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[uint64]*Guild, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.YoutubeTalents = []*YouTubeTalent{}
		}
		var (
			edgeids []string
			edges   = make(map[string][]*Guild)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   guild.YoutubeTalentsTable,
				Columns: guild.YoutubeTalentsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(guild.YoutubeTalentsPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullString)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullString)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := uint64(eout.Int64)
				inValue := ein.String
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, gq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "youtube_talents": %w`, err)
		}
		query.Where(youtubetalent.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "youtube_talents" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.YoutubeTalents = append(nodes[i].Edges.YoutubeTalents, n)
			}
		}
	}

	return nodes, nil
}

func (gq *GuildQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	_spec.Node.Columns = gq.fields
	if len(gq.fields) > 0 {
		_spec.Unique = gq.unique != nil && *gq.unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GuildQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gq *GuildQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   guild.Table,
			Columns: guild.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: guild.FieldID,
			},
		},
		From:   gq.sql,
		Unique: true,
	}
	if unique := gq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, guild.FieldID)
		for i := range fields {
			if fields[i] != guild.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GuildQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(guild.Table)
	columns := gq.fields
	if len(columns) == 0 {
		columns = guild.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.unique != nil && *gq.unique {
		selector.Distinct()
	}
	for _, m := range gq.modifiers {
		m(selector)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gq *GuildQuery) ForUpdate(opts ...sql.LockOption) *GuildQuery {
	if gq.driver.Dialect() == dialect.Postgres {
		gq.Unique(false)
	}
	gq.modifiers = append(gq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gq *GuildQuery) ForShare(opts ...sql.LockOption) *GuildQuery {
	if gq.driver.Dialect() == dialect.Postgres {
		gq.Unique(false)
	}
	gq.modifiers = append(gq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gq
}

// GuildGroupBy is the group-by builder for Guild entities.
type GuildGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GuildGroupBy) Aggregate(fns ...AggregateFunc) *GuildGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the group-by query and scans the result into the given value.
func (ggb *GuildGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ggb.path(ctx)
	if err != nil {
		return err
	}
	ggb.sql = query
	return ggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ggb *GuildGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GuildGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GuildGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ggb *GuildGroupBy) StringsX(ctx context.Context) []string {
	v, err := ggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GuildGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guild.Label}
	default:
		err = fmt.Errorf("ent: GuildGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ggb *GuildGroupBy) StringX(ctx context.Context) string {
	v, err := ggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GuildGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GuildGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ggb *GuildGroupBy) IntsX(ctx context.Context) []int {
	v, err := ggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GuildGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guild.Label}
	default:
		err = fmt.Errorf("ent: GuildGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ggb *GuildGroupBy) IntX(ctx context.Context) int {
	v, err := ggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GuildGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GuildGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ggb *GuildGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GuildGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guild.Label}
	default:
		err = fmt.Errorf("ent: GuildGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ggb *GuildGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GuildGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GuildGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ggb *GuildGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GuildGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guild.Label}
	default:
		err = fmt.Errorf("ent: GuildGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ggb *GuildGroupBy) BoolX(ctx context.Context) bool {
	v, err := ggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ggb *GuildGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ggb.fields {
		if !guild.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ggb *GuildGroupBy) sqlQuery() *sql.Selector {
	selector := ggb.sql.Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ggb.fields)+len(ggb.fns))
		for _, f := range ggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ggb.fields...)...)
}

// GuildSelect is the builder for selecting fields of Guild entities.
type GuildSelect struct {
	*GuildQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GuildSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	gs.sql = gs.GuildQuery.sqlQuery(ctx)
	return gs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gs *GuildSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gs *GuildSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GuildSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gs *GuildSelect) StringsX(ctx context.Context) []string {
	v, err := gs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gs *GuildSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guild.Label}
	default:
		err = fmt.Errorf("ent: GuildSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gs *GuildSelect) StringX(ctx context.Context) string {
	v, err := gs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gs *GuildSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GuildSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gs *GuildSelect) IntsX(ctx context.Context) []int {
	v, err := gs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gs *GuildSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guild.Label}
	default:
		err = fmt.Errorf("ent: GuildSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gs *GuildSelect) IntX(ctx context.Context) int {
	v, err := gs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gs *GuildSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GuildSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gs *GuildSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gs *GuildSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guild.Label}
	default:
		err = fmt.Errorf("ent: GuildSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gs *GuildSelect) Float64X(ctx context.Context) float64 {
	v, err := gs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gs *GuildSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GuildSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gs *GuildSelect) BoolsX(ctx context.Context) []bool {
	v, err := gs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gs *GuildSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guild.Label}
	default:
		err = fmt.Errorf("ent: GuildSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gs *GuildSelect) BoolX(ctx context.Context) bool {
	v, err := gs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gs *GuildSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gs.sql.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
