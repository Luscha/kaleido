// Code generated by ent, DO NOT EDIT.

package procedure

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the procedure type in the database.
	Label = "procedure"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldManifest holds the string denoting the manifest field in the database.
	FieldManifest = "manifest"
	// Table holds the table name of the procedure in the database.
	Table = "procedure"
)

// Columns holds all SQL columns for procedure fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldMetadata,
	FieldManifest,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Procedure queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByMetadata orders the results by the metadata field.
func ByMetadata(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMetadata, opts...).ToFunc()
}

// ByManifest orders the results by the manifest field.
func ByManifest(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldManifest, opts...).ToFunc()
}
