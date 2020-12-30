// Code generated by entc, DO NOT EDIT.

package place

import (
	"github.com/B6001186/Contagions/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PlaceName applies equality check predicate on the "PlaceName" field. It's identical to PlaceNameEQ.
func PlaceName(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlaceName), v))
	})
}

// PlaceNameEQ applies the EQ predicate on the "PlaceName" field.
func PlaceNameEQ(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlaceName), v))
	})
}

// PlaceNameNEQ applies the NEQ predicate on the "PlaceName" field.
func PlaceNameNEQ(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPlaceName), v))
	})
}

// PlaceNameIn applies the In predicate on the "PlaceName" field.
func PlaceNameIn(vs ...string) predicate.Place {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Place(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPlaceName), v...))
	})
}

// PlaceNameNotIn applies the NotIn predicate on the "PlaceName" field.
func PlaceNameNotIn(vs ...string) predicate.Place {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Place(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPlaceName), v...))
	})
}

// PlaceNameGT applies the GT predicate on the "PlaceName" field.
func PlaceNameGT(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPlaceName), v))
	})
}

// PlaceNameGTE applies the GTE predicate on the "PlaceName" field.
func PlaceNameGTE(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPlaceName), v))
	})
}

// PlaceNameLT applies the LT predicate on the "PlaceName" field.
func PlaceNameLT(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPlaceName), v))
	})
}

// PlaceNameLTE applies the LTE predicate on the "PlaceName" field.
func PlaceNameLTE(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPlaceName), v))
	})
}

// PlaceNameContains applies the Contains predicate on the "PlaceName" field.
func PlaceNameContains(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPlaceName), v))
	})
}

// PlaceNameHasPrefix applies the HasPrefix predicate on the "PlaceName" field.
func PlaceNameHasPrefix(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPlaceName), v))
	})
}

// PlaceNameHasSuffix applies the HasSuffix predicate on the "PlaceName" field.
func PlaceNameHasSuffix(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPlaceName), v))
	})
}

// PlaceNameEqualFold applies the EqualFold predicate on the "PlaceName" field.
func PlaceNameEqualFold(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPlaceName), v))
	})
}

// PlaceNameContainsFold applies the ContainsFold predicate on the "PlaceName" field.
func PlaceNameContainsFold(v string) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPlaceName), v))
	})
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Place) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Place) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
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
func Not(p predicate.Place) predicate.Place {
	return predicate.Place(func(s *sql.Selector) {
		p(s.Not())
	})
}