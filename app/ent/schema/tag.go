package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MinLen(1).
			Match(regexp.MustCompile("[a-zA-Z_ -]+$")).
			Unique().
			Comment("Name of the Tag"),
		field.String("type").
			MinLen(1).
			MaxLen(10).
			Match(regexp.MustCompile("[a-zA-Z]+$")).
			Default("Common").
			Comment("Type of Blog"),
		field.Enum("category").
			Values("Hot", "Trending", "Newest", "Controversial").
			Optional(),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		// Simple M2M
		edge.To("blogs", Blog.Type),
	}
}
