// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_mpacks

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

// NewListManagementPacksByWorkspaceParams creates a new ListManagementPacksByWorkspaceParams object
// with the default values initialized.
func NewListManagementPacksByWorkspaceParams() *ListManagementPacksByWorkspaceParams {
	var ()
	return &ListManagementPacksByWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListManagementPacksByWorkspaceParamsWithTimeout creates a new ListManagementPacksByWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListManagementPacksByWorkspaceParamsWithTimeout(timeout time.Duration) *ListManagementPacksByWorkspaceParams {
	var ()
	return &ListManagementPacksByWorkspaceParams{

		timeout: timeout,
	}
}

// NewListManagementPacksByWorkspaceParamsWithContext creates a new ListManagementPacksByWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewListManagementPacksByWorkspaceParamsWithContext(ctx context.Context) *ListManagementPacksByWorkspaceParams {
	var ()
	return &ListManagementPacksByWorkspaceParams{

		Context: ctx,
	}
}

// NewListManagementPacksByWorkspaceParamsWithHTTPClient creates a new ListManagementPacksByWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListManagementPacksByWorkspaceParamsWithHTTPClient(client *http.Client) *ListManagementPacksByWorkspaceParams {
	var ()
	return &ListManagementPacksByWorkspaceParams{
		HTTPClient: client,
	}
}

/*ListManagementPacksByWorkspaceParams contains all the parameters to send to the API endpoint
for the list management packs by workspace operation typically these are written to a http.Request
*/
type ListManagementPacksByWorkspaceParams struct {

	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list management packs by workspace params
func (o *ListManagementPacksByWorkspaceParams) WithTimeout(timeout time.Duration) *ListManagementPacksByWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list management packs by workspace params
func (o *ListManagementPacksByWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list management packs by workspace params
func (o *ListManagementPacksByWorkspaceParams) WithContext(ctx context.Context) *ListManagementPacksByWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list management packs by workspace params
func (o *ListManagementPacksByWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list management packs by workspace params
func (o *ListManagementPacksByWorkspaceParams) WithHTTPClient(client *http.Client) *ListManagementPacksByWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list management packs by workspace params
func (o *ListManagementPacksByWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceID adds the workspaceID to the list management packs by workspace params
func (o *ListManagementPacksByWorkspaceParams) WithWorkspaceID(workspaceID int64) *ListManagementPacksByWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the list management packs by workspace params
func (o *ListManagementPacksByWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListManagementPacksByWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
