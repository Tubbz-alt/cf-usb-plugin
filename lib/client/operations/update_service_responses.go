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

type UpdateServiceReader struct {
	formats strfmt.Registry
}

func (o *UpdateServiceReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUpdateServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateServiceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateServiceOK creates a UpdateServiceOK with default headers values
func NewUpdateServiceOK() *UpdateServiceOK {
	return &UpdateServiceOK{}
}

/*UpdateServiceOK

Sucessfull response
*/
type UpdateServiceOK struct {
	Payload *models.Service
}

func (o *UpdateServiceOK) Error() string {
	return fmt.Sprintf("[PUT /services/{service_id}][%d] updateServiceOK  %+v", 200, o.Payload)
}

func (o *UpdateServiceOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceNotFound creates a UpdateServiceNotFound with default headers values
func NewUpdateServiceNotFound() *UpdateServiceNotFound {
	return &UpdateServiceNotFound{}
}

/*UpdateServiceNotFound

Not Found
*/
type UpdateServiceNotFound struct {
}

func (o *UpdateServiceNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/{service_id}][%d] updateServiceNotFound ", 404)
}

func (o *UpdateServiceNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceConflict creates a UpdateServiceConflict with default headers values
func NewUpdateServiceConflict() *UpdateServiceConflict {
	return &UpdateServiceConflict{}
}

/*UpdateServiceConflict

Conflict
*/
type UpdateServiceConflict struct {
}

func (o *UpdateServiceConflict) Error() string {
	return fmt.Sprintf("[PUT /services/{service_id}][%d] updateServiceConflict ", 409)
}

func (o *UpdateServiceConflict) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceInternalServerError creates a UpdateServiceInternalServerError with default headers values
func NewUpdateServiceInternalServerError() *UpdateServiceInternalServerError {
	return &UpdateServiceInternalServerError{}
}

/*UpdateServiceInternalServerError

Unexpected error
*/
type UpdateServiceInternalServerError struct {
	Payload string
}

func (o *UpdateServiceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /services/{service_id}][%d] updateServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateServiceInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}