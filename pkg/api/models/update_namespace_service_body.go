// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateNamespaceServiceBody UpdateNamespaceServiceBody UpdateNamespaceServiceBody update namespace service body
// Example: {"cmd":"","image":"direktiv/request:v10","minScale":"1","size":"small","trafficPercent":50}
//
// swagger:model UpdateNamespaceServiceBody
type UpdateNamespaceServiceBody struct {

	// cmd
	// Required: true
	Cmd *string `json:"cmd"`

	// Target image a service will use
	// Required: true
	Image *string `json:"image"`

	// Minimum amount of service pods to be live
	// Required: true
	MinScale *int64 `json:"minScale"`

	// Size of created service pods
	// Required: true
	// Enum: [[[small medium large]]]
	Size *string `json:"size"`

	// Traffic percentage new revision will use
	// Required: true
	TrafficPercent *int64 `json:"trafficPercent"`
}

// Validate validates this update namespace service body
func (m *UpdateNamespaceServiceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCmd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinScale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrafficPercent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNamespaceServiceBody) validateCmd(formats strfmt.Registry) error {

	if err := validate.Required("cmd", "body", m.Cmd); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNamespaceServiceBody) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNamespaceServiceBody) validateMinScale(formats strfmt.Registry) error {

	if err := validate.Required("minScale", "body", m.MinScale); err != nil {
		return err
	}

	return nil
}

var updateNamespaceServiceBodyTypeSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["[[small medium large]]"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNamespaceServiceBodyTypeSizePropEnum = append(updateNamespaceServiceBodyTypeSizePropEnum, v)
	}
}

const (

	// UpdateNamespaceServiceBodySizeSmallMediumLarge captures enum value "[[small medium large]]"
	UpdateNamespaceServiceBodySizeSmallMediumLarge string = "[[small medium large]]"
)

// prop value enum
func (m *UpdateNamespaceServiceBody) validateSizeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNamespaceServiceBodyTypeSizePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateNamespaceServiceBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	// value enum
	if err := m.validateSizeEnum("size", "body", *m.Size); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNamespaceServiceBody) validateTrafficPercent(formats strfmt.Registry) error {

	if err := validate.Required("trafficPercent", "body", m.TrafficPercent); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update namespace service body based on context it is used
func (m *UpdateNamespaceServiceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNamespaceServiceBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNamespaceServiceBody) UnmarshalBinary(b []byte) error {
	var res UpdateNamespaceServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}