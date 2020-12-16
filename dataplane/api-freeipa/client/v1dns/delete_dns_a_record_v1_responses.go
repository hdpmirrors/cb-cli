// Code generated by go-swagger; DO NOT EDIT.

package v1dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteDNSARecordV1Reader is a Reader for the DeleteDNSARecordV1 structure.
type DeleteDNSARecordV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDNSARecordV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteDNSARecordV1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteDNSARecordV1Default creates a DeleteDNSARecordV1Default with default headers values
func NewDeleteDNSARecordV1Default(code int) *DeleteDNSARecordV1Default {
	return &DeleteDNSARecordV1Default{
		_statusCode: code,
	}
}

/*DeleteDNSARecordV1Default handles this case with default header values.

successful operation
*/
type DeleteDNSARecordV1Default struct {
	_statusCode int
}

// Code gets the status code for the delete Dns a record v1 default response
func (o *DeleteDNSARecordV1Default) Code() int {
	return o._statusCode
}

func (o *DeleteDNSARecordV1Default) Error() string {
	return fmt.Sprintf("[DELETE /v1/dns/record/a][%d] deleteDnsARecordV1 default ", o._statusCode)
}

func (o *DeleteDNSARecordV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}