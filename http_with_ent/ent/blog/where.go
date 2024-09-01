// Code generated by ent, DO NOT EDIT.

package blog

import (
	"example/hello/http_with_ent/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldDescription, v))
}

// Episode applies equality check predicate on the "episode" field. It's identical to EpisodeEQ.
func Episode(v int) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldEpisode, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldCreatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContainsFold(FieldDescription, v))
}

// EpisodeEQ applies the EQ predicate on the "episode" field.
func EpisodeEQ(v int) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldEpisode, v))
}

// EpisodeNEQ applies the NEQ predicate on the "episode" field.
func EpisodeNEQ(v int) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldEpisode, v))
}

// EpisodeIn applies the In predicate on the "episode" field.
func EpisodeIn(vs ...int) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldEpisode, vs...))
}

// EpisodeNotIn applies the NotIn predicate on the "episode" field.
func EpisodeNotIn(vs ...int) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldEpisode, vs...))
}

// EpisodeGT applies the GT predicate on the "episode" field.
func EpisodeGT(v int) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldEpisode, v))
}

// EpisodeGTE applies the GTE predicate on the "episode" field.
func EpisodeGTE(v int) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldEpisode, v))
}

// EpisodeLT applies the LT predicate on the "episode" field.
func EpisodeLT(v int) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldEpisode, v))
}

// EpisodeLTE applies the LTE predicate on the "episode" field.
func EpisodeLTE(v int) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldEpisode, v))
}

// EpisodeIsNil applies the IsNil predicate on the "episode" field.
func EpisodeIsNil() predicate.Blog {
	return predicate.Blog(sql.FieldIsNull(FieldEpisode))
}

// EpisodeNotNil applies the NotNil predicate on the "episode" field.
func EpisodeNotNil() predicate.Blog {
	return predicate.Blog(sql.FieldNotNull(FieldEpisode))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := newTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Blog) predicate.Blog {
	return predicate.Blog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Blog) predicate.Blog {
	return predicate.Blog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Blog) predicate.Blog {
	return predicate.Blog(sql.NotPredicates(p))
}