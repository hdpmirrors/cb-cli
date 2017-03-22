package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewDeletePrivateNetworkParams creates a new DeletePrivateNetworkParams object
// with the default values initialized.
func NewDeletePrivateNetworkParams() *DeletePrivateNetworkParams {
	var ()
	return &DeletePrivateNetworkParams{}
}

/*DeletePrivateNetworkParams contains all the parameters to send to the API endpoint
for the delete private network operation typically these are written to a http.Request
*/
type DeletePrivateNetworkParams struct {

	/*Name*/
	Name string
}

// WithName adds the name to the delete private network params
func (o *DeletePrivateNetworkParams) WithName(name string) *DeletePrivateNetworkParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePrivateNetworkParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}