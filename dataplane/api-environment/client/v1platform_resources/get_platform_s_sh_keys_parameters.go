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

// NewGetPlatformSShKeysParams creates a new GetPlatformSShKeysParams object
// with the default values initialized.
func NewGetPlatformSShKeysParams() *GetPlatformSShKeysParams {
	var ()
	return &GetPlatformSShKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlatformSShKeysParamsWithTimeout creates a new GetPlatformSShKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlatformSShKeysParamsWithTimeout(timeout time.Duration) *GetPlatformSShKeysParams {
	var ()
	return &GetPlatformSShKeysParams{

		timeout: timeout,
	}
}

// NewGetPlatformSShKeysParamsWithContext creates a new GetPlatformSShKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlatformSShKeysParamsWithContext(ctx context.Context) *GetPlatformSShKeysParams {
	var ()
	return &GetPlatformSShKeysParams{

		Context: ctx,
	}
}

// NewGetPlatformSShKeysParamsWithHTTPClient creates a new GetPlatformSShKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlatformSShKeysParamsWithHTTPClient(client *http.Client) *GetPlatformSShKeysParams {
	var ()
	return &GetPlatformSShKeysParams{
		HTTPClient: client,
	}
}

/*GetPlatformSShKeysParams contains all the parameters to send to the API endpoint
for the get platform s sh keys operation typically these are written to a http.Request
*/
type GetPlatformSShKeysParams struct {

	/*AvailabilityZone*/
	AvailabilityZone *string
	/*CredentialCrn*/
	CredentialCrn *string
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

// WithTimeout adds the timeout to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) WithTimeout(timeout time.Duration) *GetPlatformSShKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) WithContext(ctx context.Context) *GetPlatformSShKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) WithHTTPClient(client *http.Client) *GetPlatformSShKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) WithAvailabilityZone(availabilityZone *string) *GetPlatformSShKeysParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithCredentialCrn adds the credentialCrn to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) WithCredentialCrn(credentialCrn *string) *GetPlatformSShKeysParams {
	o.SetCredentialCrn(credentialCrn)
	return o
}

// SetCredentialCrn adds the credentialCrn to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) SetCredentialCrn(credentialCrn *string) {
	o.CredentialCrn = credentialCrn
}

// WithCredentialName adds the credentialName to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) WithCredentialName(credentialName *string) *GetPlatformSShKeysParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithPlatformVariant adds the platformVariant to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) WithPlatformVariant(platformVariant *string) *GetPlatformSShKeysParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) WithRegion(region *string) *GetPlatformSShKeysParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get platform s sh keys params
func (o *GetPlatformSShKeysParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlatformSShKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CredentialCrn != nil {

		// query param credentialCrn
		var qrCredentialCrn string
		if o.CredentialCrn != nil {
			qrCredentialCrn = *o.CredentialCrn
		}
		qCredentialCrn := qrCredentialCrn
		if qCredentialCrn != "" {
			if err := r.SetQueryParam("credentialCrn", qCredentialCrn); err != nil {
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