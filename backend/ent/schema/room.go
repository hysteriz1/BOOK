package schema
import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{		 
		field.String("ROOMNAME").NotEmpty(), 	
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ROOM_BOOK",Book.Type),
		edge.From("ROOM_ROOMTYPE",RoomType.Type).
            Ref("ROOMTYPE_ROOM").
			Unique(),
		edge.From("ROOM_STATUS",RoomStatus.Type).
            Ref("STATUS_ROOM").
			Unique(),
		edge.From("ROOM_INFO",RoomInfo.Type).
            Ref("INFO_ROOM").
			Unique(),
    }
}
