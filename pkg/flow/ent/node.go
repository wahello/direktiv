// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
	"github.com/vorteil/direktiv/pkg/flow/ent/cloudevents"
	"github.com/vorteil/direktiv/pkg/flow/ent/events"
	"github.com/vorteil/direktiv/pkg/flow/ent/eventswait"
	"github.com/vorteil/direktiv/pkg/flow/ent/inode"
	"github.com/vorteil/direktiv/pkg/flow/ent/instance"
	"github.com/vorteil/direktiv/pkg/flow/ent/instanceruntime"
	"github.com/vorteil/direktiv/pkg/flow/ent/logmsg"
	"github.com/vorteil/direktiv/pkg/flow/ent/namespace"
	"github.com/vorteil/direktiv/pkg/flow/ent/ref"
	"github.com/vorteil/direktiv/pkg/flow/ent/revision"
	"github.com/vorteil/direktiv/pkg/flow/ent/route"
	"github.com/vorteil/direktiv/pkg/flow/ent/vardata"
	"github.com/vorteil/direktiv/pkg/flow/ent/varref"
	"github.com/vorteil/direktiv/pkg/flow/ent/workflow"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     uuid.UUID `json:"id,omitempty"`     // node id.
	Type   string    `json:"type,omitempty"`   // node type.
	Fields []*Field  `json:"fields,omitempty"` // node fields.
	Edges  []*Edge   `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string      `json:"type,omitempty"` // edge type.
	Name string      `json:"name,omitempty"` // edge name.
	IDs  []uuid.UUID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (ce *CloudEvents) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ce.ID,
		Type:   "CloudEvents",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(ce.EventId); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "eventId",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ce.Event); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "event.Event",
		Name:  "event",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ce.Fire); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "fire",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ce.Created); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "created",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ce.Processed); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "bool",
		Name:  "processed",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Namespace",
		Name: "namespace",
	}
	err = ce.QueryNamespace().
		Select(namespace.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (e *Events) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     e.ID,
		Type:   "Events",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(e.Events); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "[]map[string]interface {}",
		Name:  "events",
		Value: string(buf),
	}
	if buf, err = json.Marshal(e.Correlations); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "[]string",
		Name:  "correlations",
		Value: string(buf),
	}
	if buf, err = json.Marshal(e.Signature); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]byte",
		Name:  "signature",
		Value: string(buf),
	}
	if buf, err = json.Marshal(e.Count); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "int",
		Name:  "count",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Workflow",
		Name: "workflow",
	}
	err = e.QueryWorkflow().
		Select(workflow.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "EventsWait",
		Name: "wfeventswait",
	}
	err = e.QueryWfeventswait().
		Select(eventswait.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Instance",
		Name: "instance",
	}
	err = e.QueryInstance().
		Select(instance.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (ew *EventsWait) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ew.ID,
		Type:   "EventsWait",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(ew.Events); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "map[string]interface {}",
		Name:  "events",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Events",
		Name: "workflowevent",
	}
	err = ew.QueryWorkflowevent().
		Select(events.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (i *Inode) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     i.ID,
		Type:   "Inode",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 4),
	}
	var buf []byte
	if buf, err = json.Marshal(i.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.Type); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.Attributes); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "[]string",
		Name:  "attributes",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Namespace",
		Name: "namespace",
	}
	err = i.QueryNamespace().
		Select(namespace.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Inode",
		Name: "children",
	}
	err = i.QueryChildren().
		Select(inode.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Inode",
		Name: "parent",
	}
	err = i.QueryParent().
		Select(inode.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "Workflow",
		Name: "workflow",
	}
	err = i.QueryWorkflow().
		Select(workflow.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (i *Instance) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     i.ID,
		Type:   "Instance",
		Fields: make([]*Field, 7),
		Edges:  make([]*Edge, 8),
	}
	var buf []byte
	if buf, err = json.Marshal(i.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.EndAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "end_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.Status); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.As); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "as",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.ErrorCode); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "errorCode",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.ErrorMessage); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "errorMessage",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Namespace",
		Name: "namespace",
	}
	err = i.QueryNamespace().
		Select(namespace.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Workflow",
		Name: "workflow",
	}
	err = i.QueryWorkflow().
		Select(workflow.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Revision",
		Name: "revision",
	}
	err = i.QueryRevision().
		Select(revision.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "LogMsg",
		Name: "logs",
	}
	err = i.QueryLogs().
		Select(logmsg.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "VarRef",
		Name: "vars",
	}
	err = i.QueryVars().
		Select(varref.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[5] = &Edge{
		Type: "InstanceRuntime",
		Name: "runtime",
	}
	err = i.QueryRuntime().
		Select(instanceruntime.FieldID).
		Scan(ctx, &node.Edges[5].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[6] = &Edge{
		Type: "InstanceRuntime",
		Name: "children",
	}
	err = i.QueryChildren().
		Select(instanceruntime.FieldID).
		Scan(ctx, &node.Edges[6].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[7] = &Edge{
		Type: "Events",
		Name: "eventlisteners",
	}
	err = i.QueryEventlisteners().
		Select(events.FieldID).
		Scan(ctx, &node.Edges[7].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (ir *InstanceRuntime) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ir.ID,
		Type:   "InstanceRuntime",
		Fields: make([]*Field, 12),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(ir.Input); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "[]byte",
		Name:  "input",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ir.Data); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "data",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ir.Controller); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "controller",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ir.Memory); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "memory",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ir.Flow); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "[]string",
		Name:  "flow",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ir.Output); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "output",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ir.StateBeginTime); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "time.Time",
		Name:  "stateBeginTime",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ir.Deadline); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "time.Time",
		Name:  "deadline",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ir.Attempts); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "int",
		Name:  "attempts",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ir.CallerData); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "string",
		Name:  "caller_data",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ir.InstanceContext); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "string",
		Name:  "instanceContext",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ir.StateContext); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "string",
		Name:  "stateContext",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Instance",
		Name: "instance",
	}
	err = ir.QueryInstance().
		Select(instance.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Instance",
		Name: "caller",
	}
	err = ir.QueryCaller().
		Select(instance.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (lm *LogMsg) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     lm.ID,
		Type:   "LogMsg",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(lm.T); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "t",
		Value: string(buf),
	}
	if buf, err = json.Marshal(lm.Msg); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "msg",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Namespace",
		Name: "namespace",
	}
	err = lm.QueryNamespace().
		Select(namespace.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Workflow",
		Name: "workflow",
	}
	err = lm.QueryWorkflow().
		Select(workflow.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Instance",
		Name: "instance",
	}
	err = lm.QueryInstance().
		Select(instance.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (n *Namespace) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     n.ID,
		Type:   "Namespace",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 6),
	}
	var buf []byte
	if buf, err = json.Marshal(n.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(n.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(n.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Inode",
		Name: "inodes",
	}
	err = n.QueryInodes().
		Select(inode.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Workflow",
		Name: "workflows",
	}
	err = n.QueryWorkflows().
		Select(workflow.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Instance",
		Name: "instances",
	}
	err = n.QueryInstances().
		Select(instance.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "LogMsg",
		Name: "logs",
	}
	err = n.QueryLogs().
		Select(logmsg.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "VarRef",
		Name: "vars",
	}
	err = n.QueryVars().
		Select(varref.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[5] = &Edge{
		Type: "CloudEvents",
		Name: "cloudevents",
	}
	err = n.QueryCloudevents().
		Select(cloudevents.FieldID).
		Scan(ctx, &node.Edges[5].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (r *Ref) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     r.ID,
		Type:   "Ref",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(r.Immutable); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "bool",
		Name:  "immutable",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Workflow",
		Name: "workflow",
	}
	err = r.QueryWorkflow().
		Select(workflow.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Revision",
		Name: "revision",
	}
	err = r.QueryRevision().
		Select(revision.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Route",
		Name: "routes",
	}
	err = r.QueryRoutes().
		Select(route.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (r *Revision) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     r.ID,
		Type:   "Revision",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(r.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Hash); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "hash",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Source); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]byte",
		Name:  "source",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Workflow",
		Name: "workflow",
	}
	err = r.QueryWorkflow().
		Select(workflow.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Ref",
		Name: "refs",
	}
	err = r.QueryRefs().
		Select(ref.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Instance",
		Name: "instances",
	}
	err = r.QueryInstances().
		Select(instance.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (r *Route) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     r.ID,
		Type:   "Route",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(r.Weight); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "weight",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Workflow",
		Name: "workflow",
	}
	err = r.QueryWorkflow().
		Select(workflow.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Ref",
		Name: "ref",
	}
	err = r.QueryRef().
		Select(ref.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (vd *VarData) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     vd.ID,
		Type:   "VarData",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(vd.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(vd.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(vd.Size); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "size",
		Value: string(buf),
	}
	if buf, err = json.Marshal(vd.Hash); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "hash",
		Value: string(buf),
	}
	if buf, err = json.Marshal(vd.Data); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "[]byte",
		Name:  "data",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "VarRef",
		Name: "varrefs",
	}
	err = vd.QueryVarrefs().
		Select(varref.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (vr *VarRef) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     vr.ID,
		Type:   "VarRef",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 4),
	}
	var buf []byte
	if buf, err = json.Marshal(vr.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "VarData",
		Name: "vardata",
	}
	err = vr.QueryVardata().
		Select(vardata.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Namespace",
		Name: "namespace",
	}
	err = vr.QueryNamespace().
		Select(namespace.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Workflow",
		Name: "workflow",
	}
	err = vr.QueryWorkflow().
		Select(workflow.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "Instance",
		Name: "instance",
	}
	err = vr.QueryInstance().
		Select(instance.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (w *Workflow) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     w.ID,
		Type:   "Workflow",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 9),
	}
	var buf []byte
	if buf, err = json.Marshal(w.Live); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "bool",
		Name:  "live",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.LogToEvents); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "logToEvents",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Inode",
		Name: "inode",
	}
	err = w.QueryInode().
		Select(inode.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Namespace",
		Name: "namespace",
	}
	err = w.QueryNamespace().
		Select(namespace.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Revision",
		Name: "revisions",
	}
	err = w.QueryRevisions().
		Select(revision.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "Ref",
		Name: "refs",
	}
	err = w.QueryRefs().
		Select(ref.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "Instance",
		Name: "instances",
	}
	err = w.QueryInstances().
		Select(instance.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[5] = &Edge{
		Type: "Route",
		Name: "routes",
	}
	err = w.QueryRoutes().
		Select(route.FieldID).
		Scan(ctx, &node.Edges[5].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[6] = &Edge{
		Type: "LogMsg",
		Name: "logs",
	}
	err = w.QueryLogs().
		Select(logmsg.FieldID).
		Scan(ctx, &node.Edges[6].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[7] = &Edge{
		Type: "VarRef",
		Name: "vars",
	}
	err = w.QueryVars().
		Select(varref.FieldID).
		Scan(ctx, &node.Edges[7].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[8] = &Edge{
		Type: "Events",
		Name: "wfevents",
	}
	err = w.QueryWfevents().
		Select(events.FieldID).
		Scan(ctx, &node.Edges[8].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id uuid.UUID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, uuid.UUID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, uuid.UUID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, uuid.UUID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id uuid.UUID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//		c.Noder(ctx, id)
//		c.Noder(ctx, id, ent.WithNodeType(pet.Table))
//
func (c *Client) Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id uuid.UUID) (Noder, error) {
	switch table {
	case cloudevents.Table:
		n, err := c.CloudEvents.Query().
			Where(cloudevents.ID(id)).
			CollectFields(ctx, "CloudEvents").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case events.Table:
		n, err := c.Events.Query().
			Where(events.ID(id)).
			CollectFields(ctx, "Events").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case eventswait.Table:
		n, err := c.EventsWait.Query().
			Where(eventswait.ID(id)).
			CollectFields(ctx, "EventsWait").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case inode.Table:
		n, err := c.Inode.Query().
			Where(inode.ID(id)).
			CollectFields(ctx, "Inode").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case instance.Table:
		n, err := c.Instance.Query().
			Where(instance.ID(id)).
			CollectFields(ctx, "Instance").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case instanceruntime.Table:
		n, err := c.InstanceRuntime.Query().
			Where(instanceruntime.ID(id)).
			CollectFields(ctx, "InstanceRuntime").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case logmsg.Table:
		n, err := c.LogMsg.Query().
			Where(logmsg.ID(id)).
			CollectFields(ctx, "LogMsg").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case namespace.Table:
		n, err := c.Namespace.Query().
			Where(namespace.ID(id)).
			CollectFields(ctx, "Namespace").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case ref.Table:
		n, err := c.Ref.Query().
			Where(ref.ID(id)).
			CollectFields(ctx, "Ref").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case revision.Table:
		n, err := c.Revision.Query().
			Where(revision.ID(id)).
			CollectFields(ctx, "Revision").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case route.Table:
		n, err := c.Route.Query().
			Where(route.ID(id)).
			CollectFields(ctx, "Route").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case vardata.Table:
		n, err := c.VarData.Query().
			Where(vardata.ID(id)).
			CollectFields(ctx, "VarData").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case varref.Table:
		n, err := c.VarRef.Query().
			Where(varref.ID(id)).
			CollectFields(ctx, "VarRef").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case workflow.Table:
		n, err := c.Workflow.Query().
			Where(workflow.ID(id)).
			CollectFields(ctx, "Workflow").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []uuid.UUID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]uuid.UUID)
	id2idx := make(map[uuid.UUID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []uuid.UUID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[uuid.UUID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case cloudevents.Table:
		nodes, err := c.CloudEvents.Query().
			Where(cloudevents.IDIn(ids...)).
			CollectFields(ctx, "CloudEvents").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case events.Table:
		nodes, err := c.Events.Query().
			Where(events.IDIn(ids...)).
			CollectFields(ctx, "Events").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case eventswait.Table:
		nodes, err := c.EventsWait.Query().
			Where(eventswait.IDIn(ids...)).
			CollectFields(ctx, "EventsWait").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case inode.Table:
		nodes, err := c.Inode.Query().
			Where(inode.IDIn(ids...)).
			CollectFields(ctx, "Inode").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case instance.Table:
		nodes, err := c.Instance.Query().
			Where(instance.IDIn(ids...)).
			CollectFields(ctx, "Instance").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case instanceruntime.Table:
		nodes, err := c.InstanceRuntime.Query().
			Where(instanceruntime.IDIn(ids...)).
			CollectFields(ctx, "InstanceRuntime").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case logmsg.Table:
		nodes, err := c.LogMsg.Query().
			Where(logmsg.IDIn(ids...)).
			CollectFields(ctx, "LogMsg").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case namespace.Table:
		nodes, err := c.Namespace.Query().
			Where(namespace.IDIn(ids...)).
			CollectFields(ctx, "Namespace").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case ref.Table:
		nodes, err := c.Ref.Query().
			Where(ref.IDIn(ids...)).
			CollectFields(ctx, "Ref").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case revision.Table:
		nodes, err := c.Revision.Query().
			Where(revision.IDIn(ids...)).
			CollectFields(ctx, "Revision").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case route.Table:
		nodes, err := c.Route.Query().
			Where(route.IDIn(ids...)).
			CollectFields(ctx, "Route").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case vardata.Table:
		nodes, err := c.VarData.Query().
			Where(vardata.IDIn(ids...)).
			CollectFields(ctx, "VarData").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case varref.Table:
		nodes, err := c.VarRef.Query().
			Where(varref.IDIn(ids...)).
			CollectFields(ctx, "VarRef").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workflow.Table:
		nodes, err := c.Workflow.Query().
			Where(workflow.IDIn(ids...)).
			CollectFields(ctx, "Workflow").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
