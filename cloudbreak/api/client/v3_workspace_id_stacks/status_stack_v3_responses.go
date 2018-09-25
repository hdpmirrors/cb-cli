// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// StatusStackV3Reader is a Reader for the StatusStackV3 structure.
type StatusStackV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusStackV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStatusStackV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStatusStackV3OK creates a StatusStackV3OK with default headers values
func NewStatusStackV3OK() *StatusStackV3OK {
	return &StatusStackV3OK{}
}

/*StatusStackV3OK handles this case with default header values.

successful operation
*/
type StatusStackV3OK struct {
	Payload StatusStackV3OKBody
}

func (o *StatusStackV3OK) Error() string {
	return fmt.Sprintf("[GET /v3/{workspaceId}/stacks/{name}/status][%d] statusStackV3OK  %+v", 200, o.Payload)
}

func (o *StatusStackV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StatusStackV3OKBody status stack v3 o k body
swagger:model StatusStackV3OKBody
*/

type StatusStackV3OKBody map[string]interface{}

// Validate validates this status stack v3 o k body
func (o StatusStackV3OKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("statusStackV3OK", "body", o); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}