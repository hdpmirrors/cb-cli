// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_imagecatalogs

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

// NewSetDefaultImageCatalogByNameInWorkspaceParams creates a new SetDefaultImageCatalogByNameInWorkspaceParams object
// with the default values initialized.
func NewSetDefaultImageCatalogByNameInWorkspaceParams() *SetDefaultImageCatalogByNameInWorkspaceParams {
	var ()
	return &SetDefaultImageCatalogByNameInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetDefaultImageCatalogByNameInWorkspaceParamsWithTimeout creates a new SetDefaultImageCatalogByNameInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetDefaultImageCatalogByNameInWorkspaceParamsWithTimeout(timeout time.Duration) *SetDefaultImageCatalogByNameInWorkspaceParams {
	var ()
	return &SetDefaultImageCatalogByNameInWorkspaceParams{

		timeout: timeout,
	}
}

// NewSetDefaultImageCatalogByNameInWorkspaceParamsWithContext creates a new SetDefaultImageCatalogByNameInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetDefaultImageCatalogByNameInWorkspaceParamsWithContext(ctx context.Context) *SetDefaultImageCatalogByNameInWorkspaceParams {
	var ()
	return &SetDefaultImageCatalogByNameInWorkspaceParams{

		Context: ctx,
	}
}

// NewSetDefaultImageCatalogByNameInWorkspaceParamsWithHTTPClient creates a new SetDefaultImageCatalogByNameInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetDefaultImageCatalogByNameInWorkspaceParamsWithHTTPClient(client *http.Client) *SetDefaultImageCatalogByNameInWorkspaceParams {
	var ()
	return &SetDefaultImageCatalogByNameInWorkspaceParams{
		HTTPClient: client,
	}
}

/*SetDefaultImageCatalogByNameInWorkspaceParams contains all the parameters to send to the API endpoint
for the set default image catalog by name in workspace operation typically these are written to a http.Request
*/
type SetDefaultImageCatalogByNameInWorkspaceParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set default image catalog by name in workspace params
func (o *SetDefaultImageCatalogByNameInWorkspaceParams) WithTimeout(timeout time.Duration) *SetDefaultImageCatalogByNameInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set default image catalog by name in workspace params
func (o *SetDefaultImageCatalogByNameInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set default image catalog by name in workspace params
func (o *SetDefaultImageCatalogByNameInWorkspaceParams) WithContext(ctx context.Context) *SetDefaultImageCatalogByNameInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set default image catalog by name in workspace params
func (o *SetDefaultImageCatalogByNameInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set default image catalog by name in workspace params
func (o *SetDefaultImageCatalogByNameInWorkspaceParams) WithHTTPClient(client *http.Client) *SetDefaultImageCatalogByNameInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set default image catalog by name in workspace params
func (o *SetDefaultImageCatalogByNameInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the set default image catalog by name in workspace params
func (o *SetDefaultImageCatalogByNameInWorkspaceParams) WithName(name string) *SetDefaultImageCatalogByNameInWorkspaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the set default image catalog by name in workspace params
func (o *SetDefaultImageCatalogByNameInWorkspaceParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the set default image catalog by name in workspace params
func (o *SetDefaultImageCatalogByNameInWorkspaceParams) WithWorkspaceID(workspaceID int64) *SetDefaultImageCatalogByNameInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the set default image catalog by name in workspace params
func (o *SetDefaultImageCatalogByNameInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *SetDefaultImageCatalogByNameInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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