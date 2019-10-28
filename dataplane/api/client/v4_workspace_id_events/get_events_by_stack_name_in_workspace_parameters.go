// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_events

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

// NewGetEventsByStackNameInWorkspaceParams creates a new GetEventsByStackNameInWorkspaceParams object
// with the default values initialized.
func NewGetEventsByStackNameInWorkspaceParams() *GetEventsByStackNameInWorkspaceParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetEventsByStackNameInWorkspaceParams{
		Page: &pageDefault,
		Size: &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventsByStackNameInWorkspaceParamsWithTimeout creates a new GetEventsByStackNameInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventsByStackNameInWorkspaceParamsWithTimeout(timeout time.Duration) *GetEventsByStackNameInWorkspaceParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetEventsByStackNameInWorkspaceParams{
		Page: &pageDefault,
		Size: &sizeDefault,

		timeout: timeout,
	}
}

// NewGetEventsByStackNameInWorkspaceParamsWithContext creates a new GetEventsByStackNameInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventsByStackNameInWorkspaceParamsWithContext(ctx context.Context) *GetEventsByStackNameInWorkspaceParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetEventsByStackNameInWorkspaceParams{
		Page: &pageDefault,
		Size: &sizeDefault,

		Context: ctx,
	}
}

// NewGetEventsByStackNameInWorkspaceParamsWithHTTPClient creates a new GetEventsByStackNameInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventsByStackNameInWorkspaceParamsWithHTTPClient(client *http.Client) *GetEventsByStackNameInWorkspaceParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetEventsByStackNameInWorkspaceParams{
		Page:       &pageDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*GetEventsByStackNameInWorkspaceParams contains all the parameters to send to the API endpoint
for the get events by stack name in workspace operation typically these are written to a http.Request
*/
type GetEventsByStackNameInWorkspaceParams struct {

	/*Name*/
	Name string
	/*Page*/
	Page *int32
	/*Size*/
	Size *int32
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) WithTimeout(timeout time.Duration) *GetEventsByStackNameInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) WithContext(ctx context.Context) *GetEventsByStackNameInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) WithHTTPClient(client *http.Client) *GetEventsByStackNameInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) WithName(name string) *GetEventsByStackNameInWorkspaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) SetName(name string) {
	o.Name = name
}

// WithPage adds the page to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) WithPage(page *int32) *GetEventsByStackNameInWorkspaceParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) SetPage(page *int32) {
	o.Page = page
}

// WithSize adds the size to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) WithSize(size *int32) *GetEventsByStackNameInWorkspaceParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) SetSize(size *int32) {
	o.Size = size
}

// WithWorkspaceID adds the workspaceID to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetEventsByStackNameInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get events by stack name in workspace params
func (o *GetEventsByStackNameInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventsByStackNameInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

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