// Code generated by go-swagger; DO NOT EDIT.

package v1kerberosmgmt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// GenerateServiceKeytabV1Reader is a Reader for the GenerateServiceKeytabV1 structure.
type GenerateServiceKeytabV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateServiceKeytabV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGenerateServiceKeytabV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGenerateServiceKeytabV1OK creates a GenerateServiceKeytabV1OK with default headers values
func NewGenerateServiceKeytabV1OK() *GenerateServiceKeytabV1OK {
	return &GenerateServiceKeytabV1OK{}
}

/*GenerateServiceKeytabV1OK handles this case with default header values.

successful operation
*/
type GenerateServiceKeytabV1OK struct {
	Payload *model.ServiceKeytabV1Response
}

func (o *GenerateServiceKeytabV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/kerberosmgmt/servicekeytab][%d] generateServiceKeytabV1OK  %+v", 200, o.Payload)
}

func (o *GenerateServiceKeytabV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ServiceKeytabV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}