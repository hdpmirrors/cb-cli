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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteInstanceDistroXV1ByNameParams creates a new DeleteInstanceDistroXV1ByNameParams object
// with the default values initialized.
func NewDeleteInstanceDistroXV1ByNameParams() *DeleteInstanceDistroXV1ByNameParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstanceDistroXV1ByNameParams{
		Forced: &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstanceDistroXV1ByNameParamsWithTimeout creates a new DeleteInstanceDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstanceDistroXV1ByNameParamsWithTimeout(timeout time.Duration) *DeleteInstanceDistroXV1ByNameParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstanceDistroXV1ByNameParams{
		Forced: &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteInstanceDistroXV1ByNameParamsWithContext creates a new DeleteInstanceDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstanceDistroXV1ByNameParamsWithContext(ctx context.Context) *DeleteInstanceDistroXV1ByNameParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstanceDistroXV1ByNameParams{
		Forced: &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteInstanceDistroXV1ByNameParamsWithHTTPClient creates a new DeleteInstanceDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstanceDistroXV1ByNameParamsWithHTTPClient(client *http.Client) *DeleteInstanceDistroXV1ByNameParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstanceDistroXV1ByNameParams{
		Forced:     &forcedDefault,
		HTTPClient: client,
	}
}

/*DeleteInstanceDistroXV1ByNameParams contains all the parameters to send to the API endpoint
for the delete instance distro x v1 by name operation typically these are written to a http.Request
*/
type DeleteInstanceDistroXV1ByNameParams struct {

	/*Forced*/
	Forced *bool
	/*InstanceID*/
	InstanceID *string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) WithTimeout(timeout time.Duration) *DeleteInstanceDistroXV1ByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) WithContext(ctx context.Context) *DeleteInstanceDistroXV1ByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) WithHTTPClient(client *http.Client) *DeleteInstanceDistroXV1ByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForced adds the forced to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) WithForced(forced *bool) *DeleteInstanceDistroXV1ByNameParams {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) SetForced(forced *bool) {
	o.Forced = forced
}

// WithInstanceID adds the instanceID to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) WithInstanceID(instanceID *string) *DeleteInstanceDistroXV1ByNameParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) SetInstanceID(instanceID *string) {
	o.InstanceID = instanceID
}

// WithName adds the name to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) WithName(name string) *DeleteInstanceDistroXV1ByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete instance distro x v1 by name params
func (o *DeleteInstanceDistroXV1ByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstanceDistroXV1ByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Forced != nil {

		// query param forced
		var qrForced bool
		if o.Forced != nil {
			qrForced = *o.Forced
		}
		qForced := swag.FormatBool(qrForced)
		if qForced != "" {
			if err := r.SetQueryParam("forced", qForced); err != nil {
				return err
			}
		}

	}

	if o.InstanceID != nil {

		// query param instanceId
		var qrInstanceID string
		if o.InstanceID != nil {
			qrInstanceID = *o.InstanceID
		}
		qInstanceID := qrInstanceID
		if qInstanceID != "" {
			if err := r.SetQueryParam("instanceId", qInstanceID); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
