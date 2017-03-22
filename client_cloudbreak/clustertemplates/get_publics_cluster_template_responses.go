package clustertemplates

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

// GetPublicsClusterTemplateReader is a Reader for the GetPublicsClusterTemplate structure.
type GetPublicsClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPublicsClusterTemplateReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsClusterTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsClusterTemplateOK creates a GetPublicsClusterTemplateOK with default headers values
func NewGetPublicsClusterTemplateOK() *GetPublicsClusterTemplateOK {
	return &GetPublicsClusterTemplateOK{}
}

/*GetPublicsClusterTemplateOK handles this case with default header values.

successful operation
*/
type GetPublicsClusterTemplateOK struct {
	Payload []*models_cloudbreak.ClusterTemplateResponse
}

func (o *GetPublicsClusterTemplateOK) Error() string {
	return fmt.Sprintf("[GET /clustertemplates/account][%d] getPublicsClusterTemplateOK  %+v", 200, o.Payload)
}

func (o *GetPublicsClusterTemplateOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}