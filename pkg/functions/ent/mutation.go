// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/vorteil/direktiv/pkg/functions/ent/predicate"
	"github.com/vorteil/direktiv/pkg/functions/ent/services"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeServices = "Services"
)

// ServicesMutation represents an operation that mutates the Services nodes in the graph.
type ServicesMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	data          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Services, error)
	predicates    []predicate.Services
}

var _ ent.Mutation = (*ServicesMutation)(nil)

// servicesOption allows management of the mutation configuration using functional options.
type servicesOption func(*ServicesMutation)

// newServicesMutation creates new mutation for the Services entity.
func newServicesMutation(c config, op Op, opts ...servicesOption) *ServicesMutation {
	m := &ServicesMutation{
		config:        c,
		op:            op,
		typ:           TypeServices,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withServicesID sets the ID field of the mutation.
func withServicesID(id int) servicesOption {
	return func(m *ServicesMutation) {
		var (
			err   error
			once  sync.Once
			value *Services
		)
		m.oldValue = func(ctx context.Context) (*Services, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Services.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withServices sets the old Services of the mutation.
func withServices(node *Services) servicesOption {
	return func(m *ServicesMutation) {
		m.oldValue = func(context.Context) (*Services, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ServicesMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ServicesMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ServicesMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *ServicesMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ServicesMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Services entity.
// If the Services object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServicesMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ServicesMutation) ResetName() {
	m.name = nil
}

// SetData sets the "data" field.
func (m *ServicesMutation) SetData(s string) {
	m.data = &s
}

// Data returns the value of the "data" field in the mutation.
func (m *ServicesMutation) Data() (r string, exists bool) {
	v := m.data
	if v == nil {
		return
	}
	return *v, true
}

// OldData returns the old "data" field's value of the Services entity.
// If the Services object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServicesMutation) OldData(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldData is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldData requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldData: %w", err)
	}
	return oldValue.Data, nil
}

// ResetData resets all changes to the "data" field.
func (m *ServicesMutation) ResetData() {
	m.data = nil
}

// Where appends a list predicates to the ServicesMutation builder.
func (m *ServicesMutation) Where(ps ...predicate.Services) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ServicesMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Services).
func (m *ServicesMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ServicesMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, services.FieldName)
	}
	if m.data != nil {
		fields = append(fields, services.FieldData)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ServicesMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case services.FieldName:
		return m.Name()
	case services.FieldData:
		return m.Data()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ServicesMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case services.FieldName:
		return m.OldName(ctx)
	case services.FieldData:
		return m.OldData(ctx)
	}
	return nil, fmt.Errorf("unknown Services field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ServicesMutation) SetField(name string, value ent.Value) error {
	switch name {
	case services.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case services.FieldData:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetData(v)
		return nil
	}
	return fmt.Errorf("unknown Services field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ServicesMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ServicesMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ServicesMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Services numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ServicesMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ServicesMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ServicesMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Services nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ServicesMutation) ResetField(name string) error {
	switch name {
	case services.FieldName:
		m.ResetName()
		return nil
	case services.FieldData:
		m.ResetData()
		return nil
	}
	return fmt.Errorf("unknown Services field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ServicesMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ServicesMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ServicesMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ServicesMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ServicesMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ServicesMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ServicesMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Services unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ServicesMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Services edge %s", name)
}
