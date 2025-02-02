// Code generated by entc, DO NOT EDIT.

package book

const (
	// Label holds the string label denoting the book type in the database.
	Label = "book"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRESERVATIONS holds the string denoting the reservations field in the database.
	FieldRESERVATIONS = "reservations"

	// EdgeBOOKUSER holds the string denoting the book_user edge name in mutations.
	EdgeBOOKUSER = "BOOK_USER"
	// EdgeBOOKROOM holds the string denoting the book_room edge name in mutations.
	EdgeBOOKROOM = "BOOK_ROOM"
	// EdgeBOOKBOOKSTATUS holds the string denoting the book_bookstatus edge name in mutations.
	EdgeBOOKBOOKSTATUS = "BOOK_BOOKSTATUS"

	// Table holds the table name of the book in the database.
	Table = "books"
	// BOOKUSERTable is the table the holds the BOOK_USER relation/edge.
	BOOKUSERTable = "books"
	// BOOKUSERInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	BOOKUSERInverseTable = "users"
	// BOOKUSERColumn is the table column denoting the BOOK_USER relation/edge.
	BOOKUSERColumn = "user_user_book"
	// BOOKROOMTable is the table the holds the BOOK_ROOM relation/edge.
	BOOKROOMTable = "books"
	// BOOKROOMInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	BOOKROOMInverseTable = "rooms"
	// BOOKROOMColumn is the table column denoting the BOOK_ROOM relation/edge.
	BOOKROOMColumn = "room_room_book"
	// BOOKBOOKSTATUSTable is the table the holds the BOOK_BOOKSTATUS relation/edge.
	BOOKBOOKSTATUSTable = "books"
	// BOOKBOOKSTATUSInverseTable is the table name for the Bookstatus entity.
	// It exists in this package in order to avoid circular dependency with the "bookstatus" package.
	BOOKBOOKSTATUSInverseTable = "bookstatuses"
	// BOOKBOOKSTATUSColumn is the table column denoting the BOOK_BOOKSTATUS relation/edge.
	BOOKBOOKSTATUSColumn = "bookstatus_bookstatus_book"
)

// Columns holds all SQL columns for book fields.
var Columns = []string{
	FieldID,
	FieldRESERVATIONS,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Book type.
var ForeignKeys = []string{
	"bookstatus_bookstatus_book",
	"room_room_book",
	"user_user_book",
}
