// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/vorteil/direktiv/pkg/flow/ent/inode"
	"github.com/vorteil/direktiv/pkg/flow/ent/namespace"
	"github.com/vorteil/direktiv/pkg/flow/ent/workflow"
)

// Inode is the model entity for the Inode schema.
type Inode struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Attributes holds the value of the "attributes" field.
	Attributes []string `json:"attributes,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InodeQuery when eager-loading is set.
	Edges            InodeEdges `json:"edges"`
	inode_children   *uuid.UUID
	namespace_inodes *uuid.UUID
}

// InodeEdges holds the relations/edges for other nodes in the graph.
type InodeEdges struct {
	// Namespace holds the value of the namespace edge.
	Namespace *Namespace `json:"namespace,omitempty"`
	// Children holds the value of the children edge.
	Children []*Inode `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Inode `json:"parent,omitempty"`
	// Workflow holds the value of the workflow edge.
	Workflow *Workflow `json:"workflow,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// NamespaceOrErr returns the Namespace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InodeEdges) NamespaceOrErr() (*Namespace, error) {
	if e.loadedTypes[0] {
		if e.Namespace == nil {
			// The edge namespace was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: namespace.Label}
		}
		return e.Namespace, nil
	}
	return nil, &NotLoadedError{edge: "namespace"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e InodeEdges) ChildrenOrErr() ([]*Inode, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InodeEdges) ParentOrErr() (*Inode, error) {
	if e.loadedTypes[2] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: inode.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InodeEdges) WorkflowOrErr() (*Workflow, error) {
	if e.loadedTypes[3] {
		if e.Workflow == nil {
			// The edge workflow was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workflow.Label}
		}
		return e.Workflow, nil
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Inode) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case inode.FieldAttributes:
			values[i] = new([]byte)
		case inode.FieldName, inode.FieldType:
			values[i] = new(sql.NullString)
		case inode.FieldCreatedAt, inode.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case inode.FieldID:
			values[i] = new(uuid.UUID)
		case inode.ForeignKeys[0]: // inode_children
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case inode.ForeignKeys[1]: // namespace_inodes
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Inode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Inode fields.
func (i *Inode) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case inode.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case inode.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case inode.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case inode.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case inode.FieldType:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[j])
			} else if value.Valid {
				i.Type = value.String
			}
		case inode.FieldAttributes:
			if value, ok := values[j].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field attributes", values[j])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &i.Attributes); err != nil {
					return fmt.Errorf("unmarshal field attributes: %w", err)
				}
			}
		case inode.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field inode_children", values[j])
			} else if value.Valid {
				i.inode_children = new(uuid.UUID)
				*i.inode_children = *value.S.(*uuid.UUID)
			}
		case inode.ForeignKeys[1]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field namespace_inodes", values[j])
			} else if value.Valid {
				i.namespace_inodes = new(uuid.UUID)
				*i.namespace_inodes = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryNamespace queries the "namespace" edge of the Inode entity.
func (i *Inode) QueryNamespace() *NamespaceQuery {
	return (&InodeClient{config: i.config}).QueryNamespace(i)
}

// QueryChildren queries the "children" edge of the Inode entity.
func (i *Inode) QueryChildren() *InodeQuery {
	return (&InodeClient{config: i.config}).QueryChildren(i)
}

// QueryParent queries the "parent" edge of the Inode entity.
func (i *Inode) QueryParent() *InodeQuery {
	return (&InodeClient{config: i.config}).QueryParent(i)
}

// QueryWorkflow queries the "workflow" edge of the Inode entity.
func (i *Inode) QueryWorkflow() *WorkflowQuery {
	return (&InodeClient{config: i.config}).QueryWorkflow(i)
}

// Update returns a builder for updating this Inode.
// Note that you need to call Inode.Unwrap() before calling this method if this Inode
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Inode) Update() *InodeUpdateOne {
	return (&InodeClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Inode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Inode) Unwrap() *Inode {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Inode is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Inode) String() string {
	var builder strings.Builder
	builder.WriteString("Inode(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(i.Name)
	builder.WriteString(", type=")
	builder.WriteString(i.Type)
	builder.WriteString(", attributes=")
	builder.WriteString(fmt.Sprintf("%v", i.Attributes))
	builder.WriteByte(')')
	return builder.String()
}

// Inodes is a parsable slice of Inode.
type Inodes []*Inode

func (i Inodes) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
