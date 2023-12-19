package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Procedure struct {
	ent.Schema
}

func (Procedure) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "procedure"},
	}
}

// Fields of the Procedure.
func (Procedure) Fields() []ent.Field {
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

// Edges of the Procedure.
func (Procedure) Edges() []ent.Edge {
	return []ent.Edge{}
}
