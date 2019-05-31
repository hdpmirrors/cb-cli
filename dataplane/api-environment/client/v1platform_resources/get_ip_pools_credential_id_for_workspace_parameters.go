// Code generated by go-swagger; DO NOT EDIT.

package v1platform_resources

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

// NewGetIPPoolsCredentialIDForWorkspaceParams creates a new GetIPPoolsCredentialIDForWorkspaceParams object
// with the default values initialized.
func NewGetIPPoolsCredentialIDForWorkspaceParams() *GetIPPoolsCredentialIDForWorkspaceParams {
	var ()
	return &GetIPPoolsCredentialIDForWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPPoolsCredentialIDForWorkspaceParamsWithTimeout creates a new GetIPPoolsCredentialIDForWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPPoolsCredentialIDForWorkspaceParamsWithTimeout(timeout time.Duration) *GetIPPoolsCredentialIDForWorkspaceParams {
	var ()
	return &GetIPPoolsCredentialIDForWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetIPPoolsCredentialIDForWorkspaceParamsWithContext creates a new GetIPPoolsCredentialIDForWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPPoolsCredentialIDForWorkspaceParamsWithContext(ctx context.Context) *GetIPPoolsCredentialIDForWorkspaceParams {
	var ()
	return &GetIPPoolsCredentialIDForWorkspaceParams{

		Context: ctx,
	}
}

// NewGetIPPoolsCredentialIDForWorkspaceParamsWithHTTPClient creates a new GetIPPoolsCredentialIDForWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPPoolsCredentialIDForWorkspaceParamsWithHTTPClient(client *http.Client) *GetIPPoolsCredentialIDForWorkspaceParams {
	var ()
	return &GetIPPoolsCredentialIDForWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetIPPoolsCredentialIDForWorkspaceParams contains all the parameters to send to the API endpoint
for the get Ip pools credential Id for workspace operation typically these are written to a http.Request
*/
type GetIPPoolsCredentialIDForWorkspaceParams struct {

	/*AvailabilityZone*/
	AvailabilityZone *string
	/*CredentialName*/
	CredentialName *string
	/*PlatformVariant*/
	PlatformVariant *string
	/*Region*/
	Region *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithTimeout(timeout time.Duration) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithContext(ctx context.Context) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithHTTPClient(client *http.Client) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithAvailabilityZone(availabilityZone *string) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithCredentialName adds the credentialName to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithCredentialName(credentialName *string) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithPlatformVariant adds the platformVariant to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithPlatformVariant(platformVariant *string) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithRegion(region *string) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AvailabilityZone != nil {

		// query param availabilityZone
		var qrAvailabilityZone string
		if o.AvailabilityZone != nil {
			qrAvailabilityZone = *o.AvailabilityZone
		}
		qAvailabilityZone := qrAvailabilityZone
		if qAvailabilityZone != "" {
			if err := r.SetQueryParam("availabilityZone", qAvailabilityZone); err != nil {
				return err
			}
		}

	}

	if o.CredentialName != nil {

		// query param credentialName
		var qrCredentialName string
		if o.CredentialName != nil {
			qrCredentialName = *o.CredentialName
		}
		qCredentialName := qrCredentialName
		if qCredentialName != "" {
			if err := r.SetQueryParam("credentialName", qCredentialName); err != nil {
				return err
			}
		}

	}

	if o.PlatformVariant != nil {

		// query param platformVariant
		var qrPlatformVariant string
		if o.PlatformVariant != nil {
			qrPlatformVariant = *o.PlatformVariant
		}
		qPlatformVariant := qrPlatformVariant
		if qPlatformVariant != "" {
			if err := r.SetQueryParam("platformVariant", qPlatformVariant); err != nil {
				return err
			}
		}

	}

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}