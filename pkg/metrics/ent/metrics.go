// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/vorteil/direktiv/pkg/metrics/ent/metrics"
)

// Metrics is the model entity for the Metrics schema.
type Metrics struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Namespace holds the value of the "namespace" field.
	Namespace string `json:"namespace,omitempty"`
	// Workflow holds the value of the "workflow" field.
	Workflow string `json:"workflow,omitempty"`
	// Instance holds the value of the "instance" field.
	Instance string `json:"instance,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// WorkflowMs holds the value of the "workflow_ms" field.
	WorkflowMs int64 `json:"workflow_ms,omitempty"`
	// IsolateMs holds the value of the "isolate_ms" field.
	IsolateMs int64 `json:"isolate_ms,omitempty"`
	// ErrorCode holds the value of the "error_code" field.
	ErrorCode string `json:"error_code,omitempty"`
	// Invoker holds the value of the "invoker" field.
	Invoker string `json:"invoker,omitempty"`
	// Next holds the value of the "next" field.
	Next int8 `json:"next,omitempty"`
	// Transition holds the value of the "transition" field.
	Transition string `json:"transition,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Metrics) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case metrics.FieldID, metrics.FieldWorkflowMs, metrics.FieldIsolateMs, metrics.FieldNext:
			values[i] = &sql.NullInt64{}
		case metrics.FieldNamespace, metrics.FieldWorkflow, metrics.FieldInstance, metrics.FieldState, metrics.FieldErrorCode, metrics.FieldInvoker, metrics.FieldTransition:
			values[i] = &sql.NullString{}
		case metrics.FieldTimestamp:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Metrics", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Metrics fields.
func (m *Metrics) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case metrics.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case metrics.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				m.Namespace = value.String
			}
		case metrics.FieldWorkflow:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field workflow", values[i])
			} else if value.Valid {
				m.Workflow = value.String
			}
		case metrics.FieldInstance:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instance", values[i])
			} else if value.Valid {
				m.Instance = value.String
			}
		case metrics.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				m.State = value.String
			}
		case metrics.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				m.Timestamp = value.Time
			}
		case metrics.FieldWorkflowMs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_ms", values[i])
			} else if value.Valid {
				m.WorkflowMs = value.Int64
			}
		case metrics.FieldIsolateMs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field isolate_ms", values[i])
			} else if value.Valid {
				m.IsolateMs = value.Int64
			}
		case metrics.FieldErrorCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error_code", values[i])
			} else if value.Valid {
				m.ErrorCode = value.String
			}
		case metrics.FieldInvoker:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invoker", values[i])
			} else if value.Valid {
				m.Invoker = value.String
			}
		case metrics.FieldNext:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field next", values[i])
			} else if value.Valid {
				m.Next = int8(value.Int64)
			}
		case metrics.FieldTransition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transition", values[i])
			} else if value.Valid {
				m.Transition = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Metrics.
// Note that you need to call Metrics.Unwrap() before calling this method if this Metrics
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Metrics) Update() *MetricsUpdateOne {
	return (&MetricsClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Metrics entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Metrics) Unwrap() *Metrics {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Metrics is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Metrics) String() string {
	var builder strings.Builder
	builder.WriteString("Metrics(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", namespace=")
	builder.WriteString(m.Namespace)
	builder.WriteString(", workflow=")
	builder.WriteString(m.Workflow)
	builder.WriteString(", instance=")
	builder.WriteString(m.Instance)
	builder.WriteString(", state=")
	builder.WriteString(m.State)
	builder.WriteString(", timestamp=")
	builder.WriteString(m.Timestamp.Format(time.ANSIC))
	builder.WriteString(", workflow_ms=")
	builder.WriteString(fmt.Sprintf("%v", m.WorkflowMs))
	builder.WriteString(", isolate_ms=")
	builder.WriteString(fmt.Sprintf("%v", m.IsolateMs))
	builder.WriteString(", error_code=")
	builder.WriteString(m.ErrorCode)
	builder.WriteString(", invoker=")
	builder.WriteString(m.Invoker)
	builder.WriteString(", next=")
	builder.WriteString(fmt.Sprintf("%v", m.Next))
	builder.WriteString(", transition=")
	builder.WriteString(m.Transition)
	builder.WriteByte(')')
	return builder.String()
}

// MetricsSlice is a parsable slice of Metrics.
type MetricsSlice []*Metrics

func (m MetricsSlice) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}