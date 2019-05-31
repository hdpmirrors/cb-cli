// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

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

// NewGetFreeIpaByEnvironmentV1Params creates a new GetFreeIpaByEnvironmentV1Params object
// with the default values initialized.
func NewGetFreeIpaByEnvironmentV1Params() *GetFreeIpaByEnvironmentV1Params {
	var ()
	return &GetFreeIpaByEnvironmentV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFreeIpaByEnvironmentV1ParamsWithTimeout creates a new GetFreeIpaByEnvironmentV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFreeIpaByEnvironmentV1ParamsWithTimeout(timeout time.Duration) *GetFreeIpaByEnvironmentV1Params {
	var ()
	return &GetFreeIpaByEnvironmentV1Params{

		timeout: timeout,
	}
}

// NewGetFreeIpaByEnvironmentV1ParamsWithContext creates a new GetFreeIpaByEnvironmentV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetFreeIpaByEnvironmentV1ParamsWithContext(ctx context.Context) *GetFreeIpaByEnvironmentV1Params {
	var ()
	return &GetFreeIpaByEnvironmentV1Params{

		Context: ctx,
	}
}

// NewGetFreeIpaByEnvironmentV1ParamsWithHTTPClient creates a new GetFreeIpaByEnvironmentV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFreeIpaByEnvironmentV1ParamsWithHTTPClient(client *http.Client) *GetFreeIpaByEnvironmentV1Params {
	var ()
	return &GetFreeIpaByEnvironmentV1Params{
		HTTPClient: client,
	}
}

/*GetFreeIpaByEnvironmentV1Params contains all the parameters to send to the API endpoint
for the get free ipa by environment v1 operation typically these are written to a http.Request
*/
type GetFreeIpaByEnvironmentV1Params struct {

	/*Environment*/
	Environment *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get free ipa by environment v1 params
func (o *GetFreeIpaByEnvironmentV1Params) WithTimeout(timeout time.Duration) *GetFreeIpaByEnvironmentV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get free ipa by environment v1 params
func (o *GetFreeIpaByEnvironmentV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get free ipa by environment v1 params
func (o *GetFreeIpaByEnvironmentV1Params) WithContext(ctx context.Context) *GetFreeIpaByEnvironmentV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get free ipa by environment v1 params
func (o *GetFreeIpaByEnvironmentV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get free ipa by environment v1 params
func (o *GetFreeIpaByEnvironmentV1Params) WithHTTPClient(client *http.Client) *GetFreeIpaByEnvironmentV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get free ipa by environment v1 params
func (o *GetFreeIpaByEnvironmentV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the get free ipa by environment v1 params
func (o *GetFreeIpaByEnvironmentV1Params) WithEnvironment(environment *string) *GetFreeIpaByEnvironmentV1Params {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the get free ipa by environment v1 params
func (o *GetFreeIpaByEnvironmentV1Params) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WriteToRequest writes these params to a swagger request
func (o *GetFreeIpaByEnvironmentV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Environment != nil {

		// query param environment
		var qrEnvironment string
		if o.Environment != nil {
			qrEnvironment = *o.Environment
		}
		qEnvironment := qrEnvironment
		if qEnvironment != "" {
			if err := r.SetQueryParam("environment", qEnvironment); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}