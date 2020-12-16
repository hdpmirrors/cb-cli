// Code generated by go-swagger; DO NOT EDIT.

package v1credentialsaudit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// GetAuditCredentialByResourceNameV1Reader is a Reader for the GetAuditCredentialByResourceNameV1 structure.
type GetAuditCredentialByResourceNameV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditCredentialByResourceNameV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuditCredentialByResourceNameV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAuditCredentialByResourceNameV1OK creates a GetAuditCredentialByResourceNameV1OK with default headers values
func NewGetAuditCredentialByResourceNameV1OK() *GetAuditCredentialByResourceNameV1OK {
	return &GetAuditCredentialByResourceNameV1OK{}
}

/*GetAuditCredentialByResourceNameV1OK handles this case with default header values.

successful operation
*/
type GetAuditCredentialByResourceNameV1OK struct {
	Payload *model.CredentialV1Response
}

func (o *GetAuditCredentialByResourceNameV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/audit/name/{name}][%d] getAuditCredentialByResourceNameV1OK  %+v", 200, o.Payload)
}

func (o *GetAuditCredentialByResourceNameV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CredentialV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}