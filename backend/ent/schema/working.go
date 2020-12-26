package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
 )
// Working holds the schema definition for the Working entity.
type Working struct {
	ent.Schema
}

// Fields of the Working.
func (Working) Fields() []ent.Field {
	return []ent.Field {
		field.Time("added_time"),
	}
}

// Edges of the Working.
func (Working) Edges() []ent.Edge {
	return nil
}
