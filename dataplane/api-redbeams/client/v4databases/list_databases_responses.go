// Code generated by go-swagger; DO NOT EDIT.

package v4databases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
)

// ListDatabasesReader is a Reader for the ListDatabases structure.
type ListDatabasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDatabasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListDatabasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListDatabasesOK creates a ListDatabasesOK with default headers values
func NewListDatabasesOK() *ListDatabasesOK {
	return &ListDatabasesOK{}
}

/*ListDatabasesOK handles this case with default header values.

successful operation
*/
type ListDatabasesOK struct {
	Payload *model.DatabaseV4Responses
}

func (o *ListDatabasesOK) Error() string {
	return fmt.Sprintf("[GET /v4/databases][%d] listDatabasesOK  %+v", 200, o.Payload)
}

func (o *ListDatabasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DatabaseV4Responses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}