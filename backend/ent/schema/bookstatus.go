package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent"
)

// Bookstatus holds the schema definition for the Bookstatus entity.
type Bookstatus struct {
	ent.Schema
}

// Fields of the Book.
func (Bookstatus) Fields() []ent.Field {
	return []ent.Field{
			field.String("BOOKSTATUSDATA").NotEmpty(),
	}
}

// Edges of the Bookstatus.
func (Bookstatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("BOOKSTATUS_BOOK",Book.Type),
    }
}
