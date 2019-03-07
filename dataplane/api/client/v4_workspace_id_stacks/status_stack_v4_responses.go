// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// StatusStackV4Reader is a Reader for the StatusStackV4 structure.
type StatusStackV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusStackV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStatusStackV4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStatusStackV4OK creates a StatusStackV4OK with default headers values
func NewStatusStackV4OK() *StatusStackV4OK {
	return &StatusStackV4OK{}
}

/*StatusStackV4OK handles this case with default header values.

successful operation
*/
type StatusStackV4OK struct {
	Payload *model.StackStatusV4Response
}

func (o *StatusStackV4OK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/stacks/{name}/status][%d] statusStackV4OK  %+v", 200, o.Payload)
}

func (o *StatusStackV4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.StackStatusV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}