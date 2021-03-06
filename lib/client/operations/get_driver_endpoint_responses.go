package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/SUSE/cf-usb-plugin/lib/models"
)

// GetDriverEndpointReader is a Reader for the GetDriverEndpoint structure.
type GetDriverEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetDriverEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDriverEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetDriverEndpointNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetDriverEndpointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDriverEndpointOK creates a GetDriverEndpointOK with default headers values
func NewGetDriverEndpointOK() *GetDriverEndpointOK {
	return &GetDriverEndpointOK{}
}

/*GetDriverEndpointOK handles this case with default header values.

OK
*/
type GetDriverEndpointOK struct {
	Payload *models.DriverEndpoint
}

func (o *GetDriverEndpointOK) Error() string {
	return fmt.Sprintf("[GET /driver_endpoints/{driver_endpoint_id}][%d] getDriverEndpointOK  %+v", 200, o.Payload)
}

func (o *GetDriverEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DriverEndpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverEndpointNotFound creates a GetDriverEndpointNotFound with default headers values
func NewGetDriverEndpointNotFound() *GetDriverEndpointNotFound {
	return &GetDriverEndpointNotFound{}
}

/*GetDriverEndpointNotFound handles this case with default header values.

Not Found
*/
type GetDriverEndpointNotFound struct {
}

func (o *GetDriverEndpointNotFound) Error() string {
	return fmt.Sprintf("[GET /driver_endpoints/{driver_endpoint_id}][%d] getDriverEndpointNotFound ", 404)
}

func (o *GetDriverEndpointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDriverEndpointInternalServerError creates a GetDriverEndpointInternalServerError with default headers values
func NewGetDriverEndpointInternalServerError() *GetDriverEndpointInternalServerError {
	return &GetDriverEndpointInternalServerError{}
}

/*GetDriverEndpointInternalServerError handles this case with default header values.

Unexpected error
*/
type GetDriverEndpointInternalServerError struct {
	Payload string
}

func (o *GetDriverEndpointInternalServerError) Error() string {
	return fmt.Sprintf("[GET /driver_endpoints/{driver_endpoint_id}][%d] getDriverEndpointInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDriverEndpointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
