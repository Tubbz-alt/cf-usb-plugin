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

// RegisterDriverEndpointReader is a Reader for the RegisterDriverEndpoint structure.
type RegisterDriverEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *RegisterDriverEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewRegisterDriverEndpointCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewRegisterDriverEndpointConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRegisterDriverEndpointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegisterDriverEndpointCreated creates a RegisterDriverEndpointCreated with default headers values
func NewRegisterDriverEndpointCreated() *RegisterDriverEndpointCreated {
	return &RegisterDriverEndpointCreated{}
}

/*RegisterDriverEndpointCreated handles this case with default header values.

Driver endpoint registered
*/
type RegisterDriverEndpointCreated struct {
	Payload *models.DriverEndpoint
}

func (o *RegisterDriverEndpointCreated) Error() string {
	return fmt.Sprintf("[POST /driver_endpoints][%d] registerDriverEndpointCreated  %+v", 201, o.Payload)
}

func (o *RegisterDriverEndpointCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DriverEndpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterDriverEndpointConflict creates a RegisterDriverEndpointConflict with default headers values
func NewRegisterDriverEndpointConflict() *RegisterDriverEndpointConflict {
	return &RegisterDriverEndpointConflict{}
}

/*RegisterDriverEndpointConflict handles this case with default header values.

Conflict
*/
type RegisterDriverEndpointConflict struct {
	Payload string
}

func (o *RegisterDriverEndpointConflict) Error() string {
	return fmt.Sprintf("[POST /driver_endpoints][%d] registerDriverEndpointConflict  %+v", 409, o.Payload)
}

func (o *RegisterDriverEndpointConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterDriverEndpointInternalServerError creates a RegisterDriverEndpointInternalServerError with default headers values
func NewRegisterDriverEndpointInternalServerError() *RegisterDriverEndpointInternalServerError {
	return &RegisterDriverEndpointInternalServerError{}
}

/*RegisterDriverEndpointInternalServerError handles this case with default header values.

Unexpected error
*/
type RegisterDriverEndpointInternalServerError struct {
	Payload string
}

func (o *RegisterDriverEndpointInternalServerError) Error() string {
	return fmt.Sprintf("[POST /driver_endpoints][%d] registerDriverEndpointInternalServerError  %+v", 500, o.Payload)
}

func (o *RegisterDriverEndpointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
