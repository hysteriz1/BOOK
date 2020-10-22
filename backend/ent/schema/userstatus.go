package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// UserStatus holds the schema definition for the UserStatus entity.
type UserStatus struct {
	ent.Schema
}

// Fields of the UserStatus.
func (UserStatus) Fields() []ent.Field {
		return []ent.Field{			
			field.String("STATUS").NotEmpty(),
		}
}

// Edges of the UserStatus.
func (UserStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("USERSTATUS_USER",User.Type),
   }
}
