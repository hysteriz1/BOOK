// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/m16_z/app/ent/predicate"
	"github.com/m16_z/app/ent/room"
	"github.com/m16_z/app/ent/roomstatus"
)

// RoomStatusQuery is the builder for querying RoomStatus entities.
type RoomStatusQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.RoomStatus
	// eager-loading edges.
	withSTATUSROOM *RoomQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (rsq *RoomStatusQuery) Where(ps ...predicate.RoomStatus) *RoomStatusQuery {
	rsq.predicates = append(rsq.predicates, ps...)
	return rsq
}

// Limit adds a limit step to the query.
func (rsq *RoomStatusQuery) Limit(limit int) *RoomStatusQuery {
	rsq.limit = &limit
	return rsq
}

// Offset adds an offset step to the query.
func (rsq *RoomStatusQuery) Offset(offset int) *RoomStatusQuery {
	rsq.offset = &offset
	return rsq
}

// Order adds an order step to the query.
func (rsq *RoomStatusQuery) Order(o ...OrderFunc) *RoomStatusQuery {
	rsq.order = append(rsq.order, o...)
	return rsq
}

// QuerySTATUSROOM chains the current query on the STATUS_ROOM edge.
func (rsq *RoomStatusQuery) QuerySTATUSROOM() *RoomQuery {
	query := &RoomQuery{config: rsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roomstatus.Table, roomstatus.FieldID, rsq.sqlQuery()),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, roomstatus.STATUSROOMTable, roomstatus.STATUSROOMColumn),
		)
		fromU = sqlgraph.SetNeighbors(rsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RoomStatus entity in the query. Returns *NotFoundError when no roomstatus was found.
func (rsq *RoomStatusQuery) First(ctx context.Context) (*RoomStatus, error) {
	rsSlice, err := rsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rsSlice) == 0 {
		return nil, &NotFoundError{roomstatus.Label}
	}
	return rsSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rsq *RoomStatusQuery) FirstX(ctx context.Context) *RoomStatus {
	rs, err := rsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return rs
}

// FirstID returns the first RoomStatus id in the query. Returns *NotFoundError when no id was found.
func (rsq *RoomStatusQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{roomstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (rsq *RoomStatusQuery) FirstXID(ctx context.Context) int {
	id, err := rsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only RoomStatus entity in the query, returns an error if not exactly one entity was returned.
func (rsq *RoomStatusQuery) Only(ctx context.Context) (*RoomStatus, error) {
	rsSlice, err := rsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(rsSlice) {
	case 1:
		return rsSlice[0], nil
	case 0:
		return nil, &NotFoundError{roomstatus.Label}
	default:
		return nil, &NotSingularError{roomstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rsq *RoomStatusQuery) OnlyX(ctx context.Context) *RoomStatus {
	rs, err := rsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return rs
}

// OnlyID returns the only RoomStatus id in the query, returns an error if not exactly one id was returned.
func (rsq *RoomStatusQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{roomstatus.Label}
	default:
		err = &NotSingularError{roomstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rsq *RoomStatusQuery) OnlyIDX(ctx context.Context) int {
	id, err := rsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RoomStatusSlice.
func (rsq *RoomStatusQuery) All(ctx context.Context) ([]*RoomStatus, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rsq *RoomStatusQuery) AllX(ctx context.Context) []*RoomStatus {
	rsSlice, err := rsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return rsSlice
}

// IDs executes the query and returns a list of RoomStatus ids.
func (rsq *RoomStatusQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rsq.Select(roomstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rsq *RoomStatusQuery) IDsX(ctx context.Context) []int {
	ids, err := rsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rsq *RoomStatusQuery) Count(ctx context.Context) (int, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rsq *RoomStatusQuery) CountX(ctx context.Context) int {
	count, err := rsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rsq *RoomStatusQuery) Exist(ctx context.Context) (bool, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rsq *RoomStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := rsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rsq *RoomStatusQuery) Clone() *RoomStatusQuery {
	return &RoomStatusQuery{
		config:     rsq.config,
		limit:      rsq.limit,
		offset:     rsq.offset,
		order:      append([]OrderFunc{}, rsq.order...),
		unique:     append([]string{}, rsq.unique...),
		predicates: append([]predicate.RoomStatus{}, rsq.predicates...),
		// clone intermediate query.
		sql:  rsq.sql.Clone(),
		path: rsq.path,
	}
}

//  WithSTATUSROOM tells the query-builder to eager-loads the nodes that are connected to
// the "STATUS_ROOM" edge. The optional arguments used to configure the query builder of the edge.
func (rsq *RoomStatusQuery) WithSTATUSROOM(opts ...func(*RoomQuery)) *RoomStatusQuery {
	query := &RoomQuery{config: rsq.config}
	for _, opt := range opts {
		opt(query)
	}
	rsq.withSTATUSROOM = query
	return rsq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		STATUSDATA string `json:"STATUSDATA,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RoomStatus.Query().
//		GroupBy(roomstatus.FieldSTATUSDATA).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rsq *RoomStatusQuery) GroupBy(field string, fields ...string) *RoomStatusGroupBy {
	group := &RoomStatusGroupBy{config: rsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rsq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		STATUSDATA string `json:"STATUSDATA,omitempty"`
//	}
//
//	client.RoomStatus.Query().
//		Select(roomstatus.FieldSTATUSDATA).
//		Scan(ctx, &v)
//
func (rsq *RoomStatusQuery) Select(field string, fields ...string) *RoomStatusSelect {
	selector := &RoomStatusSelect{config: rsq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rsq.sqlQuery(), nil
	}
	return selector
}

func (rsq *RoomStatusQuery) prepareQuery(ctx context.Context) error {
	if rsq.path != nil {
		prev, err := rsq.path(ctx)
		if err != nil {
			return err
		}
		rsq.sql = prev
	}
	return nil
}

func (rsq *RoomStatusQuery) sqlAll(ctx context.Context) ([]*RoomStatus, error) {
	var (
		nodes       = []*RoomStatus{}
		_spec       = rsq.querySpec()
		loadedTypes = [1]bool{
			rsq.withSTATUSROOM != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &RoomStatus{config: rsq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, rsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rsq.withSTATUSROOM; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*RoomStatus)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Room(func(s *sql.Selector) {
			s.Where(sql.InValues(roomstatus.STATUSROOMColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.room_status_status_room
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "room_status_status_room" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "room_status_status_room" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.STATUSROOM = append(node.Edges.STATUSROOM, n)
		}
	}

	return nodes, nil
}

func (rsq *RoomStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rsq.querySpec()
	return sqlgraph.CountNodes(ctx, rsq.driver, _spec)
}

func (rsq *RoomStatusQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rsq *RoomStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roomstatus.Table,
			Columns: roomstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomstatus.FieldID,
			},
		},
		From:   rsq.sql,
		Unique: true,
	}
	if ps := rsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rsq *RoomStatusQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rsq.driver.Dialect())
	t1 := builder.Table(roomstatus.Table)
	selector := builder.Select(t1.Columns(roomstatus.Columns...)...).From(t1)
	if rsq.sql != nil {
		selector = rsq.sql
		selector.Select(selector.Columns(roomstatus.Columns...)...)
	}
	for _, p := range rsq.predicates {
		p(selector)
	}
	for _, p := range rsq.order {
		p(selector)
	}
	if offset := rsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoomStatusGroupBy is the builder for group-by RoomStatus entities.
type RoomStatusGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rsgb *RoomStatusGroupBy) Aggregate(fns ...AggregateFunc) *RoomStatusGroupBy {
	rsgb.fns = append(rsgb.fns, fns...)
	return rsgb
}

// Scan applies the group-by query and scan the result into the given value.
func (rsgb *RoomStatusGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rsgb.path(ctx)
	if err != nil {
		return err
	}
	rsgb.sql = query
	return rsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rsgb *RoomStatusGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RoomStatusGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RoomStatusGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rsgb *RoomStatusGroupBy) StringsX(ctx context.Context) []string {
	v, err := rsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RoomStatusGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomstatus.Label}
	default:
		err = fmt.Errorf("ent: RoomStatusGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rsgb *RoomStatusGroupBy) StringX(ctx context.Context) string {
	v, err := rsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RoomStatusGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RoomStatusGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rsgb *RoomStatusGroupBy) IntsX(ctx context.Context) []int {
	v, err := rsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RoomStatusGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomstatus.Label}
	default:
		err = fmt.Errorf("ent: RoomStatusGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rsgb *RoomStatusGroupBy) IntX(ctx context.Context) int {
	v, err := rsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RoomStatusGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RoomStatusGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rsgb *RoomStatusGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RoomStatusGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomstatus.Label}
	default:
		err = fmt.Errorf("ent: RoomStatusGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rsgb *RoomStatusGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RoomStatusGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RoomStatusGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rsgb *RoomStatusGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RoomStatusGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomstatus.Label}
	default:
		err = fmt.Errorf("ent: RoomStatusGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rsgb *RoomStatusGroupBy) BoolX(ctx context.Context) bool {
	v, err := rsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rsgb *RoomStatusGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rsgb.sqlQuery().Query()
	if err := rsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rsgb *RoomStatusGroupBy) sqlQuery() *sql.Selector {
	selector := rsgb.sql
	columns := make([]string, 0, len(rsgb.fields)+len(rsgb.fns))
	columns = append(columns, rsgb.fields...)
	for _, fn := range rsgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(rsgb.fields...)
}

// RoomStatusSelect is the builder for select fields of RoomStatus entities.
type RoomStatusSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (rss *RoomStatusSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := rss.path(ctx)
	if err != nil {
		return err
	}
	rss.sql = query
	return rss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rss *RoomStatusSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (rss *RoomStatusSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RoomStatusSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rss *RoomStatusSelect) StringsX(ctx context.Context) []string {
	v, err := rss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (rss *RoomStatusSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomstatus.Label}
	default:
		err = fmt.Errorf("ent: RoomStatusSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rss *RoomStatusSelect) StringX(ctx context.Context) string {
	v, err := rss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (rss *RoomStatusSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RoomStatusSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rss *RoomStatusSelect) IntsX(ctx context.Context) []int {
	v, err := rss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (rss *RoomStatusSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomstatus.Label}
	default:
		err = fmt.Errorf("ent: RoomStatusSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rss *RoomStatusSelect) IntX(ctx context.Context) int {
	v, err := rss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (rss *RoomStatusSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RoomStatusSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rss *RoomStatusSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (rss *RoomStatusSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomstatus.Label}
	default:
		err = fmt.Errorf("ent: RoomStatusSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rss *RoomStatusSelect) Float64X(ctx context.Context) float64 {
	v, err := rss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (rss *RoomStatusSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RoomStatusSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rss *RoomStatusSelect) BoolsX(ctx context.Context) []bool {
	v, err := rss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (rss *RoomStatusSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomstatus.Label}
	default:
		err = fmt.Errorf("ent: RoomStatusSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rss *RoomStatusSelect) BoolX(ctx context.Context) bool {
	v, err := rss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rss *RoomStatusSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rss.sqlQuery().Query()
	if err := rss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rss *RoomStatusSelect) sqlQuery() sql.Querier {
	selector := rss.sql
	selector.Select(selector.Columns(rss.fields...)...)
	return selector
}
