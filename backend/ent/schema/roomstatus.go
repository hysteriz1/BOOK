package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent"
)

// RoomStatus holds the schema definition for the RoomStatus entity.
type RoomStatus struct {
	ent.Schema
}

// Fields of the Status.
func (RoomStatus) Fields() []ent.Field {
	return []ent.Field{		
		field.String("STATUSDATA").NotEmpty(),
	}
}

// Edges of the Status.
func (RoomStatus) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("STATUS_ROOM",Room.Type),
	}
}
