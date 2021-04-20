// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/vorteil/direktiv/pkg/metrics/ent/metrics"
	"github.com/vorteil/direktiv/pkg/metrics/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	metricsFields := schema.Metrics{}.Fields()
	_ = metricsFields
	// metricsDescNamespace is the schema descriptor for namespace field.
	metricsDescNamespace := metricsFields[0].Descriptor()
	// metrics.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	metrics.NamespaceValidator = metricsDescNamespace.Validators[0].(func(string) error)
	// metricsDescWorkflow is the schema descriptor for workflow field.
	metricsDescWorkflow := metricsFields[1].Descriptor()
	// metrics.WorkflowValidator is a validator for the "workflow" field. It is called by the builders before save.
	metrics.WorkflowValidator = metricsDescWorkflow.Validators[0].(func(string) error)
	// metricsDescInstance is the schema descriptor for instance field.
	metricsDescInstance := metricsFields[2].Descriptor()
	// metrics.InstanceValidator is a validator for the "instance" field. It is called by the builders before save.
	metrics.InstanceValidator = metricsDescInstance.Validators[0].(func(string) error)
	// metricsDescState is the schema descriptor for state field.
	metricsDescState := metricsFields[3].Descriptor()
	// metrics.StateValidator is a validator for the "state" field. It is called by the builders before save.
	metrics.StateValidator = metricsDescState.Validators[0].(func(string) error)
	// metricsDescWorkflowMs is the schema descriptor for workflow_ms field.
	metricsDescWorkflowMs := metricsFields[5].Descriptor()
	// metrics.WorkflowMsValidator is a validator for the "workflow_ms" field. It is called by the builders before save.
	metrics.WorkflowMsValidator = metricsDescWorkflowMs.Validators[0].(func(int64) error)
	// metricsDescIsolateMs is the schema descriptor for isolate_ms field.
	metricsDescIsolateMs := metricsFields[6].Descriptor()
	// metrics.IsolateMsValidator is a validator for the "isolate_ms" field. It is called by the builders before save.
	metrics.IsolateMsValidator = metricsDescIsolateMs.Validators[0].(func(int64) error)
	// metricsDescNext is the schema descriptor for next field.
	metricsDescNext := metricsFields[9].Descriptor()
	// metrics.NextValidator is a validator for the "next" field. It is called by the builders before save.
	metrics.NextValidator = func() func(int8) error {
		validators := metricsDescNext.Validators
		fns := [...]func(int8) error{
			validators[0].(func(int8) error),
			validators[1].(func(int8) error),
		}
		return func(next int8) error {
			for _, fn := range fns {
				if err := fn(next); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
