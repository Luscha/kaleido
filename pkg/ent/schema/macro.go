package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Macro struct {
	ent.Schema
}

func (Macro) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "macro"},
	}
}

// Fields of the Macro.
func (Macro) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Immutable(),
		field.String("name").
			Immutable(),
		field.String("description"),
		field.String("metadata"),
		field.String("code"),
	}
}

// Edges of the Macro.
func (Macro) Edges() []ent.Edge {
	return []ent.Edge{}
}
