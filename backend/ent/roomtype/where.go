// Code generated by entc, DO NOT EDIT.

package roomtype

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/m16_z/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ROOMTYPEDATA applies equality check predicate on the "ROOMTYPEDATA" field. It's identical to ROOMTYPEDATAEQ.
func ROOMTYPEDATA(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldROOMTYPEDATA), v))
	})
}

// COSTDATA applies equality check predicate on the "COSTDATA" field. It's identical to COSTDATAEQ.
func COSTDATA(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCOSTDATA), v))
	})
}

// ROOMTYPEDATAEQ applies the EQ predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATAEQ(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldROOMTYPEDATA), v))
	})
}

// ROOMTYPEDATANEQ applies the NEQ predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATANEQ(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldROOMTYPEDATA), v))
	})
}

// ROOMTYPEDATAIn applies the In predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATAIn(vs ...string) predicate.RoomType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RoomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldROOMTYPEDATA), v...))
	})
}

// ROOMTYPEDATANotIn applies the NotIn predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATANotIn(vs ...string) predicate.RoomType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RoomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldROOMTYPEDATA), v...))
	})
}

// ROOMTYPEDATAGT applies the GT predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATAGT(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldROOMTYPEDATA), v))
	})
}

// ROOMTYPEDATAGTE applies the GTE predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATAGTE(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldROOMTYPEDATA), v))
	})
}

// ROOMTYPEDATALT applies the LT predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATALT(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldROOMTYPEDATA), v))
	})
}

// ROOMTYPEDATALTE applies the LTE predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATALTE(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldROOMTYPEDATA), v))
	})
}

// ROOMTYPEDATAContains applies the Contains predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATAContains(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldROOMTYPEDATA), v))
	})
}

// ROOMTYPEDATAHasPrefix applies the HasPrefix predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATAHasPrefix(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldROOMTYPEDATA), v))
	})
}

// ROOMTYPEDATAHasSuffix applies the HasSuffix predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATAHasSuffix(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldROOMTYPEDATA), v))
	})
}

// ROOMTYPEDATAEqualFold applies the EqualFold predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATAEqualFold(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldROOMTYPEDATA), v))
	})
}

// ROOMTYPEDATAContainsFold applies the ContainsFold predicate on the "ROOMTYPEDATA" field.
func ROOMTYPEDATAContainsFold(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldROOMTYPEDATA), v))
	})
}

// COSTDATAEQ applies the EQ predicate on the "COSTDATA" field.
func COSTDATAEQ(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCOSTDATA), v))
	})
}

// COSTDATANEQ applies the NEQ predicate on the "COSTDATA" field.
func COSTDATANEQ(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCOSTDATA), v))
	})
}

// COSTDATAIn applies the In predicate on the "COSTDATA" field.
func COSTDATAIn(vs ...int) predicate.RoomType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RoomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCOSTDATA), v...))
	})
}

// COSTDATANotIn applies the NotIn predicate on the "COSTDATA" field.
func COSTDATANotIn(vs ...int) predicate.RoomType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RoomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCOSTDATA), v...))
	})
}

// COSTDATAGT applies the GT predicate on the "COSTDATA" field.
func COSTDATAGT(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCOSTDATA), v))
	})
}

// COSTDATAGTE applies the GTE predicate on the "COSTDATA" field.
func COSTDATAGTE(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCOSTDATA), v))
	})
}

// COSTDATALT applies the LT predicate on the "COSTDATA" field.
func COSTDATALT(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCOSTDATA), v))
	})
}

// COSTDATALTE applies the LTE predicate on the "COSTDATA" field.
func COSTDATALTE(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCOSTDATA), v))
	})
}

// HasROOMTYPEROOM applies the HasEdge predicate on the "ROOMTYPE_ROOM" edge.
func HasROOMTYPEROOM() predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ROOMTYPEROOMTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ROOMTYPEROOMTable, ROOMTYPEROOMColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasROOMTYPEROOMWith applies the HasEdge predicate on the "ROOMTYPE_ROOM" edge with a given conditions (other predicates).
func HasROOMTYPEROOMWith(preds ...predicate.Room) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ROOMTYPEROOMInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ROOMTYPEROOMTable, ROOMTYPEROOMColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.RoomType) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.RoomType) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RoomType) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		p(s.Not())
	})
}
