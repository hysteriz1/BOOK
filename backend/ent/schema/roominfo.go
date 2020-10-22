package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent"
)


// RoomInfo holds the schema definition for the RoomInfo entity.
type RoomInfo struct {
	ent.Schema
}

// Fields of the RoomInfo.
func (RoomInfo) Fields() []ent.Field {
	return []ent.Field{		
		field.Int("INFOBED"),
		field.Int("INFOREFRIGERAT"), 
		field.Int("INFOWARDROB"),
		field.Int("INFOOFFICEDE"), 
	}
}

// Edges of the Info.
func (RoomInfo) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("INFO_ROOM",Room.Type),
	}
}
