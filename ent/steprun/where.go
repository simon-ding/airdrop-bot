// Code generated by ent, DO NOT EDIT.

package steprun

import (
	"airdrop-bot/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.StepRun {
	return predicate.StepRun(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.StepRun {
	return predicate.StepRun(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.StepRun {
	return predicate.StepRun(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.StepRun {
	return predicate.StepRun(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.StepRun {
	return predicate.StepRun(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.StepRun {
	return predicate.StepRun(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.StepRun {
	return predicate.StepRun(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.StepRun {
	return predicate.StepRun(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.StepRun {
	return predicate.StepRun(sql.FieldLTE(FieldID, id))
}

// Step applies equality check predicate on the "step" field. It's identical to StepEQ.
func Step(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldEQ(FieldStep, v))
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldEQ(FieldReason, v))
}

// StepEQ applies the EQ predicate on the "step" field.
func StepEQ(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldEQ(FieldStep, v))
}

// StepNEQ applies the NEQ predicate on the "step" field.
func StepNEQ(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldNEQ(FieldStep, v))
}

// StepIn applies the In predicate on the "step" field.
func StepIn(vs ...string) predicate.StepRun {
	return predicate.StepRun(sql.FieldIn(FieldStep, vs...))
}

// StepNotIn applies the NotIn predicate on the "step" field.
func StepNotIn(vs ...string) predicate.StepRun {
	return predicate.StepRun(sql.FieldNotIn(FieldStep, vs...))
}

// StepGT applies the GT predicate on the "step" field.
func StepGT(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldGT(FieldStep, v))
}

// StepGTE applies the GTE predicate on the "step" field.
func StepGTE(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldGTE(FieldStep, v))
}

// StepLT applies the LT predicate on the "step" field.
func StepLT(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldLT(FieldStep, v))
}

// StepLTE applies the LTE predicate on the "step" field.
func StepLTE(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldLTE(FieldStep, v))
}

// StepContains applies the Contains predicate on the "step" field.
func StepContains(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldContains(FieldStep, v))
}

// StepHasPrefix applies the HasPrefix predicate on the "step" field.
func StepHasPrefix(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldHasPrefix(FieldStep, v))
}

// StepHasSuffix applies the HasSuffix predicate on the "step" field.
func StepHasSuffix(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldHasSuffix(FieldStep, v))
}

// StepEqualFold applies the EqualFold predicate on the "step" field.
func StepEqualFold(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldEqualFold(FieldStep, v))
}

// StepContainsFold applies the ContainsFold predicate on the "step" field.
func StepContainsFold(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldContainsFold(FieldStep, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.StepRun {
	return predicate.StepRun(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.StepRun {
	return predicate.StepRun(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.StepRun {
	return predicate.StepRun(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.StepRun {
	return predicate.StepRun(sql.FieldNotIn(FieldStatus, vs...))
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldEQ(FieldReason, v))
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldNEQ(FieldReason, v))
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...string) predicate.StepRun {
	return predicate.StepRun(sql.FieldIn(FieldReason, vs...))
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...string) predicate.StepRun {
	return predicate.StepRun(sql.FieldNotIn(FieldReason, vs...))
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldGT(FieldReason, v))
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldGTE(FieldReason, v))
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldLT(FieldReason, v))
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldLTE(FieldReason, v))
}

// ReasonContains applies the Contains predicate on the "reason" field.
func ReasonContains(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldContains(FieldReason, v))
}

// ReasonHasPrefix applies the HasPrefix predicate on the "reason" field.
func ReasonHasPrefix(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldHasPrefix(FieldReason, v))
}

// ReasonHasSuffix applies the HasSuffix predicate on the "reason" field.
func ReasonHasSuffix(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldHasSuffix(FieldReason, v))
}

// ReasonEqualFold applies the EqualFold predicate on the "reason" field.
func ReasonEqualFold(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldEqualFold(FieldReason, v))
}

// ReasonContainsFold applies the ContainsFold predicate on the "reason" field.
func ReasonContainsFold(v string) predicate.StepRun {
	return predicate.StepRun(sql.FieldContainsFold(FieldReason, v))
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := newAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNode applies the HasEdge predicate on the "node" edge.
func HasNode() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodeWith applies the HasEdge predicate on the "node" edge with a given conditions (other predicates).
func HasNodeWith(preds ...predicate.Node) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := newNodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StepRun) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StepRun) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
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
func Not(p predicate.StepRun) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		p(s.Not())
	})
}
