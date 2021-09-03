// Code generated by entc, DO NOT EDIT.

package cloudevents

import (
	"time"
)

const (
	// Label holds the string label denoting the cloudevents type in the database.
	Label = "cloud_events"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNamespace holds the string denoting the namespace field in the database.
	FieldNamespace = "namespace"
	// FieldEvent holds the string denoting the event field in the database.
	FieldEvent = "event"
	// FieldFire holds the string denoting the fire field in the database.
	FieldFire = "fire"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldProcessed holds the string denoting the processed field in the database.
	FieldProcessed = "processed"
	// Table holds the table name of the cloudevents in the database.
	Table = "cloud_events"
)

// Columns holds all SQL columns for cloudevents fields.
var Columns = []string{
	FieldID,
	FieldNamespace,
	FieldEvent,
	FieldFire,
	FieldCreated,
	FieldProcessed,
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

var (
	// NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	NamespaceValidator func(string) error
	// DefaultFire holds the default value on creation for the "fire" field.
	DefaultFire func() time.Time
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)