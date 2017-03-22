package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewDeletePrivateStackParams creates a new DeletePrivateStackParams object
// with the default values initialized.
func NewDeletePrivateStackParams() *DeletePrivateStackParams {
	var (
		deleteDependenciesDefault bool = bool(false)
		forcedDefault             bool = bool(false)
	)
	return &DeletePrivateStackParams{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,
	}
}

/*DeletePrivateStackParams contains all the parameters to send to the API endpoint
for the delete private stack operation typically these are written to a http.Request
*/
type DeletePrivateStackParams struct {

	/*DeleteDependencies*/
	DeleteDependencies *bool
	/*Forced*/
	Forced *bool
	/*Name*/
	Name string
}

// WithDeleteDependencies adds the deleteDependencies to the delete private stack params
func (o *DeletePrivateStackParams) WithDeleteDependencies(deleteDependencies *bool) *DeletePrivateStackParams {
	o.DeleteDependencies = deleteDependencies
	return o
}

// WithForced adds the forced to the delete private stack params
func (o *DeletePrivateStackParams) WithForced(forced *bool) *DeletePrivateStackParams {
	o.Forced = forced
	return o
}

// WithName adds the name to the delete private stack params
func (o *DeletePrivateStackParams) WithName(name string) *DeletePrivateStackParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePrivateStackParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.DeleteDependencies != nil {

		// query param deleteDependencies
		var qrDeleteDependencies bool
		if o.DeleteDependencies != nil {
			qrDeleteDependencies = *o.DeleteDependencies
		}
		qDeleteDependencies := swag.FormatBool(qrDeleteDependencies)
		if qDeleteDependencies != "" {
			if err := r.SetQueryParam("deleteDependencies", qDeleteDependencies); err != nil {
				return err
			}
		}

	}

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}