// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpgradeOsInWorkspaceV4InternalParams creates a new UpgradeOsInWorkspaceV4InternalParams object
// with the default values initialized.
func NewUpgradeOsInWorkspaceV4InternalParams() *UpgradeOsInWorkspaceV4InternalParams {
	var ()
	return &UpgradeOsInWorkspaceV4InternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeOsInWorkspaceV4InternalParamsWithTimeout creates a new UpgradeOsInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeOsInWorkspaceV4InternalParamsWithTimeout(timeout time.Duration) *UpgradeOsInWorkspaceV4InternalParams {
	var ()
	return &UpgradeOsInWorkspaceV4InternalParams{

		timeout: timeout,
	}
}

// NewUpgradeOsInWorkspaceV4InternalParamsWithContext creates a new UpgradeOsInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeOsInWorkspaceV4InternalParamsWithContext(ctx context.Context) *UpgradeOsInWorkspaceV4InternalParams {
	var ()
	return &UpgradeOsInWorkspaceV4InternalParams{

		Context: ctx,
	}
}

// NewUpgradeOsInWorkspaceV4InternalParamsWithHTTPClient creates a new UpgradeOsInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeOsInWorkspaceV4InternalParamsWithHTTPClient(client *http.Client) *UpgradeOsInWorkspaceV4InternalParams {
	var ()
	return &UpgradeOsInWorkspaceV4InternalParams{
		HTTPClient: client,
	}
}

/*UpgradeOsInWorkspaceV4InternalParams contains all the parameters to send to the API endpoint
for the upgrade os in workspace v4 internal operation typically these are written to a http.Request
*/
type UpgradeOsInWorkspaceV4InternalParams struct {

	/*InitiatorUserCrn*/
	InitiatorUserCrn *string
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) WithTimeout(timeout time.Duration) *UpgradeOsInWorkspaceV4InternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) WithContext(ctx context.Context) *UpgradeOsInWorkspaceV4InternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) WithHTTPClient(client *http.Client) *UpgradeOsInWorkspaceV4InternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *UpgradeOsInWorkspaceV4InternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithName adds the name to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) WithName(name string) *UpgradeOsInWorkspaceV4InternalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) WithWorkspaceID(workspaceID int64) *UpgradeOsInWorkspaceV4InternalParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the upgrade os in workspace v4 internal params
func (o *UpgradeOsInWorkspaceV4InternalParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeOsInWorkspaceV4InternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InitiatorUserCrn != nil {

		// query param initiatorUserCrn
		var qrInitiatorUserCrn string
		if o.InitiatorUserCrn != nil {
			qrInitiatorUserCrn = *o.InitiatorUserCrn
		}
		qInitiatorUserCrn := qrInitiatorUserCrn
		if qInitiatorUserCrn != "" {
			if err := r.SetQueryParam("initiatorUserCrn", qInitiatorUserCrn); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}