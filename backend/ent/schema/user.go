package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
 
// User holds the schema definition for the User entity.
type User struct {
   ent.Schema
}
 
// Fields of the User.
func (User) Fields() []ent.Field {
   return []ent.Field{       
       field.String("USEREMAIL").NotEmpty(),
       field.String("USERNAME").NotEmpty(),
   }
}
 
// Edges of the User.
func (User) Edges() []ent.Edge {
   return []ent.Edge{
      edge.To("USER_BOOK",Book.Type),
      edge.From("USER_USERSTATUS",UserStatus.Type).
            Ref("USERSTATUS_USER").
			Unique(),
 }
}
