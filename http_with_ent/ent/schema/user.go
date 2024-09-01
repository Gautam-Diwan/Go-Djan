package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	re := regexp.MustCompile(`^[a-zA-Z0-9_ !@#$%^&*()-+=\[\]{};:'",.<>?/\\|~]*$`)
	return []ent.Field{
		field.String("name").
			MinLen(3).
			Match(re).
			Comment("Name of the author/user"),
		field.String("password").
			Default("QWERTYUIO").
			Sensitive().
			Comment("Password of the author/user"),
		field.Int("age").
			Positive().
			Default(1).
			Optional().
			Comment("Age of the author/user"),
		field.Bool("is_active").
			Default(true).
			Comment("Activity of the author/user"),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Comment("Time when the user was created"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M relation to blogs
		edge.To("blogs", Blog.Type),
		// Loop Back M2M Relation which can help in adding friends
		edge.To("friends", User.Type),
	}
}
