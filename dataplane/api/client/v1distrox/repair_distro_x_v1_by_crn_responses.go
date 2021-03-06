// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RepairDistroXV1ByCrnReader is a Reader for the RepairDistroXV1ByCrn structure.
type RepairDistroXV1ByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepairDistroXV1ByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewRepairDistroXV1ByCrnDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewRepairDistroXV1ByCrnDefault creates a RepairDistroXV1ByCrnDefault with default headers values
func NewRepairDistroXV1ByCrnDefault(code int) *RepairDistroXV1ByCrnDefault {
	return &RepairDistroXV1ByCrnDefault{
		_statusCode: code,
	}
}

/*RepairDistroXV1ByCrnDefault handles this case with default header values.

successful operation
*/
type RepairDistroXV1ByCrnDefault struct {
	_statusCode int
}

// Code gets the status code for the repair distro x v1 by crn default response
func (o *RepairDistroXV1ByCrnDefault) Code() int {
	return o._statusCode
}

func (o *RepairDistroXV1ByCrnDefault) Error() string {
	return fmt.Sprintf("[POST /v1/distrox/crn/{crn}/manual_repair][%d] repairDistroXV1ByCrn default ", o._statusCode)
}

func (o *RepairDistroXV1ByCrnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
