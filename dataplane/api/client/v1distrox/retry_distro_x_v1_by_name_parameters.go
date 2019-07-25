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

// NewRetryDistroXV1ByNameParams creates a new RetryDistroXV1ByNameParams object
// with the default values initialized.
func NewRetryDistroXV1ByNameParams() *RetryDistroXV1ByNameParams {
	var ()
	return &RetryDistroXV1ByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetryDistroXV1ByNameParamsWithTimeout creates a new RetryDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetryDistroXV1ByNameParamsWithTimeout(timeout time.Duration) *RetryDistroXV1ByNameParams {
	var ()
	return &RetryDistroXV1ByNameParams{

		timeout: timeout,
	}
}

// NewRetryDistroXV1ByNameParamsWithContext creates a new RetryDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetryDistroXV1ByNameParamsWithContext(ctx context.Context) *RetryDistroXV1ByNameParams {
	var ()
	return &RetryDistroXV1ByNameParams{

		Context: ctx,
	}
}

// NewRetryDistroXV1ByNameParamsWithHTTPClient creates a new RetryDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetryDistroXV1ByNameParamsWithHTTPClient(client *http.Client) *RetryDistroXV1ByNameParams {
	var ()
	return &RetryDistroXV1ByNameParams{
		HTTPClient: client,
	}
}

/*RetryDistroXV1ByNameParams contains all the parameters to send to the API endpoint
for the retry distro x v1 by name operation typically these are written to a http.Request
*/
type RetryDistroXV1ByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retry distro x v1 by name params
func (o *RetryDistroXV1ByNameParams) WithTimeout(timeout time.Duration) *RetryDistroXV1ByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retry distro x v1 by name params
func (o *RetryDistroXV1ByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retry distro x v1 by name params
func (o *RetryDistroXV1ByNameParams) WithContext(ctx context.Context) *RetryDistroXV1ByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retry distro x v1 by name params
func (o *RetryDistroXV1ByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retry distro x v1 by name params
func (o *RetryDistroXV1ByNameParams) WithHTTPClient(client *http.Client) *RetryDistroXV1ByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retry distro x v1 by name params
func (o *RetryDistroXV1ByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the retry distro x v1 by name params
func (o *RetryDistroXV1ByNameParams) WithName(name string) *RetryDistroXV1ByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the retry distro x v1 by name params
func (o *RetryDistroXV1ByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RetryDistroXV1ByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
