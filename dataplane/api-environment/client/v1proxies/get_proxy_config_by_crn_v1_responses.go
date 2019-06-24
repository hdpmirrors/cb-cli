// Code generated by go-swagger; DO NOT EDIT.

package v1proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// GetProxyConfigByCrnV1Reader is a Reader for the GetProxyConfigByCrnV1 structure.
type GetProxyConfigByCrnV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProxyConfigByCrnV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProxyConfigByCrnV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProxyConfigByCrnV1OK creates a GetProxyConfigByCrnV1OK with default headers values
func NewGetProxyConfigByCrnV1OK() *GetProxyConfigByCrnV1OK {
	return &GetProxyConfigByCrnV1OK{}
}

/*GetProxyConfigByCrnV1OK handles this case with default header values.

successful operation
*/
type GetProxyConfigByCrnV1OK struct {
	Payload *model.ProxyResponse
}

func (o *GetProxyConfigByCrnV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/proxies/crn/{crn}][%d] getProxyConfigByCrnV1OK  %+v", 200, o.Payload)
}

func (o *GetProxyConfigByCrnV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ProxyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}