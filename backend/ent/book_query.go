// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/m16_z/app/ent/book"
	"github.com/m16_z/app/ent/bookstatus"
	"github.com/m16_z/app/ent/predicate"
	"github.com/m16_z/app/ent/room"
	"github.com/m16_z/app/ent/user"
)

// BookQuery is the builder for querying Book entities.
type BookQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Book
	// eager-loading edges.
	withBOOKUSER       *UserQuery
	withBOOKROOM       *RoomQuery
	withBOOKBOOKSTATUS *BookstatusQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (bq *BookQuery) Where(ps ...predicate.Book) *BookQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BookQuery) Limit(limit int) *BookQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BookQuery) Offset(offset int) *BookQuery {
	bq.offset = &offset
	return bq
}

// Order adds an order step to the query.
func (bq *BookQuery) Order(o ...OrderFunc) *BookQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryBOOKUSER chains the current query on the BOOK_USER edge.
func (bq *BookQuery) QueryBOOKUSER() *UserQuery {
	query := &UserQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, bq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, book.BOOKUSERTable, book.BOOKUSERColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBOOKROOM chains the current query on the BOOK_ROOM edge.
func (bq *BookQuery) QueryBOOKROOM() *RoomQuery {
	query := &RoomQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, bq.sqlQuery()),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, book.BOOKROOMTable, book.BOOKROOMColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBOOKBOOKSTATUS chains the current query on the BOOK_BOOKSTATUS edge.
func (bq *BookQuery) QueryBOOKBOOKSTATUS() *BookstatusQuery {
	query := &BookstatusQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, bq.sqlQuery()),
			sqlgraph.To(bookstatus.Table, bookstatus.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, book.BOOKBOOKSTATUSTable, book.BOOKBOOKSTATUSColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Book entity in the query. Returns *NotFoundError when no book was found.
func (bq *BookQuery) First(ctx context.Context) (*Book, error) {
	bs, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(bs) == 0 {
		return nil, &NotFoundError{book.Label}
	}
	return bs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BookQuery) FirstX(ctx context.Context) *Book {
	b, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return b
}

// FirstID returns the first Book id in the query. Returns *NotFoundError when no id was found.
func (bq *BookQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{book.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (bq *BookQuery) FirstXID(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Book entity in the query, returns an error if not exactly one entity was returned.
func (bq *BookQuery) Only(ctx context.Context) (*Book, error) {
	bs, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(bs) {
	case 1:
		return bs[0], nil
	case 0:
		return nil, &NotFoundError{book.Label}
	default:
		return nil, &NotSingularError{book.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BookQuery) OnlyX(ctx context.Context) *Book {
	b, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// OnlyID returns the only Book id in the query, returns an error if not exactly one id was returned.
func (bq *BookQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{book.Label}
	default:
		err = &NotSingularError{book.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BookQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Books.
func (bq *BookQuery) All(ctx context.Context) ([]*Book, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bq *BookQuery) AllX(ctx context.Context) []*Book {
	bs, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return bs
}

// IDs executes the query and returns a list of Book ids.
func (bq *BookQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bq.Select(book.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BookQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BookQuery) Count(ctx context.Context) (int, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BookQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BookQuery) Exist(ctx context.Context) (bool, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BookQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BookQuery) Clone() *BookQuery {
	return &BookQuery{
		config:     bq.config,
		limit:      bq.limit,
		offset:     bq.offset,
		order:      append([]OrderFunc{}, bq.order...),
		unique:     append([]string{}, bq.unique...),
		predicates: append([]predicate.Book{}, bq.predicates...),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

//  WithBOOKUSER tells the query-builder to eager-loads the nodes that are connected to
// the "BOOK_USER" edge. The optional arguments used to configure the query builder of the edge.
func (bq *BookQuery) WithBOOKUSER(opts ...func(*UserQuery)) *BookQuery {
	query := &UserQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withBOOKUSER = query
	return bq
}

//  WithBOOKROOM tells the query-builder to eager-loads the nodes that are connected to
// the "BOOK_ROOM" edge. The optional arguments used to configure the query builder of the edge.
func (bq *BookQuery) WithBOOKROOM(opts ...func(*RoomQuery)) *BookQuery {
	query := &RoomQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withBOOKROOM = query
	return bq
}

//  WithBOOKBOOKSTATUS tells the query-builder to eager-loads the nodes that are connected to
// the "BOOK_BOOKSTATUS" edge. The optional arguments used to configure the query builder of the edge.
func (bq *BookQuery) WithBOOKBOOKSTATUS(opts ...func(*BookstatusQuery)) *BookQuery {
	query := &BookstatusQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withBOOKBOOKSTATUS = query
	return bq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RESERVATIONS time.Time `json:"RESERVATIONS,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Book.Query().
//		GroupBy(book.FieldRESERVATIONS).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bq *BookQuery) GroupBy(field string, fields ...string) *BookGroupBy {
	group := &BookGroupBy{config: bq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		RESERVATIONS time.Time `json:"RESERVATIONS,omitempty"`
//	}
//
//	client.Book.Query().
//		Select(book.FieldRESERVATIONS).
//		Scan(ctx, &v)
//
func (bq *BookQuery) Select(field string, fields ...string) *BookSelect {
	selector := &BookSelect{config: bq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(), nil
	}
	return selector
}

func (bq *BookQuery) prepareQuery(ctx context.Context) error {
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BookQuery) sqlAll(ctx context.Context) ([]*Book, error) {
	var (
		nodes       = []*Book{}
		withFKs     = bq.withFKs
		_spec       = bq.querySpec()
		loadedTypes = [3]bool{
			bq.withBOOKUSER != nil,
			bq.withBOOKROOM != nil,
			bq.withBOOKBOOKSTATUS != nil,
		}
	)
	if bq.withBOOKUSER != nil || bq.withBOOKROOM != nil || bq.withBOOKBOOKSTATUS != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, book.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Book{config: bq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
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
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bq.withBOOKUSER; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Book)
		for i := range nodes {
			if fk := nodes[i].user_user_book; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_user_book" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.BOOKUSER = n
			}
		}
	}

	if query := bq.withBOOKROOM; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Book)
		for i := range nodes {
			if fk := nodes[i].room_room_book; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(room.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "room_room_book" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.BOOKROOM = n
			}
		}
	}

	if query := bq.withBOOKBOOKSTATUS; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Book)
		for i := range nodes {
			if fk := nodes[i].bookstatus_bookstatus_book; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(bookstatus.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "bookstatus_bookstatus_book" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.BOOKBOOKSTATUS = n
			}
		}
	}

	return nodes, nil
}

func (bq *BookQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BookQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (bq *BookQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   book.Table,
			Columns: book.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: book.FieldID,
			},
		},
		From:   bq.sql,
		Unique: true,
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BookQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(book.Table)
	selector := builder.Select(t1.Columns(book.Columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(book.Columns...)...)
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BookGroupBy is the builder for group-by Book entities.
type BookGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BookGroupBy) Aggregate(fns ...AggregateFunc) *BookGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scan the result into the given value.
func (bgb *BookGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bgb.path(ctx)
	if err != nil {
		return err
	}
	bgb.sql = query
	return bgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bgb *BookGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BookGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bgb *BookGroupBy) StringsX(ctx context.Context) []string {
	v, err := bgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{book.Label}
	default:
		err = fmt.Errorf("ent: BookGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bgb *BookGroupBy) StringX(ctx context.Context) string {
	v, err := bgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BookGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bgb *BookGroupBy) IntsX(ctx context.Context) []int {
	v, err := bgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{book.Label}
	default:
		err = fmt.Errorf("ent: BookGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bgb *BookGroupBy) IntX(ctx context.Context) int {
	v, err := bgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BookGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bgb *BookGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{book.Label}
	default:
		err = fmt.Errorf("ent: BookGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bgb *BookGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BookGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bgb *BookGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{book.Label}
	default:
		err = fmt.Errorf("ent: BookGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bgb *BookGroupBy) BoolX(ctx context.Context) bool {
	v, err := bgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bgb *BookGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bgb.sqlQuery().Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BookGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql
	columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
	columns = append(columns, bgb.fields...)
	for _, fn := range bgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(bgb.fields...)
}

// BookSelect is the builder for select fields of Book entities.
type BookSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (bs *BookSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := bs.path(ctx)
	if err != nil {
		return err
	}
	bs.sql = query
	return bs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bs *BookSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (bs *BookSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BookSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bs *BookSelect) StringsX(ctx context.Context) []string {
	v, err := bs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (bs *BookSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{book.Label}
	default:
		err = fmt.Errorf("ent: BookSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bs *BookSelect) StringX(ctx context.Context) string {
	v, err := bs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (bs *BookSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BookSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bs *BookSelect) IntsX(ctx context.Context) []int {
	v, err := bs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (bs *BookSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{book.Label}
	default:
		err = fmt.Errorf("ent: BookSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bs *BookSelect) IntX(ctx context.Context) int {
	v, err := bs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (bs *BookSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BookSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bs *BookSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (bs *BookSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{book.Label}
	default:
		err = fmt.Errorf("ent: BookSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bs *BookSelect) Float64X(ctx context.Context) float64 {
	v, err := bs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (bs *BookSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BookSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bs *BookSelect) BoolsX(ctx context.Context) []bool {
	v, err := bs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (bs *BookSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{book.Label}
	default:
		err = fmt.Errorf("ent: BookSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bs *BookSelect) BoolX(ctx context.Context) bool {
	v, err := bs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bs *BookSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bs.sqlQuery().Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bs *BookSelect) sqlQuery() sql.Querier {
	selector := bs.sql
	selector.Select(selector.Columns(bs.fields...)...)
	return selector
}
