package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{		 
		field.Time("RESERVATIONS"), 
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("BOOK_USER",User.Type).
            Ref("USER_BOOK").
			Unique(),
		edge.From("BOOK_ROOM",Room.Type).
            Ref("ROOM_BOOK").
			Unique(),
		edge.From("BOOK_BOOKSTATUS",Bookstatus.Type).
            Ref("BOOKSTATUS_BOOK").
            Unique(),
    }
}
