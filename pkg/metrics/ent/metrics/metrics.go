// Code generated by entc, DO NOT EDIT.

package metrics

const (
	// Label holds the string label denoting the metrics type in the database.
	Label = "metrics"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNamespace holds the string denoting the namespace field in the database.
	FieldNamespace = "namespace"
	// FieldWorkflow holds the string denoting the workflow field in the database.
	FieldWorkflow = "workflow"
	// FieldInstance holds the string denoting the instance field in the database.
	FieldInstance = "instance"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldWorkflowMs holds the string denoting the workflow_ms field in the database.
	FieldWorkflowMs = "workflow_ms"
	// FieldIsolateMs holds the string denoting the isolate_ms field in the database.
	FieldIsolateMs = "isolate_ms"
	// FieldErrorCode holds the string denoting the error_code field in the database.
	FieldErrorCode = "error_code"
	// FieldInvoker holds the string denoting the invoker field in the database.
	FieldInvoker = "invoker"
	// FieldNext holds the string denoting the next field in the database.
	FieldNext = "next"
	// FieldTransition holds the string denoting the transition field in the database.
	FieldTransition = "transition"
	// Table holds the table name of the metrics in the database.
	Table = "metrics"
)

// Columns holds all SQL columns for metrics fields.
var Columns = []string{
	FieldID,
	FieldNamespace,
	FieldWorkflow,
	FieldInstance,
	FieldState,
	FieldTimestamp,
	FieldWorkflowMs,
	FieldIsolateMs,
	FieldErrorCode,
	FieldInvoker,
	FieldNext,
	FieldTransition,
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
	// WorkflowValidator is a validator for the "workflow" field. It is called by the builders before save.
	WorkflowValidator func(string) error
	// InstanceValidator is a validator for the "instance" field. It is called by the builders before save.
	InstanceValidator func(string) error
	// StateValidator is a validator for the "state" field. It is called by the builders before save.
	StateValidator func(string) error
	// WorkflowMsValidator is a validator for the "workflow_ms" field. It is called by the builders before save.
	WorkflowMsValidator func(int64) error
	// IsolateMsValidator is a validator for the "isolate_ms" field. It is called by the builders before save.
	IsolateMsValidator func(int64) error
	// NextValidator is a validator for the "next" field. It is called by the builders before save.
	NextValidator func(int8) error
)
