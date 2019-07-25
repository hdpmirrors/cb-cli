// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStatusDistroXV1ByNameParams creates a new StatusDistroXV1ByNameParams object
// with the default values initialized.
func NewStatusDistroXV1ByNameParams() *StatusDistroXV1ByNameParams {
	var ()
	return &StatusDistroXV1ByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStatusDistroXV1ByNameParamsWithTimeout creates a new StatusDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStatusDistroXV1ByNameParamsWithTimeout(timeout time.Duration) *StatusDistroXV1ByNameParams {
	var ()
	return &StatusDistroXV1ByNameParams{

		timeout: timeout,
	}
}

// NewStatusDistroXV1ByNameParamsWithContext creates a new StatusDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewStatusDistroXV1ByNameParamsWithContext(ctx context.Context) *StatusDistroXV1ByNameParams {
	var ()
	return &StatusDistroXV1ByNameParams{

		Context: ctx,
	}
}

// NewStatusDistroXV1ByNameParamsWithHTTPClient creates a new StatusDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStatusDistroXV1ByNameParamsWithHTTPClient(client *http.Client) *StatusDistroXV1ByNameParams {
	var ()
	return &StatusDistroXV1ByNameParams{
		HTTPClient: client,
	}
}

/*StatusDistroXV1ByNameParams contains all the parameters to send to the API endpoint
for the status distro x v1 by name operation typically these are written to a http.Request
*/
type StatusDistroXV1ByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the status distro x v1 by name params
func (o *StatusDistroXV1ByNameParams) WithTimeout(timeout time.Duration) *StatusDistroXV1ByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status distro x v1 by name params
func (o *StatusDistroXV1ByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status distro x v1 by name params
func (o *StatusDistroXV1ByNameParams) WithContext(ctx context.Context) *StatusDistroXV1ByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status distro x v1 by name params
func (o *StatusDistroXV1ByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status distro x v1 by name params
func (o *StatusDistroXV1ByNameParams) WithHTTPClient(client *http.Client) *StatusDistroXV1ByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status distro x v1 by name params
func (o *StatusDistroXV1ByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the status distro x v1 by name params
func (o *StatusDistroXV1ByNameParams) WithName(name string) *StatusDistroXV1ByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the status distro x v1 by name params
func (o *StatusDistroXV1ByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *StatusDistroXV1ByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
