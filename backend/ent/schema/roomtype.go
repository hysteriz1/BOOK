package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent"
)

// RoomType holds the schema definition for the RoomType entity.
type RoomType struct {
	ent.Schema
}

// Fields of the Type.
func (RoomType) Fields() []ent.Field {
	return []ent.Field{		
		field.String("ROOMTYPEDATA").NotEmpty(),
		field.Int("COSTDATA"),
	}
}

// Edges of the Type.
func (RoomType) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("ROOMTYPE_ROOM",Room.Type),
	}
}
