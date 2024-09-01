package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MinLen(3).
			MaxLen(30).
			Comment("Title of the Blog"),
		field.String("description").
			MinLen(3).
			Comment("Description of the Blog"),
		field.Int("episode").
			Positive().
			Optional(),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Comment("Time when the Blog was created"),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return []ent.Edge{
		// Back referencing O2M from User
		edge.From("user", User.Type).Ref("blogs").Unique(),
		// Back referencing M2M from Tag
		edge.From("tags", Tag.Type).Ref("blogs"),
	}
}
