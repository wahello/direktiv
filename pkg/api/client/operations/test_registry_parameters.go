// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewTestRegistryParams creates a new TestRegistryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestRegistryParams() *TestRegistryParams {
	return &TestRegistryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestRegistryParamsWithTimeout creates a new TestRegistryParams object
// with the ability to set a timeout on a request.
func NewTestRegistryParamsWithTimeout(timeout time.Duration) *TestRegistryParams {
	return &TestRegistryParams{
		timeout: timeout,
	}
}

// NewTestRegistryParamsWithContext creates a new TestRegistryParams object
// with the ability to set a context for a request.
func NewTestRegistryParamsWithContext(ctx context.Context) *TestRegistryParams {
	return &TestRegistryParams{
		Context: ctx,
	}
}

// NewTestRegistryParamsWithHTTPClient creates a new TestRegistryParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestRegistryParamsWithHTTPClient(client *http.Client) *TestRegistryParams {
	return &TestRegistryParams{
		HTTPClient: client,
	}
}

/* TestRegistryParams contains all the parameters to send to the API endpoint
   for the test registry operation.

   Typically these are written to a http.Request.
*/
type TestRegistryParams struct {

	/* RegistryPayload.

	   Payload that contains registry data
	*/
	RegistryPayload TestRegistryBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestRegistryParams) WithDefaults() *TestRegistryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestRegistryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test registry params
func (o *TestRegistryParams) WithTimeout(timeout time.Duration) *TestRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test registry params
func (o *TestRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test registry params
func (o *TestRegistryParams) WithContext(ctx context.Context) *TestRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test registry params
func (o *TestRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test registry params
func (o *TestRegistryParams) WithHTTPClient(client *http.Client) *TestRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test registry params
func (o *TestRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegistryPayload adds the registryPayload to the test registry params
func (o *TestRegistryParams) WithRegistryPayload(registryPayload TestRegistryBody) *TestRegistryParams {
	o.SetRegistryPayload(registryPayload)
	return o
}

// SetRegistryPayload adds the registryPayload to the test registry params
func (o *TestRegistryParams) SetRegistryPayload(registryPayload TestRegistryBody) {
	o.RegistryPayload = registryPayload
}

// WriteToRequest writes these params to a swagger request
func (o *TestRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.RegistryPayload); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
