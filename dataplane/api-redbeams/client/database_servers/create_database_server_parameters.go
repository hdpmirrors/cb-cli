// Code generated by go-swagger; DO NOT EDIT.

package database_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
)

// NewCreateDatabaseServerParams creates a new CreateDatabaseServerParams object
// with the default values initialized.
func NewCreateDatabaseServerParams() *CreateDatabaseServerParams {
	var ()
	return &CreateDatabaseServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDatabaseServerParamsWithTimeout creates a new CreateDatabaseServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDatabaseServerParamsWithTimeout(timeout time.Duration) *CreateDatabaseServerParams {
	var ()
	return &CreateDatabaseServerParams{

		timeout: timeout,
	}
}

// NewCreateDatabaseServerParamsWithContext creates a new CreateDatabaseServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDatabaseServerParamsWithContext(ctx context.Context) *CreateDatabaseServerParams {
	var ()
	return &CreateDatabaseServerParams{

		Context: ctx,
	}
}

// NewCreateDatabaseServerParamsWithHTTPClient creates a new CreateDatabaseServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDatabaseServerParamsWithHTTPClient(client *http.Client) *CreateDatabaseServerParams {
	var ()
	return &CreateDatabaseServerParams{
		HTTPClient: client,
	}
}

/*CreateDatabaseServerParams contains all the parameters to send to the API endpoint
for the create database server operation typically these are written to a http.Request
*/
type CreateDatabaseServerParams struct {

	/*Body
	  Request for allocating a new database server in a provider

	*/
	Body *model.AllocateDatabaseServerV4Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create database server params
func (o *CreateDatabaseServerParams) WithTimeout(timeout time.Duration) *CreateDatabaseServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create database server params
func (o *CreateDatabaseServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create database server params
func (o *CreateDatabaseServerParams) WithContext(ctx context.Context) *CreateDatabaseServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create database server params
func (o *CreateDatabaseServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create database server params
func (o *CreateDatabaseServerParams) WithHTTPClient(client *http.Client) *CreateDatabaseServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create database server params
func (o *CreateDatabaseServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create database server params
func (o *CreateDatabaseServerParams) WithBody(body *model.AllocateDatabaseServerV4Request) *CreateDatabaseServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create database server params
func (o *CreateDatabaseServerParams) SetBody(body *model.AllocateDatabaseServerV4Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDatabaseServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}