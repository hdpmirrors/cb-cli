// Code generated by go-swagger; DO NOT EDIT.

package freeipa_environment_name

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

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// NewCreateFreeIPAParams creates a new CreateFreeIPAParams object
// with the default values initialized.
func NewCreateFreeIPAParams() *CreateFreeIPAParams {
	var ()
	return &CreateFreeIPAParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFreeIPAParamsWithTimeout creates a new CreateFreeIPAParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateFreeIPAParamsWithTimeout(timeout time.Duration) *CreateFreeIPAParams {
	var ()
	return &CreateFreeIPAParams{

		timeout: timeout,
	}
}

// NewCreateFreeIPAParamsWithContext creates a new CreateFreeIPAParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateFreeIPAParamsWithContext(ctx context.Context) *CreateFreeIPAParams {
	var ()
	return &CreateFreeIPAParams{

		Context: ctx,
	}
}

// NewCreateFreeIPAParamsWithHTTPClient creates a new CreateFreeIPAParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateFreeIPAParamsWithHTTPClient(client *http.Client) *CreateFreeIPAParams {
	var ()
	return &CreateFreeIPAParams{
		HTTPClient: client,
	}
}

/*CreateFreeIPAParams contains all the parameters to send to the API endpoint
for the create free IP a operation typically these are written to a http.Request
*/
type CreateFreeIPAParams struct {

	/*Body*/
	Body *model.CreateFreeIpaRequest
	/*EnvironmentName*/
	EnvironmentName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create free IP a params
func (o *CreateFreeIPAParams) WithTimeout(timeout time.Duration) *CreateFreeIPAParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create free IP a params
func (o *CreateFreeIPAParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create free IP a params
func (o *CreateFreeIPAParams) WithContext(ctx context.Context) *CreateFreeIPAParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create free IP a params
func (o *CreateFreeIPAParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create free IP a params
func (o *CreateFreeIPAParams) WithHTTPClient(client *http.Client) *CreateFreeIPAParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create free IP a params
func (o *CreateFreeIPAParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create free IP a params
func (o *CreateFreeIPAParams) WithBody(body *model.CreateFreeIpaRequest) *CreateFreeIPAParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create free IP a params
func (o *CreateFreeIPAParams) SetBody(body *model.CreateFreeIpaRequest) {
	o.Body = body
}

// WithEnvironmentName adds the environmentName to the create free IP a params
func (o *CreateFreeIPAParams) WithEnvironmentName(environmentName string) *CreateFreeIPAParams {
	o.SetEnvironmentName(environmentName)
	return o
}

// SetEnvironmentName adds the environmentName to the create free IP a params
func (o *CreateFreeIPAParams) SetEnvironmentName(environmentName string) {
	o.EnvironmentName = environmentName
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFreeIPAParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param environmentName
	if err := r.SetPathParam("environmentName", o.EnvironmentName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}