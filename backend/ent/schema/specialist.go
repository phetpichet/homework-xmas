package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
 )

// Specialist holds the schema definition for the Specialist entity.
type Specialist struct {
	ent.Schema
}

// Fields of the Specialist.
func (Specialist) Fields() []ent.Field {
	return []ent.Field {
		field.String("sname").NotEmpty(),
	}
}

// Edges of the Specialist.
func (Specialist) Edges() []ent.Edge {
	return nil
}
