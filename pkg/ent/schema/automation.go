package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Automation struct {
	ent.Schema
}

func (Automation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "automation"},
	}
}

// Fields of the Automation.
func (Automation) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Immutable(),
		field.String("name").
			Immutable(),
		field.String("description"),
		field.String("trigger"),
		field.String("manifest"),
	}
}

// Edges of the Automation.
func (Automation) Edges() []ent.Edge {
	return []ent.Edge{}
}
