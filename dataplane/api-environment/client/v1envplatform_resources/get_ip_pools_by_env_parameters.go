// Code generated by go-swagger; DO NOT EDIT.

package v1envplatform_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetIPPoolsByEnvParams creates a new GetIPPoolsByEnvParams object
// with the default values initialized.
func NewGetIPPoolsByEnvParams() *GetIPPoolsByEnvParams {
	var ()
	return &GetIPPoolsByEnvParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPPoolsByEnvParamsWithTimeout creates a new GetIPPoolsByEnvParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPPoolsByEnvParamsWithTimeout(timeout time.Duration) *GetIPPoolsByEnvParams {
	var ()
	return &GetIPPoolsByEnvParams{

		timeout: timeout,
	}
}

// NewGetIPPoolsByEnvParamsWithContext creates a new GetIPPoolsByEnvParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPPoolsByEnvParamsWithContext(ctx context.Context) *GetIPPoolsByEnvParams {
	var ()
	return &GetIPPoolsByEnvParams{

		Context: ctx,
	}
}

// NewGetIPPoolsByEnvParamsWithHTTPClient creates a new GetIPPoolsByEnvParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPPoolsByEnvParamsWithHTTPClient(client *http.Client) *GetIPPoolsByEnvParams {
	var ()
	return &GetIPPoolsByEnvParams{
		HTTPClient: client,
	}
}

/*GetIPPoolsByEnvParams contains all the parameters to send to the API endpoint
for the get Ip pools by env operation typically these are written to a http.Request
*/
type GetIPPoolsByEnvParams struct {

	/*AvailabilityZone*/
	AvailabilityZone *string
	/*EnvironmentCrn*/
	EnvironmentCrn *string
	/*PlatformVariant*/
	PlatformVariant *string
	/*Region*/
	Region *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) WithTimeout(timeout time.Duration) *GetIPPoolsByEnvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) WithContext(ctx context.Context) *GetIPPoolsByEnvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) WithHTTPClient(client *http.Client) *GetIPPoolsByEnvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) WithAvailabilityZone(availabilityZone *string) *GetIPPoolsByEnvParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithEnvironmentCrn adds the environmentCrn to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) WithEnvironmentCrn(environmentCrn *string) *GetIPPoolsByEnvParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) SetEnvironmentCrn(environmentCrn *string) {
	o.EnvironmentCrn = environmentCrn
}

// WithPlatformVariant adds the platformVariant to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) WithPlatformVariant(platformVariant *string) *GetIPPoolsByEnvParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) WithRegion(region *string) *GetIPPoolsByEnvParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get Ip pools by env params
func (o *GetIPPoolsByEnvParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPPoolsByEnvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EnvironmentCrn != nil {

		// query param environmentCrn
		var qrEnvironmentCrn string
		if o.EnvironmentCrn != nil {
			qrEnvironmentCrn = *o.EnvironmentCrn
		}
		qEnvironmentCrn := qrEnvironmentCrn
		if qEnvironmentCrn != "" {
			if err := r.SetQueryParam("environmentCrn", qEnvironmentCrn); err != nil {
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