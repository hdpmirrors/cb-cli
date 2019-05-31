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

// NewGetPlatformSShKeysForWorkspaceParams creates a new GetPlatformSShKeysForWorkspaceParams object
// with the default values initialized.
func NewGetPlatformSShKeysForWorkspaceParams() *GetPlatformSShKeysForWorkspaceParams {
	var ()
	return &GetPlatformSShKeysForWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlatformSShKeysForWorkspaceParamsWithTimeout creates a new GetPlatformSShKeysForWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlatformSShKeysForWorkspaceParamsWithTimeout(timeout time.Duration) *GetPlatformSShKeysForWorkspaceParams {
	var ()
	return &GetPlatformSShKeysForWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetPlatformSShKeysForWorkspaceParamsWithContext creates a new GetPlatformSShKeysForWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlatformSShKeysForWorkspaceParamsWithContext(ctx context.Context) *GetPlatformSShKeysForWorkspaceParams {
	var ()
	return &GetPlatformSShKeysForWorkspaceParams{

		Context: ctx,
	}
}

// NewGetPlatformSShKeysForWorkspaceParamsWithHTTPClient creates a new GetPlatformSShKeysForWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlatformSShKeysForWorkspaceParamsWithHTTPClient(client *http.Client) *GetPlatformSShKeysForWorkspaceParams {
	var ()
	return &GetPlatformSShKeysForWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetPlatformSShKeysForWorkspaceParams contains all the parameters to send to the API endpoint
for the get platform s sh keys for workspace operation typically these are written to a http.Request
*/
type GetPlatformSShKeysForWorkspaceParams struct {

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

// WithTimeout adds the timeout to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) WithTimeout(timeout time.Duration) *GetPlatformSShKeysForWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) WithContext(ctx context.Context) *GetPlatformSShKeysForWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) WithHTTPClient(client *http.Client) *GetPlatformSShKeysForWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) WithAvailabilityZone(availabilityZone *string) *GetPlatformSShKeysForWorkspaceParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithCredentialName adds the credentialName to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) WithCredentialName(credentialName *string) *GetPlatformSShKeysForWorkspaceParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithPlatformVariant adds the platformVariant to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) WithPlatformVariant(platformVariant *string) *GetPlatformSShKeysForWorkspaceParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) WithRegion(region *string) *GetPlatformSShKeysForWorkspaceParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get platform s sh keys for workspace params
func (o *GetPlatformSShKeysForWorkspaceParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlatformSShKeysForWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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