package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

// GetDriverEndpointsReader is a Reader for the GetDriverEndpoints structure.
type GetDriverEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetDriverEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDriverEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetDriverEndpointsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDriverEndpointsOK creates a GetDriverEndpointsOK with default headers values
func NewGetDriverEndpointsOK() *GetDriverEndpointsOK {
	return &GetDriverEndpointsOK{}
}

/*GetDriverEndpointsOK handles this case with default header values.

OK
*/
type GetDriverEndpointsOK struct {
	Payload []*models.DriverEndpoint
}

func (o *GetDriverEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /driver_endpoints][%d] getDriverEndpointsOK  %+v", 200, o.Payload)
}

func (o *GetDriverEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverEndpointsInternalServerError creates a GetDriverEndpointsInternalServerError with default headers values
func NewGetDriverEndpointsInternalServerError() *GetDriverEndpointsInternalServerError {
	return &GetDriverEndpointsInternalServerError{}
}

/*GetDriverEndpointsInternalServerError handles this case with default header values.

Unexpected error
*/
type GetDriverEndpointsInternalServerError struct {
	Payload string
}

func (o *GetDriverEndpointsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /driver_endpoints][%d] getDriverEndpointsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDriverEndpointsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}