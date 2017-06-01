package flexsubscriptions

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

// GetPrivateFlexSubscriptionsReader is a Reader for the GetPrivateFlexSubscriptions structure.
type GetPrivateFlexSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPrivateFlexSubscriptionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateFlexSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateFlexSubscriptionsOK creates a GetPrivateFlexSubscriptionsOK with default headers values
func NewGetPrivateFlexSubscriptionsOK() *GetPrivateFlexSubscriptionsOK {
	return &GetPrivateFlexSubscriptionsOK{}
}

/*GetPrivateFlexSubscriptionsOK handles this case with default header values.

successful operation
*/
type GetPrivateFlexSubscriptionsOK struct {
	Payload []*models_cloudbreak.FlexSubscriptionResponse
}

func (o *GetPrivateFlexSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /flexsubscriptions/user][%d] getPrivateFlexSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetPrivateFlexSubscriptionsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}