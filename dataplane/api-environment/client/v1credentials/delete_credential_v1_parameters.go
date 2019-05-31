// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

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

// NewDeleteCredentialV1Params creates a new DeleteCredentialV1Params object
// with the default values initialized.
func NewDeleteCredentialV1Params() *DeleteCredentialV1Params {
	var ()
	return &DeleteCredentialV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCredentialV1ParamsWithTimeout creates a new DeleteCredentialV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCredentialV1ParamsWithTimeout(timeout time.Duration) *DeleteCredentialV1Params {
	var ()
	return &DeleteCredentialV1Params{

		timeout: timeout,
	}
}

// NewDeleteCredentialV1ParamsWithContext creates a new DeleteCredentialV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCredentialV1ParamsWithContext(ctx context.Context) *DeleteCredentialV1Params {
	var ()
	return &DeleteCredentialV1Params{

		Context: ctx,
	}
}

// NewDeleteCredentialV1ParamsWithHTTPClient creates a new DeleteCredentialV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCredentialV1ParamsWithHTTPClient(client *http.Client) *DeleteCredentialV1Params {
	var ()
	return &DeleteCredentialV1Params{
		HTTPClient: client,
	}
}

/*DeleteCredentialV1Params contains all the parameters to send to the API endpoint
for the delete credential v1 operation typically these are written to a http.Request
*/
type DeleteCredentialV1Params struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete credential v1 params
func (o *DeleteCredentialV1Params) WithTimeout(timeout time.Duration) *DeleteCredentialV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete credential v1 params
func (o *DeleteCredentialV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete credential v1 params
func (o *DeleteCredentialV1Params) WithContext(ctx context.Context) *DeleteCredentialV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete credential v1 params
func (o *DeleteCredentialV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete credential v1 params
func (o *DeleteCredentialV1Params) WithHTTPClient(client *http.Client) *DeleteCredentialV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete credential v1 params
func (o *DeleteCredentialV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete credential v1 params
func (o *DeleteCredentialV1Params) WithName(name string) *DeleteCredentialV1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete credential v1 params
func (o *DeleteCredentialV1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCredentialV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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