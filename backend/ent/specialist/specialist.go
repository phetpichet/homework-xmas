// Code generated by entc, DO NOT EDIT.

package specialist

const (
	// Label holds the string label denoting the specialist type in the database.
	Label = "specialist"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSname holds the string denoting the sname field in the database.
	FieldSname = "sname"

	// Table holds the table name of the specialist in the database.
	Table = "specialists"
)

// Columns holds all SQL columns for specialist fields.
var Columns = []string{
	FieldID,
	FieldSname,
}

var (
	// SnameValidator is a validator for the "sname" field. It is called by the builders before save.
	SnameValidator func(string) error
)
