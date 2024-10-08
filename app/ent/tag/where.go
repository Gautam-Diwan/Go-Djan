// Code generated by ent, DO NOT EDIT.

package tag

import (
	"go/djan/app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Tag {
	return predicate.Tag(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Tag {
	return predicate.Tag(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Tag {
	return predicate.Tag(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Tag {
	return predicate.Tag(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Tag {
	return predicate.Tag(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Tag {
	return predicate.Tag(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldName, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldType, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Tag {
	return predicate.Tag(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Tag {
	return predicate.Tag(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Tag {
	return predicate.Tag(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Tag {
	return predicate.Tag(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Tag {
	return predicate.Tag(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Tag {
	return predicate.Tag(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Tag {
	return predicate.Tag(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Tag {
	return predicate.Tag(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Tag {
	return predicate.Tag(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Tag {
	return predicate.Tag(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Tag {
	return predicate.Tag(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Tag {
	return predicate.Tag(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Tag {
	return predicate.Tag(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Tag {
	return predicate.Tag(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Tag {
	return predicate.Tag(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Tag {
	return predicate.Tag(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Tag {
	return predicate.Tag(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Tag {
	return predicate.Tag(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Tag {
	return predicate.Tag(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Tag {
	return predicate.Tag(sql.FieldContainsFold(FieldType, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v Category) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v Category) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...Category) predicate.Tag {
	return predicate.Tag(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...Category) predicate.Tag {
	return predicate.Tag(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryIsNil applies the IsNil predicate on the "category" field.
func CategoryIsNil() predicate.Tag {
	return predicate.Tag(sql.FieldIsNull(FieldCategory))
}

// CategoryNotNil applies the NotNil predicate on the "category" field.
func CategoryNotNil() predicate.Tag {
	return predicate.Tag(sql.FieldNotNull(FieldCategory))
}

// HasBlogs applies the HasEdge predicate on the "blogs" edge.
func HasBlogs() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, BlogsTable, BlogsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlogsWith applies the HasEdge predicate on the "blogs" edge with a given conditions (other predicates).
func HasBlogsWith(preds ...predicate.Blog) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		step := newBlogsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tag) predicate.Tag {
	return predicate.Tag(sql.NotPredicates(p))
}
