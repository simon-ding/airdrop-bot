package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StepRun holds the schema definition for the StepRun entity.
type StepRun struct {
	ent.Schema
}

// Fields of the StepRun.
func (StepRun) Fields() []ent.Field {
	return []ent.Field{
		field.String("step"),
		field.Enum("status").Values("pending", "running", "success", "failed"),
		field.String("reason"),
	}
}

// Edges of the StepRun.
func (StepRun) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("step_runs").Unique(),
		edge.From("node", Node.Type).
			Ref("step_runs").Unique(),
	}
}
