// Code generated by entc, DO NOT EDIT.

package bookstatus

const (
	// Label holds the string label denoting the bookstatus type in the database.
	Label = "bookstatus"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBOOKSTATUSDATA holds the string denoting the bookstatusdata field in the database.
	FieldBOOKSTATUSDATA = "bookstatusdata"

	// EdgeBOOKSTATUSBOOK holds the string denoting the bookstatus_book edge name in mutations.
	EdgeBOOKSTATUSBOOK = "BOOKSTATUS_BOOK"

	// Table holds the table name of the bookstatus in the database.
	Table = "bookstatuses"
	// BOOKSTATUSBOOKTable is the table the holds the BOOKSTATUS_BOOK relation/edge.
	BOOKSTATUSBOOKTable = "books"
	// BOOKSTATUSBOOKInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	BOOKSTATUSBOOKInverseTable = "books"
	// BOOKSTATUSBOOKColumn is the table column denoting the BOOKSTATUS_BOOK relation/edge.
	BOOKSTATUSBOOKColumn = "bookstatus_bookstatus_book"
)

// Columns holds all SQL columns for bookstatus fields.
var Columns = []string{
	FieldID,
	FieldBOOKSTATUSDATA,
}

var (
	// BOOKSTATUSDATAValidator is a validator for the "BOOKSTATUSDATA" field. It is called by the builders before save.
	BOOKSTATUSDATAValidator func(string) error
)
