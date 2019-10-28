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

// ListRetryableFlowsV4Reader is a Reader for the ListRetryableFlowsV4 structure.
type ListRetryableFlowsV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRetryableFlowsV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRetryableFlowsV4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRetryableFlowsV4OK creates a ListRetryableFlowsV4OK with default headers values
func NewListRetryableFlowsV4OK() *ListRetryableFlowsV4OK {
	return &ListRetryableFlowsV4OK{}
}

/*ListRetryableFlowsV4OK handles this case with default header values.

successful operation
*/
type ListRetryableFlowsV4OK struct {
	Payload []*model.RetryableFlowResponse
}

func (o *ListRetryableFlowsV4OK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/stacks/{name}/retry][%d] listRetryableFlowsV4OK  %+v", 200, o.Payload)
}

func (o *ListRetryableFlowsV4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}