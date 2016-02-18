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

type GetServiceReader struct {
	formats strfmt.Registry
}

func (o *GetServiceReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceOK creates a GetServiceOK with default headers values
func NewGetServiceOK() *GetServiceOK {
	return &GetServiceOK{}
}

/*GetServiceOK

Sucessfull response
*/
type GetServiceOK struct {
	Payload *models.Service
}

func (o *GetServiceOK) Error() string {
	return fmt.Sprintf("[GET /services/{service_id}][%d] getServiceOK  %+v", 200, o.Payload)
}

func (o *GetServiceOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceNotFound creates a GetServiceNotFound with default headers values
func NewGetServiceNotFound() *GetServiceNotFound {
	return &GetServiceNotFound{}
}

/*GetServiceNotFound

Not Found
*/
type GetServiceNotFound struct {
}

func (o *GetServiceNotFound) Error() string {
	return fmt.Sprintf("[GET /services/{service_id}][%d] getServiceNotFound ", 404)
}

func (o *GetServiceNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceInternalServerError creates a GetServiceInternalServerError with default headers values
func NewGetServiceInternalServerError() *GetServiceInternalServerError {
	return &GetServiceInternalServerError{}
}

/*GetServiceInternalServerError

Unexpected error
*/
type GetServiceInternalServerError struct {
	Payload string
}

func (o *GetServiceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /services/{service_id}][%d] getServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetServiceInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}