package util

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// TestRdsConnectionByIDUtilReader is a Reader for the TestRdsConnectionByIDUtil structure.
type TestRdsConnectionByIDUtilReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *TestRdsConnectionByIDUtilReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTestRdsConnectionByIDUtilOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestRdsConnectionByIDUtilOK creates a TestRdsConnectionByIDUtilOK with default headers values
func NewTestRdsConnectionByIDUtilOK() *TestRdsConnectionByIDUtilOK {
	return &TestRdsConnectionByIDUtilOK{}
}

/*TestRdsConnectionByIDUtilOK handles this case with default header values.

successful operation
*/
type TestRdsConnectionByIDUtilOK struct {
	Payload *models_cloudbreak.RdsTestResult
}

func (o *TestRdsConnectionByIDUtilOK) Error() string {
	return fmt.Sprintf("[GET /util/rds/{id}][%d] testRdsConnectionByIdUtilOK  %+v", 200, o.Payload)
}

func (o *TestRdsConnectionByIDUtilOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.RdsTestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}