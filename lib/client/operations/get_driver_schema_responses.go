package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

// GetDriverSchemaReader is a Reader for the GetDriverSchema structure.
type GetDriverSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetDriverSchemaReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDriverSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetDriverSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetDriverSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDriverSchemaOK creates a GetDriverSchemaOK with default headers values
func NewGetDriverSchemaOK() *GetDriverSchemaOK {
	return &GetDriverSchemaOK{}
}

/*GetDriverSchemaOK handles this case with default header values.

OK
*/
type GetDriverSchemaOK struct {
	Payload models.DriverSchema
}

func (o *GetDriverSchemaOK) Error() string {
	return fmt.Sprintf("[GET /drivers/{driver_id}/config_schema][%d] getDriverSchemaOK  %+v", 200, o.Payload)
}

func (o *GetDriverSchemaOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverSchemaNotFound creates a GetDriverSchemaNotFound with default headers values
func NewGetDriverSchemaNotFound() *GetDriverSchemaNotFound {
	return &GetDriverSchemaNotFound{}
}

/*GetDriverSchemaNotFound handles this case with default header values.

Not Found
*/
type GetDriverSchemaNotFound struct {
}

func (o *GetDriverSchemaNotFound) Error() string {
	return fmt.Sprintf("[GET /drivers/{driver_id}/config_schema][%d] getDriverSchemaNotFound ", 404)
}

func (o *GetDriverSchemaNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDriverSchemaInternalServerError creates a GetDriverSchemaInternalServerError with default headers values
func NewGetDriverSchemaInternalServerError() *GetDriverSchemaInternalServerError {
	return &GetDriverSchemaInternalServerError{}
}

/*GetDriverSchemaInternalServerError handles this case with default header values.

Unexpected error
*/
type GetDriverSchemaInternalServerError struct {
	Payload string
}

func (o *GetDriverSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /drivers/{driver_id}/config_schema][%d] getDriverSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDriverSchemaInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
